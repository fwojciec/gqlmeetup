package postgres_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/fwojciec/gqlmeetup"
	"github.com/fwojciec/gqlmeetup/postgres"
	"github.com/jmoiron/sqlx"
)

func TestAgentCreate(t *testing.T) {
	t.Parallel()
	pgt.Runner(t, []string{"agents"}, func(t *testing.T, sdb *sqlx.DB) {
		ctx := context.Background()
		repo := &postgres.Repository{DB: sdb}
		t.Run("create", func(t *testing.T) {
			res, err := repo.AgentCreate(ctx, testAgentCreate)
			ok(t, err)
			equals(t, &testAgentCreate, res)
			t.Run("assert agent was created", func(t *testing.T) {
				t.Parallel()
				res, _ := repo.AgentGetByID(ctx, testAgentCreate.ID)
				equals(t, &testAgentCreate, res)
			})
		})
	})
}

func TestAgentDelete(t *testing.T) {
	t.Parallel()
	pgt.Runner(t, []string{"agents"}, func(t *testing.T, sdb *sqlx.DB) {
		ctx := context.Background()
		repo := &postgres.Repository{DB: sdb}
		t.Run("delete", func(t *testing.T) {
			res, err := repo.AgentDelete(ctx, testAgent1.ID)
			ok(t, err)
			equals(t, &testAgent1, res)
			t.Run("assert agent was deleted", func(t *testing.T) {
				t.Parallel()
				_, err := repo.AgentGetByID(ctx, testAgent1.ID)
				equals(t, gqlmeetup.ErrNotFound, err)
			})
		})
	})
}

func TestAgentGetByID(t *testing.T) {
	t.Parallel()
	pgt.Runner(t, []string{"agents"}, func(t *testing.T, sdb *sqlx.DB) {
		repo := &postgres.Repository{DB: sdb}
		res, err := repo.AgentGetByID(context.Background(), 1)
		ok(t, err)
		equals(t, &testAgent1, res)
	})
}

func TestAgentList(t *testing.T) {
	t.Parallel()
	pgt.Runner(t, []string{"agents"}, func(t *testing.T, sdb *sqlx.DB) {
		repo := &postgres.Repository{DB: sdb}
		res, err := repo.AgentList(context.Background())
		ok(t, err)
		equals(t, []*gqlmeetup.Agent{&testAgent1, &testAgent2}, res)
	})
}

func TestAgentListByIDs(t *testing.T) {
	t.Parallel()
	tests := []struct {
		ids []int64
		exp []*gqlmeetup.Agent
	}{
		{[]int64{}, []*gqlmeetup.Agent{}},
		{[]int64{1}, []*gqlmeetup.Agent{&testAgent1}},
		{[]int64{1, 2}, []*gqlmeetup.Agent{&testAgent1, &testAgent2}},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("%v", tc.ids), func(t *testing.T) {
			t.Parallel()
			pgt.Runner(t, []string{"agents"}, func(t *testing.T, sdb *sqlx.DB) {
				ctx := context.Background()
				repo := &postgres.Repository{DB: sdb}
				res, err := repo.AgentListByIDs(ctx, tc.ids)
				ok(t, err)
				equals(t, tc.exp, res)
			})
		})
	}
}

func TestAgentUpdate(t *testing.T) {
	t.Parallel()
	pgt.Runner(t, []string{"agents"}, func(t *testing.T, sdb *sqlx.DB) {
		ctx := context.Background()
		repo := &postgres.Repository{DB: sdb}
		t.Run("update", func(t *testing.T) {
			res, err := repo.AgentUpdate(ctx, testAgent1.ID, testAgentUpdate)
			ok(t, err)
			equals(t, &testAgentUpdate, res)
			t.Run("assert agent was updated", func(t *testing.T) {
				t.Parallel()
				res, _ := repo.AgentGetByID(ctx, testAgent1.ID)
				equals(t, &testAgentUpdate, res)
			})
		})
	})
}
