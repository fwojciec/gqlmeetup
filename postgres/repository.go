package postgres

import (
	"github.com/fwojciec/gqlmeetup"
	"github.com/jmoiron/sqlx"
)

// Repository implements the gqlmeetup.Repository interface.
type Repository struct {
	DB *sqlx.DB
}

var _ gqlmeetup.Repository = (*Repository)(nil)
