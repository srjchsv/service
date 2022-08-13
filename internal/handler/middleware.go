package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userID"
)

// userIdentity is a authentification middleware
func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)

	if header == "" {
		newErrorReponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}
	headerParts := strings.Split(string(header), " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorReponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		newErrorReponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	userID, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorReponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(userCtx, userID)
}
