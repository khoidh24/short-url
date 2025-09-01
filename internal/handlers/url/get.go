package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khoidh24/short-url/internal/database"
	"github.com/khoidh24/short-url/internal/models"
)

func RedirectURL(c *fiber.Ctx) error {
	shortID := c.Params("shortId")

	var url models.URL

	if err := database.ConnectDB().First(&url, "short_id = ?", shortID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "URL not found",
		})
	}

	return c.Redirect(url.OriginalURL, fiber.StatusFound)
}
