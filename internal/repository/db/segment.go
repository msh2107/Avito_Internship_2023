package db

import (
	"Avito/pkg/db/postgresql"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"strings"
)

type SegmentRepository struct {
	*postgresql.Postgres
}

func NewSegmentRepository(pg *postgresql.Postgres) *SegmentRepository {
	return &SegmentRepository{Postgres: pg}
}

func (r *SegmentRepository) CreateSegment(slug string) (int, error) {
	q := `INSERT INTO segments
    	(slug) 
	VALUES 
	    ($1) 
	RETURNING id`
	var id int
	if err := r.Pool.QueryRow(context.Background(), q, slug).Scan(&id); err != nil {
		fmt.Println(err)
		if strings.Contains(err.Error(), "SQLSTATE 23505") {
			return 0, ErrAlreadyExists
		}
		return 0, fmt.Errorf("SegmentRepository.CreateSegment - r.Pool.QueryRow: %v", err)
	}
	return id, nil
}

func (r *SegmentRepository) DeleteSegment(slug string) (int, error) {
	q := `DELETE FROM segments
    	WHERE
    	    slug = $1
    	    RETURNING id`
	var id int
	if err := r.Pool.QueryRow(context.Background(), q, slug).Scan(&id); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, ErrNothingToDelete
		}
		return 0, fmt.Errorf("SegmentRepository.DeleteSegment - r.Pool.QueryRow: %v", err)
	}

	return id, nil
}
