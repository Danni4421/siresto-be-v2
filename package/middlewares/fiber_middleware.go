package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func FiberMiddleware(app *fiber.App) {
	app.Use(
		cors.New(cors.Config{
			AllowOrigins: "*",
			AllowMethods: "GET,POST,PUT,DELETE,OPTIONS,PATCH",
			AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		}),
		logger.New(),
	)

	app.Use(func(c *fiber.Ctx) error {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"status":  "fail",
					"message": "Internal Server Error",
				})
			}
		}()

		return c.Next()
	})
}
