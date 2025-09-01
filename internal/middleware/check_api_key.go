package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func CheckAPIKey() fiber.Handler {
	requiredKey := os.Getenv("API_KEY")

	return func(c *fiber.Ctx) error {
		apiKey := c.Get("X-API-Key")

		if apiKey == "" || apiKey != requiredKey {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized - missing or invalid API key",
			})
		}
		return c.Next()
	}
}
