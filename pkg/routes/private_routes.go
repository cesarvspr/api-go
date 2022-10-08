package routes

import (
	"github.com/cesarvspr/api-go/app/controllers"
	"github.com/cesarvspr/api-go/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for POST method:
	route.Post("/book", middleware.ApiKeyProtected(), controllers.CreateBook) // create a new book
}
