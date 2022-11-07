package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	userCtx = "UserID"
)

// userIdentity is a authentification middleware
func (h *Handler) userIdentity(c *gin.Context) {
	token, err := c.Cookie("access_token")
	if err != nil {
		if err == http.ErrNoCookie {
			newErrorReponse(c, http.StatusUnauthorized, "no cookie access_token")
			return
		}
		newErrorReponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	userID, err := h.services.Authorization.ParseToken(token)
	if err != nil {
		newErrorReponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(userCtx, userID)
}
