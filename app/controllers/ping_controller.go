package controllers

import (
	"github.com/Danni4421/siresto-be-v2/app/models"
	"github.com/gofiber/fiber/v2"
)

func Ping(c *fiber.Ctx) error {
	return c.JSON(models.Ping{
		Message: "pong",
	})
}
