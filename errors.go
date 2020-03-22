package gqlmeetup

import (
	"errors"
)

var (
	// ErrInvalid is returned when a params or input values are invalid.
	ErrInvalid = errors.New("invalid")

	// ErrNotFound is returned when a queried resource was not found.
	ErrNotFound = errors.New("not found")

	// ErrPwdCheck is returned when password verification has failed.
	ErrPwdCheck = errors.New("password check failed")

	// ErrUnauthorized is return on unauthorized access attempt.
	ErrUnauthorized = errors.New("unauthorized")
)
