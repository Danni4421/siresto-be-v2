package middlewares

import (
	"fmt"

	"github.com/Danni4421/siresto-be-v2/package/exceptions"
	"github.com/Danni4421/siresto-be-v2/package/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret []byte

func init() {
	jwtSecret = []byte(utils.GetEnv("AUTH_SECRET", ""))

	if len(jwtSecret) == 0 {
		panic("You must set your secret key for authentication")
	}
}

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		actualToken := c.Get("Authorization")

		if actualToken == "" {
			return exceptions.NewUnauthorized("Your not authorized, please authenticate first")
		}

		var authToken string
		_, err := fmt.Sscanf(actualToken, "Bearer %s", &authToken)

		if err != nil {
			return exceptions.NewUnauthorized("Invalid token format")
		}

		claims, err := utils.ValidateToken(authToken, jwtSecret)

		if err != nil {
			return exceptions.NewUnauthorized("Invalid token")
		}

		c.Locals("userID", claims.(jwt.MapClaims)["user_id"])

		return c.Next()
	}
}
