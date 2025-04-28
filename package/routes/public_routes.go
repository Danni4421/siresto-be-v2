package routes

import (
	"github.com/Danni4421/siresto-be-v2/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App) {
	route := app.Group("/api/v2")

	// Default ping route
	route.Get("/ping", controllers.Ping)
	route.Post("/users", userController.Register)

	// Authentication routes
	route.Post("/login", authController.Login)
	route.Put("/refresh-token", authController.RefreshToken)

	route.Get("/menu-categories", menuCategoryController.GetAllCategories)
	route.Get("/menu-categories/:id", menuCategoryController.GetCategoryByID)
	
	route.Get("/menus", menuController.GetMenus)
	route.Get("/menus/:id", menuController.GetMenuByID)
}
