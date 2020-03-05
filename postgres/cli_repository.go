package postgres

import (
	"github.com/fwojciec/gqlmeetup"
	"github.com/jmoiron/sqlx"
)

// CLIRepository implements the gqlmeetup.CLIRepository interface.
type CLIRepository struct {
	DB *sqlx.DB
}

var _ gqlmeetup.CLIRepository = (*CLIRepository)(nil)

const userCreateQuery = `
INSERT INTO users (email, password, admin) VALUES ($1, $2, $3);`

// UserCreate creates a user.
func (r *CLIRepository) UserCreate(data gqlmeetup.User) error {
	if _, err := r.DB.Exec(userCreateQuery, &data.Email, &data.Password, &data.Admin); err != nil {
		return err
	}
	return nil
}
