package api

import (
	"boilerplate/api/handler/user"
	"boilerplate/api/handler/webrtc"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes Set up routes
func SetupRoutes(app *fiber.App, userHandler *user.UserHandler) {
	api := app.Group("/api/v1/user")
	// GET /api/v1/user/:id
	api.Get("/:id", userHandler.GetUserById)
	// POST /api/v1/user
	api.Post("/", userHandler.CreateUser)
}

// SetupWebrtcRouters handler webrtc signal api
func SetupWebrtcRouters(app *fiber.App, handler *webrtc.ConnectHandler) {
	api := app.Group("/api/v1/webrtc")
	api.Post("/start-call", handler.StartCall)
}
