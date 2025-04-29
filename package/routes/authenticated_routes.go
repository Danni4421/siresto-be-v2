package routes

import (
	"github.com/Danni4421/siresto-be-v2/package/middlewares"
	"github.com/gofiber/fiber/v2"
)

func AuthenticatedRoutes(app *fiber.App) {
	route := app.Group("/api/v2", middlewares.JWTProtected())

	route.Delete("/logout", authController.Logout)

	route.Get("/users", userController.GetUsers)
	route.Post("/users/internal", userController.CreateInternalUser)

	route.Get("/users/:id", userController.GetUserByID)
	route.Patch("/users/:id", userController.UpdateUser)
	route.Delete("/users/:id", userController.DeleteUser)

	route.Post("/menu-categories", menuCategoryController.CreateCategory)
	route.Put("/menu-categories/:id", menuCategoryController.UpdateCategory)
	route.Delete("/menu-categories/:id", menuCategoryController.DeleteCategory)

	route.Post("/menus", menuController.CreateMenu)
	route.Put("/menus/:id", menuController.UpdateMenu)
	route.Delete("/menus/:id", menuController.DeleteMenu)

	route.Post("/transactions", transactionController.CreateTransaction)
	route.Get("/transactions", transactionController.GetTransactions)
	route.Get("/transactions/:id", transactionController.GetTransactionByID)
}
