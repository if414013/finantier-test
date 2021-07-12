package main

import (
	"encryption-service/config"
	"encryption-service/controller"
	"encryption-service/exception"
	"encryption-service/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"os"
)

func main() {
	// Setup Configuration
	config.New()
	// Setup Service
	encryptionService := service.NewEncryptionService()

	// Setup Controller
	encryptionController := controller.NewEncryptionController(&encryptionService)

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	// Setup Routing
	encryptionController.Route(app)

	// Start App
	err := app.Listen(":" + os.Getenv("APPLICATION_PORT"))
	exception.PanicIfNeeded(err)
}
