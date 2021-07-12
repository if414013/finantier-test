package main

import (
	"auth-service/config"
	"auth-service/controller"
	"auth-service/exception"
	"auth-service/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"os"
)

func main() {
	// Setup Configuration
	config.New()
	// Setup Service
	authService := service.NewAuthService()

	// Setup Controller
	authController := controller.NewAuthController(&authService)

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	// Setup Routing
	authController.Route(app)

	// Start App
	err := app.Listen(":" + os.Getenv("APPLICATION_PORT"))
	exception.PanicIfNeeded(err)
}
