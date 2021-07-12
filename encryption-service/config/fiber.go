package config

import (
	"github.com/gofiber/fiber/v2"
	"encryption-service/exception"
)

func NewFiberConfig() fiber.Config {
	//use custom error handler
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}
