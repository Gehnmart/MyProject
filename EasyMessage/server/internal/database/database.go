package database

import (
	"context"
	"database/sql"
	"server/internal/config"
	"time"

	_ "github.com/lib/pq"
)

func New(cfg config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DB.DSN)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, err
}
