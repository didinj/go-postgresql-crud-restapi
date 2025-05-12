package main

import (
	"github.com/didinj/go-crud-api/database"
	"github.com/didinj/go-crud-api/models"
	"github.com/didinj/go-crud-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Connect to DB
	database.Connect()

	// Auto-migrate Book model to create the table
	database.DB.AutoMigrate(&models.Book{})

	// Setup API routes
	routes.SetupRoutes(app)

	// Start the server
	app.Listen(":3000")
}
