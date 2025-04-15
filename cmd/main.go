package main

import (
	"fmt"
	"gin-jwt-auth/config"
	"gin-jwt-auth/pkg/db"
	"gin-jwt-auth/pkg/logger"
	"gin-jwt-auth/pkg/router"

	"github.com/gin-gonic/gin"
)

func main() {
	conf := config.LoadConfig()
	logger.Init()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.DBHost, conf.DBPort, conf.DBUser, conf.DBPassword, conf.DBName,
	)
	conn := db.ConnectDB(dsn)

	r := gin.Default()
	router.AddRoutes(r, conn)

	r.Run(":" + conf.ServerPort)
}
