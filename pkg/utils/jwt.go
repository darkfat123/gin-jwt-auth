package utils

import (
	"gin-jwt-auth/config"
	"gin-jwt-auth/model"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var conf = config.LoadConfig()

var jwtSecret = []byte(conf.JwtSecret)

func GenerateJWT(username string) (*model.TokenPair, error) {
	var pair model.TokenPair
	accessClaims := jwt.MapClaims{
		"iss": username,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
		"iat": time.Now().Unix(),
		"sub": "access_token",
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)

	accessString, err := accessToken.SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}

	refreshClaims := jwt.MapClaims{
		"iss": username,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
		"iat": time.Now().Unix(),
		"sub": "refresh_token",
		"jti": generateUUID(),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	refreshString, err := refreshToken.SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}

	pair = model.TokenPair{
		AccessToken:  accessString,
		RefreshToken: refreshString,
	}

	return &pair, nil
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(strings.ToLower(authHeader), "bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtSecret, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("user", claims["iss"])
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
		}
	}
}

func generateUUID() uuid.UUID {
	return uuid.New()
}
