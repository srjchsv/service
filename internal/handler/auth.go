package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/srjchsv/service/internal/repository"
)

//SignUp performs users sign-up on handlers level
func (h *Handler) signUp(c *fiber.Ctx) error {
	var input repository.User

	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//SignIn performs users sign-in on handlers level
func (h *Handler) signIn(c *fiber.Ctx) error {
	var input signInInput

	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
	})

	return c.JSON(map[string]interface{}{
		"token": token,
	})
}
