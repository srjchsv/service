package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/srjchsv/service/internal/repository"
)

// SignUp performs users sign-up on handlers level
func (h *Handler) signUp(c *gin.Context) {
	var input repository.User

	if err := c.BindJSON(&input); err != nil {
		newErrorReponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorReponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// SignIn performs users sign-in on handlers level
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorReponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorReponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	
	c.SetCookie(
		"access_token",
		token,
		60*60*24,
		"/",
		"/",
		true,
		true,
	)
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
