package service

import (
	"Avito/internal/models"
	"Avito/internal/repository"
)

type Segment interface {
	CreateSegment(slug string) (models.Segment, error)
	DeleteSegment(slug string) (models.Segment, error)
}

type UserSegment interface {
	GetActiveSegmentsByUser(user models.User) ([]models.Segment, error)
	ChangeSegments(user models.User, segmentsToAdd, segmentsToRemove []string) error
}

type Services struct {
	Segment
	UserSegment
}

func NewServices(repositories *repository.Repositories) *Services {
	return &Services{

		Segment:     NewSegmentService(repositories.Segment),
		UserSegment: NewUserSegmentService(repositories.UserSegments),
	}
}
