package store

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type store struct {
	db *sql.DB
}

func NewStore(ctx context.Context, connString string) (*store, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}
	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}
	return &store{
		db: db,
	}, nil
}

// Do database operations here
