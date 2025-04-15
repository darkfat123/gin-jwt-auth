package login

import (
	"context"
	"gin-jwt-auth/internal/login/dto"
	"gin-jwt-auth/model"
	"gin-jwt-auth/pkg/logger"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type LoginRepository interface {
	Login(ctx context.Context, req dto.LoginRequest) (*model.User, error)
}

type loginRepository struct {
	db *sqlx.DB
}

func NewLoginRepository(db *sqlx.DB) LoginRepository {
	return &loginRepository{db: db}
}

func (r *loginRepository) Login(ctx context.Context, req dto.LoginRequest) (*model.User, error) {
	var user model.User

	query := `SELECT username, password FROM users WHERE username = $1`
	err := r.db.GetContext(ctx, &user, query, req.Username)
	if err != nil {
		logger.Error("failed to fetch user", zap.String("username", req.Username), zap.Error(err))
		return nil, err
	}

	return &user, nil
}
