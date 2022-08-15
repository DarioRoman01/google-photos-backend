package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/DarioRoman01/photos/bucket"
	"github.com/DarioRoman01/photos/database"
	"github.com/DarioRoman01/photos/mailpb"
	"github.com/DarioRoman01/photos/models"
	"github.com/DarioRoman01/photos/uploadpb"
	"github.com/DarioRoman01/photos/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// CommandService is the service that handles the commands, commands are the write operations
type CommandService struct {
	mailService   mailpb.MailServiceClient     // mailService is the mail service client
	uploadService uploadpb.UploadServiceClient // uploadService is the upload service
}

// NewCommandService creates a new command service
func NewCommandService() (*CommandService, error) {
	// create the database repository
	db, err := database.NewPostgresRepository(os.Getenv("POSTGRES_URL"))
	if err != nil {
		return nil, err
	}

	// create the bucket repository
	s3Bucket, err := bucket.NewS3BucketRepository()
	if err != nil {
		return nil, err
	}

	// create the mail service client
	mailConn, err := grpc.Dial("mailService:5060", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	uploadConn, err := grpc.Dial("uploadService:5070", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	mailService := mailpb.NewMailServiceClient(mailConn)
	uploadService := uploadpb.NewUploadServiceClient(uploadConn)
	database.SetDatabaseRepository(db)
	bucket.SetBucketRepository(s3Bucket)

	return &CommandService{mailService, uploadService}, nil
}

// RegisterHandler handles the registration of a new user request
func (s *CommandService) RegisterHandler(c *fiber.Ctx) error {
	input := new(models.UserLoginRegister)
	if err := c.BodyParser(input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.JsonError("Invalid body"))
	}

	hashPwd, err := utils.GeneratePassword(utils.GetDefaultPasswordConfig(), input.Password)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.JsonError("Error generating password"))
	}

	input.Password = hashPwd
	user := &models.User{
		ID:       uuid.NewString(),
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}

	if err := database.InsertUser(user); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.JsonError("Error creating user"))
	}

	token, err := utils.CreateToken(user.Username, user.ID, models.VerifyToken.String())
	if err != nil {
		log.Printf("Error creating token: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(utils.JsonError("Error creating token"))
	}

	_, err = s.mailService.SendMail(c.Context(), &mailpb.SendMailRequest{
		Type:     models.VerifyToken.String(),
		Receiver: input.Email,
		Subject:  "Verify your email",
		Body:     "Please verify your email",
		User:     user.Username,
		Token:    token,
	})

	if err != nil {
		log.Printf("token: %v", token)
		log.Printf("Error sending email: %v", err)
	}

	return c.Status(http.StatusCreated).JSON(user)
}

// HandleVerifyEmail handles the verification of a new user request
func (s *CommandService) HandleVerify(c *fiber.Ctx) error {
	rawToken := c.Query("token")
	if rawToken == "" {
		return c.Status(http.StatusBadRequest).JSON(utils.JsonError("Invalid token"))
	}

	claims, err := utils.VerifyToken(rawToken)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.JsonError("Invalid token"))
	}

	if claims.Type != models.VerifyToken.String() {
		return c.Status(http.StatusBadRequest).JSON(utils.JsonError("Invalid token"))
	}

	if err := database.UpdateUserStatus(claims.UserID); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.JsonError("Error updating user"))
	}

	return c.Status(http.StatusOK).JSON(map[string]string{"message": "User verified"})
}

func (s *CommandService) LoginHandler(c *fiber.Ctx) error {
	input := new(models.UserLoginRegister)
	if err := c.BodyParser(input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.JsonError("Invalid body"))
	}

	user, err := database.GetUserByEmail(input.Email)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.JsonError("Error getting user"))
	}

	if user == nil {
		return c.Status(http.StatusBadRequest).JSON(utils.JsonError("User not found"))
	}

	ok, _ := utils.ComparePasswords(input.Password, user.Password)

	if !ok {
		return c.Status(http.StatusBadRequest).JSON(utils.JsonError("Invalid password"))
	}

	token, err := utils.CreateToken(user.Username, user.ID, models.AccessToken.String())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.JsonError("Error creating token"))
	}

	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(time.Hour * 24 * 7),
	})

	return c.Status(http.StatusOK).JSON(map[string]string{"token": token})
}

func (s *CommandService) LogoutHandler(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
	})

	return c.Status(http.StatusOK).JSON(map[string]string{"message": "User logged out"})
}

func (s *CommandService) getUploadData(c *fiber.Ctx) (*models.UploadRequest, error) {
	folder := c.Query("path")
	if folder == "" {
		folder = "default"
	}

	userId := c.Locals("user_id").(string)
	username := c.Locals("username").(string)
	folderId, err := database.CheckFolder(userId, folder)
	if err != nil {
		return nil, err
	}

	fileHeader, err := c.FormFile("file")
	if err != nil {
		return nil, err
	}

	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}

	return &models.UploadRequest{
		FolderName: folder,
		FolderID:   folderId,
		UserID:     userId,
		Username:   username,
		Filename:   fileHeader.Filename,
		File:       file,
	}, nil
}

// streamData streams the data to the upload service
func (s *CommandService) streamData(c *fiber.Ctx, req *models.UploadRequest) (string, error) {
	stream, err := s.uploadService.Upload(c.Context())
	if err != nil {
		return "", c.Status(http.StatusInternalServerError).JSON(utils.JsonError("Error uploading file"))
	}

	buf := make([]byte, 1024)
	for {
		n, err := req.File.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}

			return "", c.Status(http.StatusInternalServerError).JSON(utils.JsonError("Error reading file"))
		}

		err = stream.Send(&uploadpb.UploadRequest{
			Username: req.Username,
			Folder:   req.FolderName,
			Filename: req.Filename,
			Chunk:    buf[:n],
		})

		if err != nil {
			return "", c.Status(http.StatusInternalServerError).JSON(utils.JsonError("Error uploading file"))
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return "", c.Status(http.StatusInternalServerError).JSON(utils.JsonError("Error uploading file"))
	}

	return res.Location, nil
}

func (s *CommandService) UploadHandler(c *fiber.Ctx) error {
	req, err := s.getUploadData(c)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.JsonError("Error getting upload data"))
	}

	defer req.File.Close()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.JsonError("Invalid file"))
	}

	location, err := s.streamData(c, req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.JsonError("Error uploading file"))
	}

	img := &models.Image{
		ID:       uuid.NewString(),
		UserID:   req.UserID,
		Name:     req.Filename,
		FolderID: req.FolderID,
		URL:      location,
	}

	if err := database.InsertImage(img); err != nil {
		log.Printf("Error inserting image: %s", err)
		return c.Status(http.StatusInternalServerError).JSON(utils.JsonError("Error inserting file"))
	}

	return c.Status(http.StatusOK).JSON(img)
}

func (s *CommandService) DeleteImageHandler(c *fiber.Ctx) error {
	fileName := c.Params("filename")
	folder := c.Query("path")
	if folder == "" {
		folder = "default"
	}

	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(utils.JsonError("Invalid id"))
	}

	if err := bucket.Delete(c.Locals("username").(string), fileName, folder); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.JsonError("Error deleting file"))
	}

	if err := database.DeleteImage(id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.JsonError("Error deleting image"))
	}

	return c.Status(http.StatusOK).JSON(map[string]string{"message": "Image deleted"})
}
