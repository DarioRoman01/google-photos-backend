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
	app.Post("/signup", commandService.RegisterHandler)
	app.Post("/login", commandService.LoginHandler)
	app.Post("/upload", commandService.UploadHandler)
	app.Post("/move", commandService.MoveFileHandler)
	app.Post("/verify", commandService.HandleVerify)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
