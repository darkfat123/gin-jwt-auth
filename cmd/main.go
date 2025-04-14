package main

import (
	"fmt"
	"gin-jwt-auth/config"
	"gin-jwt-auth/internal/db"

	"gin-jwt-auth/internal/handler"
	"gin-jwt-auth/internal/repository"
	"gin-jwt-auth/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	conf := config.LoadConfig()

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.DBHost, conf.DBPort, conf.DBUser, conf.DBPassword, conf.DBName,
	)
	conn := db.ConnectDB(dsn)

	userRepo := repository.NewRegisterRepository(conn)
	userSvc := service.NewRegisterService(userRepo)
	userHandler := handler.NewRegisterHandler(userSvc)

	r := gin.Default()
	api := r.Group("/api")
	userHandler.RegisterRoutes(api)

	r.Run(":" + conf.ServerPort)
}
