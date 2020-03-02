package gqlmeetup

import (
	"errors"
)

var (
	// ErrNotFound is returned when a queried resource was not found.
	ErrNotFound = errors.New("not found")
)
