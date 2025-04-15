package router

import (
	"gin-jwt-auth/internal/login"
	"gin-jwt-auth/internal/register"
	"gin-jwt-auth/internal/user"
	"gin-jwt-auth/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func AddRoutes(r *gin.Engine, db *sqlx.DB) {
	//No need bearer token
	auth := r.Group("/auth")

	registerRepo := register.NewRegisterRepository(db)
	registerService := register.NewRegisterService(registerRepo)
	registerHandler := register.NewRegisterHandler(registerService)
	auth.POST("/register", registerHandler.RegisterUser)

	loginRepo := login.NewLoginRepository(db)
	loginService := login.NewLoginService(loginRepo)
	loginHandler := login.NewLoginHandler(loginService)
	auth.POST("/login", loginHandler.LoginUser)

	//Need bearer token
	api := r.Group("/api")
	api.Use(utils.JWTAuthMiddleware())
	{
		userRepo := user.NewUserRepository(db)
		userService := user.NewUserService(userRepo)
		userHandler := user.NewUserHandler(userService)
		api.GET("/users/:id", userHandler.GetUserByID)
	}

}
