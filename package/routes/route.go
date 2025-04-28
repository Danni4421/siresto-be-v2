package routes

import (
	"github.com/Danni4421/siresto-be-v2/app/controllers"
	"github.com/Danni4421/siresto-be-v2/app/services"
	"github.com/Danni4421/siresto-be-v2/platform/database"
)

var (
	userController         *controllers.UserController
	authController         *controllers.AuthController
	menuCategoryController *controllers.MenuCategoryController
	menuController         *controllers.MenuController
)

func init() {
	dbInstance := database.GetDatabase()
	
	userService := &services.UserService{DB: dbInstance}
	menuCategoryService := &services.MenuCategoryService{DB: dbInstance}
	
	userController = &controllers.UserController{
		UserService: userService,
	}

	authController = &controllers.AuthController{
		AuthService: &services.AuthService{DB: dbInstance},
		UserService: userService,
	}

	menuCategoryController = &controllers.MenuCategoryController{
		MenuCategoryService: menuCategoryService,
		UserService: userService,
	}

	menuController = &controllers.MenuController{
		MenuService: &services.MenuService{DB: dbInstance},
		MenuCategoryService: menuCategoryService,
		UserService: userService,
	}
}
