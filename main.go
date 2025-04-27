package main

import (
	"fmt"

	"github.com/Danni4421/siresto-be-v2/package/configs"
	"github.com/Danni4421/siresto-be-v2/package/middlewares"
	"github.com/Danni4421/siresto-be-v2/package/routes"
	"github.com/Danni4421/siresto-be-v2/package/utils"
	"github.com/Danni4421/siresto-be-v2/platform/database"
	"github.com/Danni4421/siresto-be-v2/platform/migrations"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fiberConfig := configs.FiberConfig()

	application := fiber.New(fiberConfig)

	databaseInstance := database.GetDatabase()
	migrationError := migrations.AutoMigrate(databaseInstance)

	if migrationError != nil {
		panic("Error migrating system database")
	}

	// Bind middlewares
	middlewares.FiberMiddleware(application)

	// Bind routes
	routes.PublicRoutes(application)
	routes.AuthenticatedRoutes(application)

	err := application.Listen(fmt.Sprintf(":%s", utils.GetEnv("APP_PORT", "8585")))

	if err != nil {
		panic(err)
	}
}
