package login_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLoginHandler(t *testing.T) {
	// Set Gin to Test mode
	gin.SetMode(gin.TestMode)

	// Create a new Gin router
	router := gin.Default()

	// Define your API route handler
	router.POST("/auth/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Login Successfully!"})
	})

	// Create a test request
	req, _ := http.NewRequest("POST", "/auth/login", nil)
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert the response status
	assert.Equal(t, http.StatusOK, w.Code)

	// Assert the response body
	assert.JSONEq(t, `{"message": "Login Successfully!"}`, w.Body.String())
}
