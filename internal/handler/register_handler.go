package handler

import (
	"gin-jwt-auth/internal/model"
	"gin-jwt-auth/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterHandler struct {
	service service.RegisterService
}

func NewRegisterHandler(service service.RegisterService) *RegisterHandler {
	return &RegisterHandler{service: service}
}

func (h *RegisterHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/register", h.RegisterUser)
}

func (h *RegisterHandler) RegisterUser(c *gin.Context) {
	var req model.RegisterRequest
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
