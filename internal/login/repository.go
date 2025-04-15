package login

import (
	"context"
	"fmt"
	"gin-jwt-auth/internal/login/dto"
	"gin-jwt-auth/model"
	"gin-jwt-auth/pkg/logger"
	"gin-jwt-auth/pkg/utils"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type LoginRepository interface {
	Login(ctx context.Context, req dto.LoginRequest) error
}

type loginRepository struct {
	db *sqlx.DB
}

func NewLoginRepository(db *sqlx.DB) LoginRepository {
	return &loginRepository{db: db}
}

func (r *loginRepository) Login(ctx context.Context, req dto.LoginRequest) error {
	var user model.User

	query := `SELECT username, password FROM users WHERE username = $1`
	err := r.db.GetContext(ctx, &user, query, req.Username)
	if err != nil {
		logger.Info("failed to fetch user", zap.String("username", req.Username))
		return fmt.Errorf("invalid username or password")
	}

	isValid := utils.CheckPasswordHash(req.Password, user.Password)
	if !isValid {
		logger.Info("invalid password attempt", zap.String("username", req.Username))
		return fmt.Errorf("invalid username or password")
	}

	logger.Info("user logged in successfully", zap.String("username", user.Username))

	return nil
}
