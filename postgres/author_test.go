package postgres_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/fwojciec/gqlmeetup"
	"github.com/fwojciec/gqlmeetup/postgres"
	"github.com/jmoiron/sqlx"
)

func TestAuthorCreate(t *testing.T) {
	t.Parallel()
	pgt.Runner(t, []string{"authors"}, func(t *testing.T, sdb *sqlx.DB) {
		ctx := context.Background()
		repo := &postgres.Repository{DB: sdb}
		t.Run("create", func(t *testing.T) {
			res, err := repo.AuthorCreate(ctx, testAuthorCreate)
			ok(t, err)
			equals(t, &testAuthorCreate, res)
			t.Run("assert agent was created", func(t *testing.T) {
				t.Parallel()
				res, _ := repo.AuthorGetByID(ctx, testAuthorCreate.ID)
				equals(t, &testAuthorCreate, res)
			})
		})
	})
}

func TestAuthorDelete(t *testing.T) {
	t.Parallel()
	pgt.Runner(t, []string{"authors"}, func(t *testing.T, sdb *sqlx.DB) {
		ctx := context.Background()
		repo := &postgres.Repository{DB: sdb}
		t.Run("delete", func(t *testing.T) {
			res, err := repo.AuthorDelete(ctx, testAuthor1.ID)
			ok(t, err)
			equals(t, &testAuthor1, res)
			t.Run("assert agent was deleted", func(t *testing.T) {
				t.Parallel()
				_, err := repo.AuthorGetByID(ctx, testAuthor1.ID)
				equals(t, gqlmeetup.ErrNotFound, err)
			})
		})
	})
}

func TestAuthorGetByID(t *testing.T) {
	t.Parallel()
	pgt.Runner(t, []string{"authors"}, func(t *testing.T, sdb *sqlx.DB) {
		repo := &postgres.Repository{DB: sdb}
		res, err := repo.AuthorGetByID(context.Background(), 1)
		ok(t, err)
		equals(t, &testAuthor1, res)
	})
}

func TestAuthorList(t *testing.T) {
	t.Parallel()
	pgt.Runner(t, []string{"authors"}, func(t *testing.T, sdb *sqlx.DB) {
		repo := &postgres.Repository{DB: sdb}
		res, err := repo.AuthorList(context.Background())
		ok(t, err)
		equals(t, []*gqlmeetup.Author{&testAuthor1, &testAuthor2}, res)
	})
}

func TestAuthorListByAgentIDs(t *testing.T) {
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
				repo := &postgres.Repository{DB: sdb}
				res, err := repo.AuthorListByAgentIDs(ctx, tc.agentIDs)
				ok(t, err)
				equals(t, tc.exp, res)
			})
		})
	}
}

func TestAuthorListByBookIDs(t *testing.T) {
	t.Parallel()
	tests := []struct {
		bookIDs []int64
		exp     []*gqlmeetup.Author
	}{
		{[]int64{}, []*gqlmeetup.Author{}},
		{[]int64{1}, []*gqlmeetup.Author{
			{ID: 1, Name: "Test Author 1", AgentID: 1, BookIDs: []int64{1}},
		}},
		{[]int64{2}, []*gqlmeetup.Author{
			{ID: 2, Name: "Test Author 2", AgentID: 2, BookIDs: []int64{2}},
		}},
		{[]int64{1, 2}, []*gqlmeetup.Author{
			{ID: 1, Name: "Test Author 1", AgentID: 1, BookIDs: []int64{1}},
			{ID: 2, Name: "Test Author 2", AgentID: 2, BookIDs: []int64{2}},
		}},
		{[]int64{1, 2, 3}, []*gqlmeetup.Author{
			{ID: 1, Name: "Test Author 1", AgentID: 1, BookIDs: []int64{1, 3}},
			{ID: 2, Name: "Test Author 2", AgentID: 2, BookIDs: []int64{2, 3}},
		}},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("%v", tc.bookIDs), func(t *testing.T) {
			t.Parallel()
			pgt.Runner(t, []string{"book_authors"}, func(t *testing.T, sdb *sqlx.DB) {
				ctx := context.Background()
				repo := &postgres.Repository{DB: sdb}
				res, err := repo.AuthorListByBookIDs(ctx, tc.bookIDs)
				ok(t, err)
				equals(t, tc.exp, res)
			})
		})
	}
}

func TestAuthorUpdate(t *testing.T) {
	t.Parallel()
	pgt.Runner(t, []string{"authors"}, func(t *testing.T, sdb *sqlx.DB) {
		ctx := context.Background()
		repo := &postgres.Repository{DB: sdb}
		t.Run("update", func(t *testing.T) {
			res, err := repo.AuthorUpdate(ctx, testAuthor1.ID, testAuthorUpdate)
			ok(t, err)
			equals(t, &testAuthorUpdate, res)
			t.Run("assert agent was updated", func(t *testing.T) {
				t.Parallel()
				res, _ := repo.AuthorGetByID(ctx, testAuthor1.ID)
				equals(t, &testAuthorUpdate, res)
			})
		})
	})
}
