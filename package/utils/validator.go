package utils

import (
	"errors"

	"github.com/Danni4421/siresto-be-v2/package/exceptions"
	validator "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type Validatable interface {
	ErrorMessages() map[string]string
}

func ParseAndValidate(c *fiber.Ctx, payload any) error {
	if err := c.BodyParser(payload); err != nil {
		return exceptions.NewBadRequest("Request unprocessable")
	}

	if err := validate.Struct(payload); err != nil {
		errorBag := make(map[string]string)

		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			customMessages := make(map[string]string)

			if v, ok := payload.(Validatable); ok {
				customMessages = v.ErrorMessages()
			}

			for _, fieldErr := range validationErrors {
				field := fieldErr.Field()
				tag := fieldErr.Tag()
				key := field + "." + tag

				if msg, exists := customMessages[key]; exists {
					errorBag[field] = msg
				} else {
					errorBag[field] = field + " is invalid"
				}
			}
		}

		return exceptions.ValidationError{
			Message: "Request unprocessable",
			Errors:  errorBag,
		}
	}

	return nil
}
