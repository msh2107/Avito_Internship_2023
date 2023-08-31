package repository

import (
	"Avito/internal/models"
	"Avito/internal/repository/db"
	"Avito/pkg/db/postgresql"
)

type Segment interface {
	CreateSegment(slug string) (int, error)
	DeleteSegment(slug string) (int, error)
}

type UserSegments interface {
	GetActiveSegmentsByUser(userID int) ([]models.Segment, error)
	ChangeActiveSegments(userID int, slugsToAdd, slugsToDelete []string) error
}

type Repositories struct {
	Segment
	UserSegments
}

func NewRepositories(pg *postgresql.Postgres) *Repositories {
	return &Repositories{
		Segment:      db.NewSegmentRepository(pg),
		UserSegments: db.NewUserSegmentRepository(pg),
	}
}
