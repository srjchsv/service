package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/srjchsv/service/internal/repository"
)

const (
	salt             = "fsfdf34444dijisjdfjdfi"
	signingKey       = "1@#edwdDSD$$"
	tokenTTL         = 15 * time.Minute
	cookieAgeSignIn  = int(tokenTTL / time.Second)
	cookieAgeRefresh = 60 * 5
)

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

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

// SignIn performs users sign-in on handlers level
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorReponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorReponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userID, err := h.services.Authorization.ParseToken(token)
	if err != nil {
		newErrorReponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.SetCookie(
		"access_token",
		token,
		cookieAgeSignIn,
		"/",
		"/",
		true,
		true,
	)
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
		"id":    userID,
	})
}

func (h *Handler) refreshToken(c *gin.Context) {
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

	newToken, err := h.services.Authorization.RefreshToken(token, userID)
	if err != nil {
		newErrorReponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.SetCookie(
		"access_token",
		newToken,
		cookieAgeRefresh,
		"/",
		"/",
		true,
		true,
	)
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": newToken,
	})

}

func (h *Handler) logout(c *gin.Context) {
	token, err := c.Cookie("access_token")
	if err != nil {
		if err == http.ErrNoCookie {
			newErrorReponse(c, http.StatusUnauthorized, "no cookie access_token")
			return
		}
		newErrorReponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	_, err = h.services.Authorization.ParseToken(token)
	if err != nil {
		newErrorReponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.SetCookie(
		"access_token",
		"",
		-1,
		"/",
		"/",
		true,
		true,
	)
	c.JSON(http.StatusOK, map[string]interface{}{
		"logout": "success",
	})
}
