// internal/infra/db/postgres.go

package db

import (
	"context"
	"fmt"
	"time"

	"github.com/rayfiyo/yamabiko/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresHistoryRepo struct {
	pool *pgxpool.Pool
}

func NewPostgresHistoryRepo(pool *pgxpool.Pool) domain.HistoryRepository {
	return &postgresHistoryRepo{pool: pool}
}

// Save inserts a new shout history record into the DB.
func (r *postgresHistoryRepo) Save(h *domain.ShoutHistory) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
    INSERT INTO shout_history (
      voice_text,
      response1,
      response2,
      response3,
      response4,
      response5,
      response6,
      created_at
    ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    RETURNING id;
    `
	err := r.pool.QueryRow(ctx, query,
		h.Voice,
		h.Response1,
		h.Response2,
		h.Response3,
		h.Response4,
		h.Response5,
		h.Response6,
		h.CreatedAt,
	).Scan(&h.ID)
	if err != nil {
		return fmt.Errorf("postgresHistoryRepo.Save: %w", err)
	}

	return nil
}

// FindAll returns all shout history, newest first
func (r *postgresHistoryRepo) FindAll() ([]*domain.ShoutHistory, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
    SELECT
      id,
      voice_text,
      response1,
      response2,
      response3,
      response4,
      response5,
      response6,
      created_at
    FROM shout_history
    ORDER BY created_at DESC;
    `
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("postgresHistoryRepo.FindAll: %w", err)
	}
	defer rows.Close()

	var results []*domain.ShoutHistory
	for rows.Next() {
		var h domain.ShoutHistory
		err := rows.Scan(
			&h.ID,
			&h.Voice,
			&h.Response1,
			&h.Response2,
			&h.Response3,
			&h.Response4,
			&h.Response5,
			&h.Response6,
			&h.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, &h)
	}
	return results, nil
}
