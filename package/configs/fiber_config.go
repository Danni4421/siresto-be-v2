package configs

import (
	"os"
	"strconv"
	"time"

	"github.com/Danni4421/siresto-be-v2/package/exceptions"
	"github.com/gofiber/fiber/v2"
)

func FiberConfig() fiber.Config {
	readTimeout, _ := strconv.Atoi(os.Getenv("APP_READ_TIMEOUT"))

	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(readTimeout),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			switch e := err.(type) {
			case exceptions.ValidationError:
				return c.Status(e.StatusCode()).JSON(fiber.Map{
					"status":  "error",
					"message": e.Message,
					"errors":  e.Errors,
				})
			case exceptions.ClientError:
				return c.Status(e.StatusCode()).JSON(fiber.Map{
					"status":  "fail",
					"message": e.Message,
				})
			default:
				code := fiber.StatusInternalServerError
				if fiberErr, ok := err.(*fiber.Error); ok {
					code = fiberErr.Code
				}
				return c.Status(code).JSON(fiber.Map{
					"status":  "error",
					"message": err.Error(),
				})
			}
		},
	}
}
