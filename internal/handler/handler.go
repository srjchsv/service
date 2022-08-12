package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/srjchsv/service/internal/services"
)

type Handler struct {
	services *services.Service
}

// NewHandler creates a handler instance of services
func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

// InitRouter register router and sets handlers
func (h *Handler) InitRouter(app *fiber.App) *fiber.App {
	app.Use(logger.New())

	auth := app.Group("/auth")
	auth.Post("/sign-up", h.signUp)
	auth.Post("/sign-in", h.signIn)

	apiV1 := app.Group("/api", h.userIdentity)
	apiV1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(c.Get(userCtx))
	})

	return app
}
