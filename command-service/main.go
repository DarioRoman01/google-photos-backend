// command service is a service that handles commands from the client this commands are just write actions to the repositorys
package main

import (
	"log"

	"github.com/DarioRoman01/photos/middlewares"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	commandService, err := NewCommandService()
	if err != nil {
		log.Fatalf("Error creating command service: %v", err)
	}

	app.Use(middlewares.CheckAuthMiddleware())
	app.Post("/users/signup", commandService.RegisterHandler)
	app.Post("/folders/create", commandService.CreateFolderHandler)
	app.Post("/users/login", commandService.LoginHandler)
	app.Post("/images/upload", commandService.UploadHandler)
	app.Put("/images/move", commandService.MoveFileHandler)
	app.Post("/users/verify", commandService.HandleVerify)
	app.Delete("/images/delete/:filename/:id", commandService.DeleteImageHandler)

	app.Listen(":3000")
}
