package routes

import (
	"github.com/Danni4421/siresto-be-v2/app/controllers"
	"github.com/Danni4421/siresto-be-v2/app/services"
	"github.com/Danni4421/siresto-be-v2/platform/database"
)

var userController *controllers.UserController
var authController *controllers.AuthController
var menuCategoryController *controllers.MenuCategoryController

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

	menuCategoryController = &controllers.MenuCategoryController{
		MenuCategoryService: &services.MenuCategoryService{
			DB: dbInstance,
		},
	}

}
