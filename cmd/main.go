package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/helmet/v2"
	"github.com/joho/godotenv"
	"github.com/khoidh24/short-url/internal/config"
	"github.com/khoidh24/short-url/internal/database"
	"github.com/khoidh24/short-url/internal/models"
	"github.com/khoidh24/short-url/internal/routes"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env file not loaded")
	}

	// Load config after .env
	cfg := config.LoadConfig()

	// Init app
	app := fiber.New()

	// Helmet middleware
	app.Use(helmet.New())

	// Connect DB
	db := database.ConnectDB()
	if err := db.AutoMigrate(&models.URL{}); err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}

	// Routes
	routes.SetupRoutes(app)

	// Start server
	log.Printf("Server is running at http://localhost:%s", cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}
