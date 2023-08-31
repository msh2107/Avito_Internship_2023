package db

import (
	"Avito/internal/models"
	"Avito/pkg/db/postgresql"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"strings"
)

type UserSegmentRepository struct {
	*postgresql.Postgres
}

func NewUserSegmentRepository(pg *postgresql.Postgres) *UserSegmentRepository {
	return &UserSegmentRepository{Postgres: pg}
}

func (r *UserSegmentRepository) GetActiveSegmentsByUser(userID int) ([]models.Segment, error) {
	q := `SELECT id, slug FROM segments INNER JOIN user_segments ON segments.id = user_segments.segment_id WHERE user_id = $1;`
	rows, err := r.Pool.Query(context.Background(), q, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("UserSegmentRepository.GetActiveSegmentsByUser - r.Pool.Query: %v", err)
	}
	var segments []models.Segment
	for rows.Next() {
		var sgm models.Segment
		err = rows.Scan(&sgm.ID, &sgm.Slug)
		if err != nil {
			return nil, fmt.Errorf("UserSegmentRepository.GetActiveSegmentsByUser - rows.Scan: %v", err)
		}
		segments = append(segments, sgm)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	if len(segments) == 0 {
		return nil, ErrNotFound
	}
	return segments, nil
}

func (r *UserSegmentRepository) ChangeActiveSegments(userID int, slugsToAdd, slugsToDelete []string) error {
	tx, err := r.Pool.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("UserSegmentRepository.ChangeActiveSegments - r.Pool.Begin: %v", err)
	}
	defer func() {
		_ = tx.Rollback(context.Background())
	}()

	q := `INSERT  INTO users (id) VALUES ($1) ON CONFLICT DO NOTHING`
	_, err = tx.Exec(context.Background(), q, userID)
	if err != nil {
		return fmt.Errorf("UserRepository.ChangeActiveSegments - tx.Exec: %v", err)
	}

	addQ := `INSERT INTO user_segments (user_id, segment_id) SELECT users.id, segments.id FROM users, segments WHERE users.id = $1 AND segments.slug = $2;`

	for _, slug := range slugsToAdd {
		_, err = tx.Exec(context.Background(), addQ, userID, slug)
		if err != nil {
			if strings.Contains(err.Error(), "SQLSTATE 23505") {
				return fmt.Errorf("segment %s is already added", slug)
			}

			return fmt.Errorf("UserSegmentRepository.ChangeActiveSegments - tx.Exec: %v", err)
		}
	}

	deleteQ := `DELETE FROM user_segments WHERE user_id = $1 AND segment_id = (SELECT id FROM segments WHERE slug = $2);`

	for _, slug := range slugsToDelete {
		_, err = tx.Exec(context.Background(), deleteQ, userID, slug)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {

				return fmt.Errorf("no such slug as %s", slug)
			}
			return fmt.Errorf("UserSegmentRepository.ChangeActiveSegments - tx.Exec: %v", err)
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return fmt.Errorf("UserSegmentRepository.ChangeActiveSegments - tx.Commit: %v", err)
	}

	return nil
}
