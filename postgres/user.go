package postgres

import (
	"context"
	"database/sql"

	"github.com/fwojciec/gqlmeetup"
)

const userCreateQuery = `
INSERT INTO users (email, password, admin) VALUES ($1, $2, $3);`

// UserCreate creates a user.
func (r *Repository) UserCreate(data gqlmeetup.User) error {
	if _, err := r.DB.Exec(userCreateQuery, &data.Email, &data.Password, &data.Admin); err != nil {
		return err
	}
	return nil
}

const userGetByEmailQuery = `
SELECT * FROM users WHERE email = $1`

// UserGetByEmail gets a user by email address.
func (r *Repository) UserGetByEmail(ctx context.Context, email string) (*gqlmeetup.User, error) {
	res := &gqlmeetup.User{}
	if err := r.DB.GetContext(ctx, res, userGetByEmailQuery, email); err != nil {
		if err == sql.ErrNoRows {
			return nil, gqlmeetup.ErrNotFound
		}
		return nil, err
	}
	return res, nil
}
