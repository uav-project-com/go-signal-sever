package api

import (
	"boilerplate/api/handler/user"

	"github.com/gofiber/fiber/v2"
)

// Set up routes
func SetupRoutes(app *fiber.App, userHandler *user.UserHandler) {
	api := app.Group("/api/v1")

	userGroup := api.Group("/user")
	// GET /api/v1/user/:id
	userGroup.Get("/:id", userHandler.GetUserById)
	// POST /api/v1/user
	userGroup.Post("/", userHandler.CreateUser)

	// Webrtc router
}
