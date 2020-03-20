package postgres_test

import (
	"context"
	"testing"

	"github.com/fwojciec/gqlmeetup"
	"github.com/fwojciec/gqlmeetup/postgres"
	"github.com/jmoiron/sqlx"
)

func TestUserCreate(t *testing.T) {
	t.Parallel()
	pgt.Runner(t, []string{"users"}, func(t *testing.T, sdb *sqlx.DB) {
		repo := &postgres.Repository{DB: sdb}
		t.Run("create", func(t *testing.T) {
			err := repo.UserCreate(testUserCreate)
			ok(t, err)
			t.Run("assert user was created", func(t *testing.T) {
				t.Parallel()
				q := `SELECT * FROM users WHERE email = $1;`
				res := gqlmeetup.User{}
				_ = sdb.Get(&res, q, testUserCreate.Email)
				equals(t, testUserCreate, res)
			})
		})
	})
}

func TestUserGetByEmail(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		email string
		exp   *gqlmeetup.User
		err   error
	}{
		{"exists", testUser1.Email, &testUser1, nil},
		{"doesn't exist", "wrong", nil, gqlmeetup.ErrNotFound},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			pgt.Runner(t, []string{"users"}, func(t *testing.T, sdb *sqlx.DB) {
				repo := &postgres.Repository{DB: sdb}
				res, err := repo.UserGetByEmail(context.Background(), tc.email)
				equals(t, tc.err, err)
				equals(t, tc.exp, res)
			})
		})
	}
}
