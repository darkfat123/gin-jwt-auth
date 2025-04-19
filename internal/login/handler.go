package login

import (
	"gin-jwt-auth/internal/login/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	service LoginService
}

func NewLoginHandler(service LoginService) *LoginHandler {
	return &LoginHandler{service: service}
}

func (h *LoginHandler) LoginUser(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	token, user, err := h.service.LoginUser(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.SetCookie("refresh_token", token.RefreshToken, 86400, "/auth/refresh", "localhost", true, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Successfully!",
		"token":   token,
		"id":      user.ID,
	})
}
