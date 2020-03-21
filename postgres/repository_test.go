package postgres_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/fwojciec/gqlmeetup"
	"github.com/fwojciec/gqlmeetup/postgres"
	"github.com/jmoiron/sqlx"
)

// Agent ----------------------------------------------------------------------

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

// Author ----------------------------------------------------------------------

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

// Book ------------------------------------------------------------------------

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
	tests := []struct {
		name   string
		limit  *int
		offset *int
		exp    []*gqlmeetup.Book
	}{
		{"no limit or offset", nil, nil, []*gqlmeetup.Book{&testBook1, &testBook2, &testBook3}},
		{"the second result", intToPtr(1), intToPtr(1), []*gqlmeetup.Book{&testBook2}},
		{"the first result", intToPtr(1), nil, []*gqlmeetup.Book{&testBook1}},
		{"the first result explicit offset", intToPtr(1), intToPtr(0), []*gqlmeetup.Book{&testBook1}},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			pgt.Runner(t, []string{"books"}, func(t *testing.T, sdb *sqlx.DB) {
				repo := &postgres.Repository{DB: sdb}
				res, err := repo.BookList(context.Background(), tc.limit, tc.offset)
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

// User ------------------------------------------------------------------------

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

func intToPtr(i int) *int {
	return &i
}
