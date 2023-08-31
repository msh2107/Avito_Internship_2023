package service

import "errors"

var (
	ErrSegmentAlreadyExists = errors.New("segment already exists")
	ErrSegmentDoesNotExist  = errors.New("segment does not exist")
	ErrNoActiveSegments     = errors.New("no active segments")

	ErrCannotCreateSegment     = errors.New("cannot create segment")
	ErrCannotDeleteSegment     = errors.New("cannot delete segment")
	ErrCannotGetActiveSegments = errors.New("cannot get active segments")
)
