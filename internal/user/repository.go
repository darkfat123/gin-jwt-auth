package user

import (
	"context"
	"database/sql"
	"errors"
	"gin-jwt-auth/internal/user/dto"
	"gin-jwt-auth/pkg/logger"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id string) (*dto.UserResponse, error)
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserByID(ctx context.Context, id string) (*dto.UserResponse, error) {
	var user dto.UserResponse
	query := `SELECT id, username, email FROM users WHERE id = $1`

	err := r.db.GetContext(ctx, &user, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		logger.Error("failed to query user by id", zap.String("id", id), zap.Error(err))
		return nil, errors.New("internal server error")
	}

	return &user, nil
}
