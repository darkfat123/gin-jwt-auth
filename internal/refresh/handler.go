package refresh

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RefreshHandler struct {
	service RefreshService
}

func NewRefreshHandler(service RefreshService) *RefreshHandler {
	return &RefreshHandler{service: service}
}

func (h *RefreshHandler) RefreshUser(c *gin.Context) {
	reqRefresh, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	tokens, err := h.service.RefreshUser(c.Request.Context(), reqRefresh)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("refresh_token", tokens.RefreshToken, 1*86400, "/auth/refresh", "localhost", true, true)

	c.JSON(http.StatusOK, gin.H{"message": "Refreshed Successfully!", "token": tokens})
}
