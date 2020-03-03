package postgres_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/fwojciec/gqlmeetup"
	"github.com/fwojciec/gqlmeetup/postgres"
	"github.com/jmoiron/sqlx"
)

func TestAuthorsListByAgentIDs(t *testing.T) {
	t.Parallel()
	tests := []struct {
		agentIDs []int64
		exp      []*gqlmeetup.Author
	}{
		{[]int64{}, []*gqlmeetup.Author{}},
		{[]int64{1}, []*gqlmeetup.Author{&testAuthor1}},
		{[]int64{1, 2}, []*gqlmeetup.Author{&testAuthor1, &testAuthor2}},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("%v", tc.agentIDs), func(t *testing.T) {
			t.Parallel()
			pgt.Runner(t, []string{"authors"}, func(t *testing.T, sdb *sqlx.DB) {
				ctx := context.Background()
				repo := &postgres.DataLoaderRepository{DB: sdb}
				res, err := repo.AuthorsListByAgentIDs(ctx, tc.agentIDs)
				ok(t, err)
				equals(t, tc.exp, res)
			})
		})
	}
}
