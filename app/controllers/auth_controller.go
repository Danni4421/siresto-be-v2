package controllers

import (
	"github.com/Danni4421/siresto-be-v2/app/dtos"
	"github.com/Danni4421/siresto-be-v2/app/services"
	"github.com/Danni4421/siresto-be-v2/package/exceptions"
	"github.com/Danni4421/siresto-be-v2/package/utils"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	UserService *services.UserService
	AuthService *services.AuthService
}

func (controller AuthController) Login(c *fiber.Ctx) error {
	authDTO := new(dtos.AuthDTO)

	if err := utils.ParseAndValidate(c, authDTO); err != nil {
		return err
	}

	user, err := controller.UserService.FindUserByEmail(authDTO.Email)

	if err != nil {
		return exceptions.NewUnauthorized("Invalid email or password")
	}

	token, refreshToken, err := controller.AuthService.Authenticate(user, authDTO.Password)

	if err != nil {
		return exceptions.NewUnauthorized("Invalid email or password")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Successfully logged in",
		"data": fiber.Map{
			"access_token":  token,
			"refresh_token": refreshToken,
		},
	})
}
