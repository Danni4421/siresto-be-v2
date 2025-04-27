package controllers

import (
	"github.com/Danni4421/siresto-be-v2/app/dtos"
	"github.com/Danni4421/siresto-be-v2/app/services"
	"github.com/Danni4421/siresto-be-v2/package/exceptions"
	"github.com/Danni4421/siresto-be-v2/package/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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

func (controller AuthController) RefreshToken(c *fiber.Ctx) error {
	refreshTokenDTO := new(dtos.RefreshTokenDTO)

	if err := utils.ParseAndValidate(c, refreshTokenDTO); err != nil {
		return exceptions.NewBadRequest("Invalid token")
	}

	claims, err := controller.AuthService.ValidateToken(refreshTokenDTO.RefreshToken)

	if err != nil {
		return exceptions.NewUnauthorized("Invalid token")
	}

	user, err := controller.UserService.FindUserByID(uint(claims.(jwt.MapClaims)["user_id"].(float64)))

	if err != nil {
		return exceptions.NewUnauthorized("You are not authorized to refresh this token")
	}

	token, err := controller.AuthService.RenewateToken(user)

	if err != nil {
		return exceptions.NewUnauthorized("Failed to refresh token")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Token refreshed successfully",
		"data": fiber.Map{
			"access_token": token,
		},
	})
}

func (controller AuthController) Logout(c *fiber.Ctx) error {
	userID := uint(c.Locals("userID").(float64))

	err := controller.AuthService.InvalidateToken(userID)

	if err != nil {
		return exceptions.NewUnauthorized("Failed to logout")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Successfully logged out",
	})
}
