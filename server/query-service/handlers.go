package main

import (
	"os"
	"strconv"

	"github.com/DarioRoman01/photos/bucket"
	"github.com/DarioRoman01/photos/database"
	"github.com/DarioRoman01/photos/utils"
	"github.com/gofiber/fiber/v2"
)

type QueryService struct{}

func NewQueryService() (*QueryService, error) {
	s3repo, err := bucket.NewS3BucketRepository()
	if err != nil {
		return nil, err
	}

	postgresRepo, err := database.NewPostgresRepository(os.Getenv("POSTGRES_URL"))
	if err != nil {
		return nil, err
	}

	database.SetDatabaseRepository(postgresRepo)
	bucket.SetBucketRepository(s3repo)
	return &QueryService{}, nil
}

func (s *QueryService) GetImagesHandler(c *fiber.Ctx) error {
	cursor := c.Query("cursor")
	strLimit := c.Query("limit")
	if strLimit == "" {
		strLimit = "10"
	}

	limit, err := strconv.Atoi(strLimit)
	if err != nil {
		return c.Status(500).JSON(utils.JsonError("Error parsing limit"))
	}

	userID := c.Locals("user_id").(string)
	images, err, hasMore := database.GetImages(userID, cursor, limit)
	if err != nil {
		return c.Status(500).JSON(utils.JsonError("Error getting images"))
	}

	return c.Status(200).JSON(fiber.Map{
		"images":  images,
		"hasMore": hasMore,
	})
}

func (s *QueryService) GetImageHandler(c *fiber.Ctx) error {
	imageID := c.Params("imageID")
	image, err := database.GetImage(imageID)
	if err != nil {
		return c.Status(500).JSON(utils.JsonError("Error getting image"))
	}

	return c.Status(200).JSON(image)
}

func (s *QueryService) GetImageByFolder(c *fiber.Ctx) error {
	folder := c.Params("folder")
	cursor := c.Query("cursor")
	strLimit := c.Query("limit")
	if strLimit == "" {
		strLimit = "10"
	}

	limit, err := strconv.Atoi(strLimit)
	if err != nil {
		return c.Status(500).JSON(utils.JsonError("Error parsing limit"))
	}

	userID := c.Locals("user_id").(string)

	images, hasMore, err := database.GetImagesByFolder(userID, folder, cursor, limit)
	if err != nil {
		return c.Status(500).JSON(utils.JsonError("Error getting image"))
	}

	return c.Status(200).JSON(fiber.Map{
		"images":  images,
		"hasMore": hasMore,
	})
}

func (s *QueryService) GetFoldersHandler(c *fiber.Ctx) error {
	cursor := c.Query("cursor")
	strLimit := c.Query("limit")
	if strLimit == "" {
		strLimit = "10"
	}

	limit, err := strconv.Atoi(strLimit)
	if err != nil {
		return c.Status(500).JSON(utils.JsonError("Error parsing limit"))
	}

	userID := c.Locals("user_id").(string)
	folders, hasMore, err := database.GetFolders(userID, cursor, limit)
	if err != nil {
		return c.Status(500).JSON(utils.JsonError("Error getting folders"))
	}

	return c.Status(200).JSON(fiber.Map{
		"folders": folders,
		"hasMore": hasMore,
	})
}
