package db

import (
	"errors"
)

var (
	ErrAlreadyExists   = errors.New("already exists")
	ErrNothingToDelete = errors.New("nothing to delete")
	ErrNotFound        = errors.New("not found")
)
