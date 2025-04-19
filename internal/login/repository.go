package login

import (
	"context"
	"errors"
	"gin-jwt-auth/internal/login/dto"
	"gin-jwt-auth/model"
	"gin-jwt-auth/pkg/logger"
	"gin-jwt-auth/pkg/utils"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type LoginRepository interface {
	Login(ctx context.Context, req dto.LoginRequest) (*model.User, error)
	InsertRefreshToken(ctx context.Context, token string, user *model.User) error
}

type loginRepository struct {
	db *sqlx.DB
}

func NewLoginRepository(db *sqlx.DB) LoginRepository {
	return &loginRepository{db: db}
}

func (r *loginRepository) Login(ctx context.Context, req dto.LoginRequest) (*model.User, error) {
	var user model.User

	query := `SELECT id,username, password FROM users WHERE username = $1`
	err := r.db.GetContext(ctx, &user, query, req.Username)
	if err != nil {
		logger.Error("failed to fetch user", zap.String("username", req.Username), zap.Error(err))
		return nil, err
	}

	return &user, nil
}

func (r *loginRepository) InsertRefreshToken(ctx context.Context, token string, user *model.User) error {
	claims, err := utils.DecodeJWT(token)
	if err != nil {
		log.Println("Error:", err)
		return errors.New("unable to decode jwt")
	}
	username, ok := claims["iss"].(string)
	if !ok {
		log.Println("username is not a string")
		return errors.New("username is not a string")
	}
	log.Println(username)

	jti, ok := claims["jti"].(string)
	if !ok {
		log.Println("jti is not a string")
		return errors.New("jti is not a string")
	}
	log.Println(jti)

	expFloat, ok := claims["exp"].(float64)
	if !ok {
		log.Println("exp is not a number")
		return errors.New("exp is not a number")
	}

	expTime := time.Unix(int64(expFloat), 0)
	log.Println("exp time:", expTime)

	now := time.Now()
	query := `INSERT INTO refresh_tokens (user_id, jti, refresh_token,expires_at,revoked,created_at,updated_at) values ($1,$2,$3,$4,$5,$6,$7)`
	_, err = r.db.ExecContext(ctx, query, user.ID, jti, token, expTime, false, now, now)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}
