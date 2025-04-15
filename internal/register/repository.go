package register

import (
	"context"
	"errors"
	"gin-jwt-auth/internal/register/dto"
	"gin-jwt-auth/pkg/logger"
	"time"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type RegisterRepository interface {
	Register(ctx context.Context, req dto.RegisterRequest) error
	CheckDuplicate(ctx context.Context, req dto.RegisterRequest) error
}

type registerRepository struct {
	db *sqlx.DB
}

func NewRegisterRepository(db *sqlx.DB) RegisterRepository {
	return &registerRepository{db: db}
}

func (r *registerRepository) Register(ctx context.Context, req dto.RegisterRequest) error {
	now := time.Now()
	insertUser := `INSERT INTO users (username, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.ExecContext(ctx, insertUser, req.Username, req.Email, req.Password, now, now)
	if err != nil {
		logger.Error("failed to insert user", zap.String("username", req.Username), zap.String("email", req.Email), zap.Error(err))
		return err
	}

	return nil
}

func (r *registerRepository) CheckDuplicate(ctx context.Context, req dto.RegisterRequest) error {
	var count int
	duplicate := `SELECT COUNT(*) FROM users WHERE username = $1 OR email = $2`
	err := r.db.GetContext(ctx, &count, duplicate, req.Username, req.Email)
	if err != nil {
		logger.Error("failed to check username/email existence", zap.String("username", req.Username), zap.String("email", req.Email), zap.Error(err))
		return errors.New("failed to check username/email existence")
	}

	if count > 0 {
		logger.Error("username or email already exists", zap.String("username", req.Username), zap.String("email", req.Email))
		return errors.New("username or email already exists")
	}

	return nil
}
