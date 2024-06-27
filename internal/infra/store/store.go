package store

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	psql "github.com/karthikraobr/bot/internal/querier"
	_ "github.com/lib/pq"
)

type store struct {
	querier psql.Querier
}

func NewStore(ctx context.Context, pool *pgxpool.Pool) *store {
	return &store{
		querier: psql.New(pool),
	}
}

func (s *store) InsertReview(ctx context.Context, arg psql.InsertReviewParams) (int32, error) {
	return s.querier.InsertReview(ctx, arg)
}
