package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your-secret-key")

func GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"iss": username,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
		"iat": time.Now().Unix(),
		"sub": "access_token",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string, username string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		return false, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expUnix, ok := claims["exp"].(float64)
		if !ok {
			return false, errors.New("invalid exp claim")
		}

		if int64(expUnix) < time.Now().Unix() {
			return false, errors.New("token expired")
		}

		if claims["iss"] != username {
			return false, errors.New("invalid issuer")
		}

		return true, nil
	}

	return false, errors.New("invalid token claims")
}
