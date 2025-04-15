package router

import (
	"gin-jwt-auth/internal/login"
	"gin-jwt-auth/internal/register"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func AddRoutes(r *gin.Engine, db *sqlx.DB) {
	api := r.Group("/api")

	// Register module
	registerRepo := register.NewRegisterRepository(db)
	registerService := register.NewRegisterService(registerRepo)
	registerHandler := register.NewRegisterHandler(registerService)
	api.POST("/register", registerHandler.RegisterUser)

	loginRepo := login.NewLoginRepository(db)
	loginService := login.NewLoginService(loginRepo)
	loginHandler := login.NewLoginHandler(loginService)
	api.POST("/login", loginHandler.LoginUser)
	// Add more modules here...
}
