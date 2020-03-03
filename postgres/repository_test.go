package postgres_test

import (
	"context"
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
