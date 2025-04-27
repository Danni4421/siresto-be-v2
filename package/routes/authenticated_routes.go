package routes

import (
	"github.com/Danni4421/siresto-be-v2/package/middlewares"
	"github.com/gofiber/fiber/v2"
)

func AuthenticatedRoutes(app *fiber.App) {
	route := app.Group("/api/v2", middlewares.JWTProtected())

	route.Get("/users", userController.GetUsers)
	route.Get("/users/:id", userController.GetUserByID)
	route.Patch("/users/:id", userController.UpdateUser)
	route.Delete("/users/:id", userController.DeleteUser)
}
