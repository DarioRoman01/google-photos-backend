package main

import (
	"net/http"
	"os"

	"github.com/DarioRoman01/photos/bucket"
	"github.com/DarioRoman01/photos/database"
	"github.com/DarioRoman01/photos/models"
	"github.com/DarioRoman01/photos/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CommandService struct {
	db     database.DatabaseRepository
	bucket bucket.BucketRepository
}

func NewCommandService() (*CommandService, error) {
	db, err := database.NewPostgresRepository(os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	s3Bucket, err := bucket.NewS3BucketRepository()
	if err != nil {
		return nil, err
	}

	database.SetDatabaseRepository(db)
	bucket.SetBucketRepository(s3Bucket)

	return &CommandService{db, s3Bucket}, nil
}

func (s *CommandService) CreateUserHandler(c *fiber.Ctx) error {
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

	return c.Status(http.StatusCreated).JSON(user)
}
