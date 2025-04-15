package register

import (
	"gin-jwt-auth/internal/register/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterHandler struct {
	service RegisterService
}

func NewRegisterHandler(service RegisterService) *RegisterHandler {
	return &RegisterHandler{service: service}
}

func (h *RegisterHandler) RegisterUser(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := h.service.RegisterUser(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"error": "Registered Successfully!"})
}
