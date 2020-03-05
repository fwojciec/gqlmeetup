package postgres_test

import (
	"testing"

	"github.com/fwojciec/gqlmeetup"
	"github.com/fwojciec/gqlmeetup/postgres"
	"github.com/jmoiron/sqlx"
)

func TestUserCreate(t *testing.T) {
	t.Parallel()
	pgt.Runner(t, []string{"users"}, func(t *testing.T, sdb *sqlx.DB) {
		repo := &postgres.CLIRepository{DB: sdb}
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
