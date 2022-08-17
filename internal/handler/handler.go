package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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
func (h *Handler) InitRouter(app *gin.Engine) *gin.Engine {
	auth := app.Group("/auth")
	auth.POST("/sign-up", h.signUp)
	auth.POST("/sign-in", h.signIn)
	auth.POST("/refresh-token", h.refreshToken)
	auth.POST("/logout", h.logout)

	apiV1 := app.Group("/api", h.userIdentity)
	apiV1.GET("", func(c *gin.Context) {
		userID := c.GetString("UserID")
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Hi user #%v you are in the secured route...", userID),
		})
	})

	return app
}
