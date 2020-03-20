package postgres_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/fwojciec/gqlmeetup"
	"github.com/fwojciec/gqlmeetup/postgres"
	"github.com/jmoiron/sqlx"
)

func TestBookCreate(t *testing.T) {
	t.Parallel()
	pgt.Runner(t, []string{"book_authors"}, func(t *testing.T, sdb *sqlx.DB) {
		ctx := context.Background()
		repo := &postgres.Repository{DB: sdb}
		t.Run("create", func(t *testing.T) {
			res, err := repo.BookCreate(ctx, testBookCreate, []int64{1, 2})
			ok(t, err)
			equals(t, &testBookCreate, res)
			t.Run("confirm create", func(t *testing.T) {
				t.Parallel()
				res, _ := repo.BookGetByID(ctx, testBookCreate.ID)
				equals(t, &testBookCreate, res)
			})
			t.Run("confirm author associations", func(t *testing.T) {
				t.Parallel()
				q := `SELECT author_id FROM book_authors WHERE book_id = $1;`
				authorIDs := make([]int64, 0)
				_ = sdb.SelectContext(ctx, &authorIDs, q, testBookCreate.ID)
				equals(t, []int64{1, 2}, authorIDs)
			})
		})
	})
}

func TestBookDelete(t *testing.T) {
	t.Parallel()
	pgt.Runner(t, []string{"book_authors"}, func(t *testing.T, sdb *sqlx.DB) {
		ctx := context.Background()
		repo := &postgres.Repository{DB: sdb}
		t.Run("delete", func(t *testing.T) {
			res, err := repo.BookDelete(ctx, testBook3.ID)
			ok(t, err)
			equals(t, &testBook3, res)
			t.Run("confirm author association delete", func(t *testing.T) {
				t.Parallel()
				q := `SELECT author_id FROM book_authors WHERE book_id = $1;`
				authorIDs := make([]int64, 0)
				_ = sdb.SelectContext(ctx, &authorIDs, q, testBook3.ID)
				equals(t, 0, len(authorIDs))
			})
		})
	})
}

func TestBookGetByID(t *testing.T) {
	t.Parallel()
	pgt.Runner(t, []string{"books"}, func(t *testing.T, sdb *sqlx.DB) {
		repo := &postgres.Repository{DB: sdb}
		res, err := repo.BookGetByID(context.Background(), 1)
		ok(t, err)
		equals(t, &testBook1, res)
	})
}

func TestBookList(t *testing.T) {
	t.Parallel()
	pgt.Runner(t, []string{"books"}, func(t *testing.T, sdb *sqlx.DB) {
		repo := &postgres.Repository{DB: sdb}
		res, err := repo.BookList(context.Background())
		ok(t, err)
		equals(t, []*gqlmeetup.Book{&testBook1, &testBook2, &testBook3}, res)
	})
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
				repo := &postgres.Repository{DB: sdb}
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

func TestBookUpdate(t *testing.T) {
	t.Parallel()
	pgt.Runner(t, []string{"book_authors"}, func(t *testing.T, sdb *sqlx.DB) {
		ctx := context.Background()
		repo := &postgres.Repository{DB: sdb}
		t.Run("update", func(t *testing.T) {
			res, err := repo.BookUpdate(ctx, testBook1.ID, testBookUpdate, []int64{2})
			ok(t, err)
			equals(t, &testBookUpdate, res)
			t.Run("confirm update", func(t *testing.T) {
				t.Parallel()
				res, _ := repo.BookGetByID(ctx, testBook1.ID)
				equals(t, &testBookUpdate, res)
			})
			t.Run("confirm author associations", func(t *testing.T) {
				t.Parallel()
				q := `SELECT author_id FROM book_authors WHERE book_id = $1;`
				authorIDs := make([]int64, 0)
				_ = sdb.SelectContext(ctx, &authorIDs, q, testBookUpdate.ID)
				equals(t, []int64{2}, authorIDs)
			})
		})
	})
}
