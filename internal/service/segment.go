package service

import (
	"Avito/internal/models"
	"Avito/internal/repository"
	"Avito/internal/repository/db"
	"errors"
)

type SegmentService struct {
	segmentRepo repository.Segment
}

func NewSegmentService(segmentRepo repository.Segment) *SegmentService {
	return &SegmentService{segmentRepo: segmentRepo}
}

func (s *SegmentService) CreateSegment(slug string) (models.Segment, error) {
	newSegment := models.Segment{Slug: slug}
	var err error
	newSegment.ID, err = s.segmentRepo.CreateSegment(slug)
	if err != nil {
		if errors.Is(err, db.ErrAlreadyExists) {
			return models.Segment{}, ErrSegmentAlreadyExists
		}
		return models.Segment{}, ErrCannotCreateSegment
	}
	return newSegment, nil
}

func (s *SegmentService) DeleteSegment(slug string) (models.Segment, error) {
	deletedSegment := models.Segment{Slug: slug}
	var err error
	deletedSegment.ID, err = s.segmentRepo.DeleteSegment(slug)
	if err != nil {
		if errors.Is(err, db.ErrNothingToDelete) {
			return models.Segment{}, ErrSegmentDoesNotExist
		}
		return models.Segment{}, ErrCannotDeleteSegment
	}
	return deletedSegment, nil
}
