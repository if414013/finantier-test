package exception

import (
	"auth-service/model"
	"github.com/gofiber/fiber/v2"
)


func ErrorHandler(ctx *fiber.Ctx, err error) error {

	return ctx.JSON(model.WebResponse{
		Code:   500,
		Status: "INTERNAL_SERVER_ERROR",
		Data:   err.Error(),
	})
}
