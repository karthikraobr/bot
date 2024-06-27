package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// NewDB creates a new connection to the database
func NewDB(ctx context.Context, dsn string) (*pgxpool.Pool, error) {

	pgxpoolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	conn, err := pgxpool.NewWithConfig(ctx, pgxpoolConfig)
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
	}

	// try pinging the database to check if the connection is alive
	// retry 3 times with a 1 second delay
	for i := 0; i < 3; i++ {
		err = conn.Ping(ctx)
		if err == nil {
			break
		}

		if i == 2 {
			return nil, fmt.Errorf("could not ping database: %w", err)
		}

		time.Sleep(time.Second)
	}

	return conn, nil
}

// CloseDB closes the connection to the database
func CloseDB(ctx context.Context, conn *pgx.Conn) error {
	if conn != nil {
		err := conn.Close(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
