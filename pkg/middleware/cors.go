package middleware

import (
	"os"
	"strings"

	"gin-jwt-auth/pkg/logger"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	originEnv := os.Getenv("ALLOWED_ORIGINS")
	if originEnv == "" {
		logger.Error("CORS: Cannot load environment.")
	}

	allowedOrigins := strings.Split(originEnv, ",")

	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		for _, allowedOrigin := range allowedOrigins {
			if origin == strings.TrimSpace(allowedOrigin) {
				c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
				break
			}
		}

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
