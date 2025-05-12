package routes

import (
	"github.com/didinj/go-crud-api/database"
	"github.com/didinj/go-crud-api/models"

	"github.com/gofiber/fiber/v2"
)

// GET /api/books
func GetBooks(c *fiber.Ctx) error {
	var books []models.Book
	database.DB.Find(&books)
	return c.JSON(books)
}

// GET /api/books/:id
func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	result := database.DB.First(&book, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Book not found"})
	}
	return c.JSON(book)
}

// POST /api/books
func CreateBook(c *fiber.Ctx) error {
	book := new(models.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	database.DB.Create(&book)
	return c.JSON(book)
}

// PUT /api/books/:id
func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	if err := database.DB.First(&book, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Book not found"})
	}

	if err := c.BodyParser(&book); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	database.DB.Save(&book)
	return c.JSON(book)
}

// DELETE /api/books/:id
func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	result := database.DB.Delete(&models.Book{}, id)
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Book not found"})
	}
	return c.JSON(fiber.Map{"message": "Book deleted successfully"})
}

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	books := api.Group("/books")

	books.Get("/", GetBooks)
	books.Get("/:id", GetBook)
	books.Post("/", CreateBook)
	books.Put("/:id", UpdateBook)
	books.Delete("/:id", DeleteBook)
}
