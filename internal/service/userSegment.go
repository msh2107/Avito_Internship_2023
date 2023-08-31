package service

import (
	"Avito/internal/models"
	"Avito/internal/repository"
	"Avito/internal/repository/db"
	"errors"
)

type UserSegmentService struct {
	userSegmentRepo repository.UserSegments
}

func NewUserSegmentService(userSegmentRepo repository.UserSegments) *UserSegmentService {
	return &UserSegmentService{userSegmentRepo: userSegmentRepo}
}

func (s *UserSegmentService) GetActiveSegmentsByUser(user models.User) ([]models.Segment, error) {
	segments, err := s.userSegmentRepo.GetActiveSegmentsByUser(user.ID)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil, ErrNoActiveSegments
		}
		return nil, ErrCannotGetActiveSegments
	}
	return segments, nil
}

func (s *UserSegmentService) ChangeSegments(user models.User, segmentsToAdd, segmentsToRemove []string) error {
	err := s.userSegmentRepo.ChangeActiveSegments(user.ID, segmentsToAdd, segmentsToRemove)
	if err != nil {
		return err
	}
	return nil
}
