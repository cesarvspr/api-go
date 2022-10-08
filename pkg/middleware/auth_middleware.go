package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"

	SecretMiddleware "github.com/gofiber/jwt/v2"
)

// ApiKeyProtected func for specify routes group with API-KEY authentication.
func ApiKeyProtected() func(*fiber.Ctx) error {

	// Create config for API-KEY authentication middleware.
	config := SecretMiddleware.Config{
		SigningKey:   []byte(os.Getenv("API-KEY")),
		ContextKey:   "api-key", // used in private routes
		ErrorHandler: Validate,
	}

	return SecretMiddleware.New(config)
}

func Validate(c *fiber.Ctx, err error) error {
	// Get token from header.
	token := c.Get("Authorization")

	if token == os.Getenv("API-KEY") {
		return c.Next()
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "Unauthorized",
	})
}
