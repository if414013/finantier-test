package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"os"
	"stock-service/config"
	"stock-service/controller"
	"stock-service/exception"
	"stock-service/repository"
	"stock-service/service"

)

func main() {
	// Setup Configuration
	config.New()
	// Setup Repository
	stockRepository := repository.NewStockRepository()

	// Setup Service
	stockService := service.NewStockService(&stockRepository)

	// Setup Controller
	stockController := controller.NewStockController(&stockService)

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	// Setup Routing
	stockController.Route(app)

	// Start App
	err := app.Listen(":" + os.Getenv("APPLICATION_PORT"))
	exception.PanicIfNeeded(err)
}
