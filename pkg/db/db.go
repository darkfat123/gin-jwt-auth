package db

import (
	"gin-jwt-auth/pkg/logger"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDB(dsn string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		logger.Error("DB connection failed")
		return nil
	}
	return db
}
