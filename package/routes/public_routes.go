package routes

import (
	"github.com/Danni4421/siresto-be-v2/app/controllers"
	"github.com/Danni4421/siresto-be-v2/app/services"
	"github.com/Danni4421/siresto-be-v2/platform/database"
	"github.com/gofiber/fiber/v2"
)

var userController *controllers.UserController
var authController *controllers.AuthController

func init() {
	dbInstance := database.GetDatabase()

	userController = &controllers.UserController{
		UserService: &services.UserService{
			DB: dbInstance,
		},
	}

	authController = &controllers.AuthController{
		AuthService: &services.AuthService{
			DB: dbInstance,
		},
		UserService: &services.UserService{
			DB: dbInstance,
		},
	}
}

func PublicRoutes(app *fiber.App) {
	route := app.Group("/api/v2")

	// Default ping route
	route.Get("/ping", controllers.Ping)
	route.Post("/users", userController.Register)

	// Authentication routes
	route.Post("/login", authController.Login)
}
