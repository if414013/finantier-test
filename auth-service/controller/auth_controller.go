package controller

import (
	"auth-service/model"
	"auth-service/service"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	AuthService service.AuthService
}

func NewAuthController(authService *service.AuthService) AuthController {
	return AuthController{AuthService: *authService}
}

func (controller *AuthController) Route(app *fiber.App) {
	app.Get("/token", controller.getToken)
	app.Get("/authenticate", controller.authenticate)
}

func (controller *AuthController) getToken(c *fiber.Ctx) error {
	tokenData := controller.AuthService.GetJwt()
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   tokenData,
	})
}

func (controller *AuthController) authenticate(c *fiber.Ctx) error {
	token := c.Query("token")
	result := controller.AuthService.Authenticate(
		model.AuthRequest{
			Token: token,
		},
	)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	})
}
