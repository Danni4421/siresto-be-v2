package routes

import (
	"github.com/gofiber/fiber/v2"
)

func AuthenticatedRoutes(app *fiber.App) {
	route := app.Group("/api/v2")

	route.Get("/users", userController.GetUsers)
	route.Get("/users/:id", userController.GetUserByID)
	route.Patch("/users/:id", userController.UpdateUser)
	route.Delete("/users/:id", userController.DeleteUser)
}
