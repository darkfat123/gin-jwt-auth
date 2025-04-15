package config

import (
	"os"

	"gin-jwt-auth/pkg/logger"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	ServerPort string
	Env        string
	JwtSecret  string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		logger.Error("No .env file found, using system environment variables")
	}

	return &Config{
		DBUser:     mustGetEnv("DB_USER"),
		DBPassword: mustGetEnv("DB_PASSWORD"),
		DBHost:     mustGetEnv("DB_HOST"),
		DBPort:     mustGetEnv("DB_PORT"),
		DBName:     mustGetEnv("DB_NAME"),
		ServerPort: mustGetEnv("SERVER_PORT"),
		Env:        mustGetEnv("ENVIRONMENT"),
		JwtSecret:  mustGetEnv("JWT_SECRET"),
	}

}

func mustGetEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		logger.Error("Missing required environment variable", zap.String("key", key))
	}
	return val
}
