package handler

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userID"
)

//userIdentity is a authentification middleware
func (h *Handler) userIdentity(c *fiber.Ctx) error {
	header := c.Request().Header.Peek(authorizationHeader)

	if string(header) == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("empty auth header")
	}
	headerParts := strings.Split(string(header), " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).SendString("empty auth header")
	}

	if len(headerParts[1]) == 0 {
		return c.Status(fiber.StatusUnauthorized).SendString("token is empty")
	}

	userID, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
	}
	userIdStr := strconv.Itoa(userID)

	c.Set(userCtx, userIdStr)

	return nil
}
