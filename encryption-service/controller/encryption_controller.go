package controller

import (
	"encryption-service/exception"
	"encryption-service/model"
	"encryption-service/service"
	"encryption-service/validation"
	"github.com/gofiber/fiber/v2"
)

type EncryptionController struct {
	StockService service.EncryprionService
}

func NewEncryptionController(encryptionService *service.EncryprionService) EncryptionController {
	return EncryptionController{StockService: *encryptionService}
}

func (controller *EncryptionController) Route(app *fiber.App) {
	app.Post("/encrypt", controller.encrypt)
}

func (controller *EncryptionController) encrypt(c *fiber.Ctx) error {
	var request model.StockInfo
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)

	validation.ValidateStockSymbol(request)

	response := controller.StockService.EncryptStockInfo(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}
