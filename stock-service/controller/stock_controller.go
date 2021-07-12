package controller

import (
	"github.com/gofiber/fiber/v2"
	middleware "stock-service/auth"
	"stock-service/model"
	"stock-service/service"
	"stock-service/validation"
)

type StockController struct {
	StockService service.StockService
}

func NewStockController(stockService *service.StockService) StockController {
	return StockController{StockService: *stockService}
}

func (controller *StockController) Route(app *fiber.App) {
	app.Get("/symbol/:id", controller.getStockInfoBySymbol)
}

func (controller *StockController) getStockInfoBySymbol(c *fiber.Ctx) error {
	token := c.Get("token")
	if !middleware.AuthenticateRequest(token) {
		return c.JSON(model.WebResponse{
			Code:   401,
			Status: "Unauthorized",
			Data:   "Missing token param in Request Header or Invalid or Expired token try to re-fetch it from /token",
		})
	}
	stockName := c.Params("id")
	validation.ValidateStockSymbol(stockName)
	response := controller.StockService.GetStockInfoBySymbol(stockName)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}
