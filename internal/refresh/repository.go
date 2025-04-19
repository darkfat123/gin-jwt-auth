package refresh

import (
	"context"
	"database/sql"
	"errors"
	"gin-jwt-auth/pkg/logger"
	"gin-jwt-auth/pkg/utils"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

type RefreshRepository interface {
	Refresh(ctx context.Context, jti string, username string) (int, error)
	UpdateRefreshToken(ctx context.Context, token string, id int) error
}

type refreshRepository struct {
	db *sqlx.DB
}

func NewRefreshRepository(db *sqlx.DB) RefreshRepository {
	return &refreshRepository{db: db}
}

func (r *refreshRepository) Refresh(ctx context.Context, jti string, username string) (int, error) {
	query := `
		SELECT refresh_tokens.id
		FROM refresh_tokens 
		LEFT JOIN users ON users.id = refresh_tokens.user_id 
		WHERE users.username = $1 AND refresh_tokens.jti = $2 AND refresh_tokens.revoked = false
	`

	var id int
	err := r.db.GetContext(ctx, &id, query, username, jti)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("failed to query refresh token", err)
			return 0, nil
		}
		log.Println("failed to query refresh token", err)
		return 0, err
	}

	logger.Info("refresh token exists")
	return id, nil
}

func (r *refreshRepository) UpdateRefreshToken(ctx context.Context, token string, id int) error {

	claims, err := utils.DecodeJWT(token)
	if err != nil {
		log.Println(err)
		return errors.New("invalid token")
	}
	query := `
		UPDATE refresh_tokens 
		SET jti = $1, refresh_token = $2, expires_at = $3, updated_at = $4
		WHERE id = $5
	`
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

	now := time.Now()
	expTime := time.Unix(int64(expFloat), 0)

	result, err := r.db.ExecContext(ctx, query, jti, token, expTime, now, id)
	if err != nil {
		log.Println("failed to execute refresh token update", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("failed to get rows affected", err)
		return err
	}

	if rowsAffected == 0 {
		log.Println("no rows were updated - possibly invalid ID")
		return errors.New("data not found")
	}

	logger.Info("refresh token has been updated")
	return nil
}
