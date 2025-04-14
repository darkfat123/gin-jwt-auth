package repository

import (
	"context"
	"gin-jwt-auth/internal/model"
	"gin-jwt-auth/pkg/utils"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

type RegisterRepository interface {
	Register(ctx context.Context, req model.RegisterRequest) error
}

type registerRepository struct {
	db *sqlx.DB
}

func NewRegisterRepository(db *sqlx.DB) RegisterRepository {
	return &registerRepository{db: db}
}

func (r *registerRepository) Register(ctx context.Context, req model.RegisterRequest) error {
	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}
	query := `INSERT INTO users (username, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`
	_, err = r.db.ExecContext(ctx, query, req.Username, req.Email, hashPassword, time.Now(), time.Now())
	if err != nil {
		log.Println("Register error:", err)
		return err
	}
	return nil
}
