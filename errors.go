package gqlmeetup

import (
	"errors"
)

var (
	// ErrInvalid is returned when a params or input values are invalid.
	ErrInvalid = errors.New("not found")

	// ErrNotFound is returned when a queried resource was not found.
	ErrNotFound = errors.New("not found")
)
