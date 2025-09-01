package routes

import (
	"github.com/gofiber/fiber/v2"

	handlers "github.com/khoidh24/short-url/internal/handlers/url"
	"github.com/khoidh24/short-url/internal/middleware"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/shorten", middleware.CheckAPIKey(), handlers.ShortenURL)
	app.Get("/:shortId", handlers.RedirectURL)
}
