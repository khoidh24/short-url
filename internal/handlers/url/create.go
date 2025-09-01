package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khoidh24/short-url/internal/database"
	"github.com/khoidh24/short-url/internal/models"
	"github.com/teris-io/shortid"
)

func ShortenURL(c *fiber.Ctx) error {
	type Request struct {
		OriginalURL string `json:"originalUrl"`
	}

	var body Request
	if err := c.BodyParser(&body); err != nil || body.OriginalURL == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	sid, _ := shortid.Generate()

	url := models.URL{
		ShortID:     sid,
		OriginalURL: body.OriginalURL,
	}

	if err := database.ConnectDB().Create(&url).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create URL",
		})
	}

	return c.JSON(fiber.Map{
		"shortUrl": c.BaseURL() + "/" + sid,
	})
}
