package routes

import (
	"github.com/Danni4421/siresto-be-v2/app/controllers"
	"github.com/gofiber/fiber/v2"
)

type PingResponse struct {
	Message string `json:"message"`
}

func PublicRoutes(app *fiber.App) {
	route := app.Group("/api/v2")

	// Default ping route
	route.Get("/ping", controllers.Ping)
}
