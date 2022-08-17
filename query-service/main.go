package main

import (
	"log"

	"github.com/DarioRoman01/photos/middlewares"
	"github.com/gofiber/fiber/v2"
)

func main() {
	svc, err := NewQueryService()
	if err != nil {
		log.Fatalf("Error creating query service: %v", err)
	}

	app := fiber.New()

	app.Use(middlewares.CheckAuthMiddleware())
	app.Get("/images", svc.GetImagesHandler)
	app.Get("/images/:imageID", svc.GetImageHandler)
	app.Get("folders", svc.GetFoldersHandler)
	app.Get("folders/:folderID", svc.GetImageByFolder)

	app.Listen(":3001")
}
