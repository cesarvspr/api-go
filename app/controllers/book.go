package controllers

import (
	"time"

	"github.com/cesarvspr/api-go/app/models"
	"github.com/cesarvspr/api-go/pkg/utils"
	dbClient "github.com/cesarvspr/api-go/platform/database"
	"github.com/gofiber/fiber/v2"
)

// GetBooks func gets all exists books.
// @Description Get all exists books.
// @Summary get all exists books
// @Tags Books
// @Accept json
// @Produce json
// @Success 200 {array} models.Book
// @Router /v1/books [get]
func GetBooks(c *fiber.Ctx) error {

	books, err := dbClient.GetAllBooks()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(books)
}

// CreateBook func for creates a new book.
// @Description Create a new book.
// @Summary create a new book
// @Tags Book
// @Accept json
// @Produce json
// @Param title body string true "Title"
// @Param author body string true "Author"
// @Param book_attrs body models.BookAttrs true "Book attributes"
// @Success 200 {object} models.Book
// @Security ApiKeyAuth
// @Router /v1/book [post]
func CreateBook(c *fiber.Ctx) error {

	// Get now time.
	now := time.Now()

	// Create new Book struct
	bodyBook := &models.Book{}

	// Create a new validator for a Book model.
	validate := utils.NewValidator()

	// Check, if received JSON data is valid.
	if err := c.BodyParser(&bodyBook); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	bodyBook.CreatedAt = now
	bodyBook.UpdatedAt = now

	// Validate book fields.
	if err := validate.Struct(bodyBook); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	_, err := dbClient.CreateBook(bodyBook)

	if err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"res":   bodyBook,
	})
}
