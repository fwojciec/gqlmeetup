package postgres_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/fwojciec/gqlmeetup"
	"github.com/fwojciec/gqlmeetup/postgres"
	"github.com/jmoiron/sqlx"
)

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
				repo := &postgres.DataLoaderRepository{DB: sdb}
				res, err := repo.AuthorListByAgentIDs(ctx, tc.agentIDs)
				ok(t, err)
				equals(t, tc.exp, res)
			})
		})
	}
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
				repo := &postgres.DataLoaderRepository{DB: sdb}
				res, err := repo.AgentListByIDs(ctx, tc.ids)
				ok(t, err)
				equals(t, tc.exp, res)
			})
		})
	}
}

func TestBookListByAuthorIDs(t *testing.T) {
	t.Parallel()
	tests := []struct {
		authorIDs []int64
		exp       []*gqlmeetup.Book
	}{
		{[]int64{}, []*gqlmeetup.Book{}},
		{[]int64{1}, []*gqlmeetup.Book{
			{ID: 1, Title: "Test Book 1", AuthorIDs: []int64{1}},
			{ID: 3, Title: "Test Book 3", AuthorIDs: []int64{1}},
		}},
		{[]int64{2}, []*gqlmeetup.Book{
			{ID: 2, Title: "Test Book 2", AuthorIDs: []int64{2}},
			{ID: 3, Title: "Test Book 3", AuthorIDs: []int64{2}},
		}},
		{[]int64{1, 2}, []*gqlmeetup.Book{
			{ID: 1, Title: "Test Book 1", AuthorIDs: []int64{1}},
			{ID: 2, Title: "Test Book 2", AuthorIDs: []int64{2}},
			{ID: 3, Title: "Test Book 3", AuthorIDs: []int64{1, 2}},
		}},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("%v", tc.authorIDs), func(t *testing.T) {
			t.Parallel()
			pgt.Runner(t, []string{"book_authors"}, func(t *testing.T, sdb *sqlx.DB) {
				ctx := context.Background()
				repo := &postgres.DataLoaderRepository{DB: sdb}
				res, err := repo.BookListByAuthorIDs(ctx, tc.authorIDs)
				for _, r := range res {
					t.Log(*r)
				}
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
				repo := &postgres.DataLoaderRepository{DB: sdb}
				res, err := repo.AuthorListByBookIDs(ctx, tc.bookIDs)
				for _, r := range res {
					t.Log(*r)
				}
				ok(t, err)
				equals(t, tc.exp, res)
			})
		})
	}
}
