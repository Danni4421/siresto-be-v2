package routes

import (
	"github.com/Danni4421/siresto-be-v2/package/middlewares"
	"github.com/gofiber/fiber/v2"
)

func AuthenticatedRoutes(app *fiber.App) {
	route := app.Group("/api/v2")

	route.Get("/users", middlewares.JWTProtected(), userController.GetUsers)
	route.Get("/users/:id", middlewares.JWTProtected(), userController.GetUserByID)
	route.Patch("/users/:id", middlewares.JWTProtected(), userController.UpdateUser)
	route.Delete("/users/:id", middlewares.JWTProtected(), userController.DeleteUser)
}
