package routes

import "github.com/gofiber/fiber/v2"

type PingResponse struct {
	Message string `json:"message"`
}

func PublicRoutes(app *fiber.App) {
	route := app.Group("/api/v2")

	// Default ping route
	route.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(PingResponse{
			Message: "pong",
		})
	})
}
