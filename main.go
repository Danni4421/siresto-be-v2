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
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	configs := configs.FiberConfig()

	app := fiber.New(configs)

	// Auto migrations
	db, db_err := database.ConnectDB()

	if db_err != nil {
		panic(db_err)
	}

	migrations.AutoMigrate(db)

	// Bind middlewares
	middlewares.FiberMiddleware(app)

	// Bind routes
	routes.PublicRoutes(app)

	err := app.Listen(fmt.Sprintf(":%s", utils.GetEnv("APP_PORT", "8585")))

	if err != nil {
		panic(err)
	}
}
