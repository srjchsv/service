package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/srjchsv/service/internal/services"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouter(app *fiber.App) *fiber.App {
	app.Use(logger.New())
	// app.Use(favicon.New(favicon.Config{
	// 	File: "./static/resources/favicon.ico",
	// }))

	auth := app.Group("/auth")
	auth.Post("/sign-up", h.signUp)
	auth.Get("/sign-up", func(c *fiber.Ctx) error {
		return c.SendString("Signup here...")
	})
	auth.Post("/sign-in", h.signIn)

	apiV1 := app.Group("/api", h.userIdentity)
	apiV1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(c.Get(userCtx))
	})

	return app
}
