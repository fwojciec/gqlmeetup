package postgres_test

import (
	"context"
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/fwojciec/gqlmeetup"
	"github.com/fwojciec/gqlmeetup/postgres"
	"github.com/fwojciec/pgtester"
	"github.com/jmoiron/sqlx"
)

var (
	pgt          *pgtester.PGT
	dbConnString = "dbname=test_db sslmode=disable"
)

func init() {
	var err error
	if pgt, err = pgtester.New(dbConnString, schema); err != nil {
		panic(err)
	}
}

func TestAgentCreate(t *testing.T) {
	t.Parallel()
	pgt.Runner(t, []string{"agents"}, func(t *testing.T, sdb *sqlx.DB) {
		ctx := context.Background()
		repo := postgres.Repository{DB: sdb}
		t.Run("create", func(t *testing.T) {
			res, err := repo.AgentCreate(ctx, testAgent3)
			ok(t, err)
			equals(t, &testAgent3, res)
			t.Run("assert agent was created", func(t *testing.T) {
				t.Parallel()
				res, _ := repo.AgentGetByID(ctx, testAgent3.ID)
				equals(t, &testAgent3, res)
			})
		})
	})
}

func TestAgentGetByID(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		id   int64
		exp  *gqlmeetup.Agent
		err  error
	}{
		{
			name: "agent exists",
			id:   1,
			exp:  &testAgent1,
			err:  nil,
		},
		{
			name: "agent doesn't exist",
			id:   11,
			exp:  nil,
			err:  gqlmeetup.ErrNotFound,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			pgt.Runner(t, []string{"agents"}, func(t *testing.T, sdb *sqlx.DB) {
				repo := postgres.Repository{DB: sdb}
				res, err := repo.AgentGetByID(context.Background(), tc.id)
				equals(t, tc.err, err)
				equals(t, tc.exp, res)
			})
		})
	}
}

func TestAgentsList(t *testing.T) {
	t.Parallel()
	pgt.Runner(t, []string{"agents"}, func(t *testing.T, sdb *sqlx.DB) {
		repo := postgres.Repository{DB: sdb}
		res, err := repo.AgentsList(context.Background())
		ok(t, err)
		equals(t, []*gqlmeetup.Agent{&testAgent1, &testAgent2}, res)
	})
}

func TestAuthorCreate(t *testing.T) {
	t.Parallel()
	pgt.Runner(t, []string{"authors"}, func(t *testing.T, sdb *sqlx.DB) {
		ctx := context.Background()
		repo := postgres.Repository{DB: sdb}
		t.Run("create", func(t *testing.T) {
			res, err := repo.AuthorCreate(ctx, testAuthor3)
			ok(t, err)
			equals(t, &testAuthor3, res)
			t.Run("assert agent was created", func(t *testing.T) {
				t.Parallel()
				res, _ := repo.AuthorGetByID(ctx, testAuthor3.ID)
				equals(t, &testAuthor3, res)
			})
		})
	})
}

func TestAuthorGetByID(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		id   int64
		exp  *gqlmeetup.Author
		err  error
	}{
		{
			name: "author exists",
			id:   1,
			exp:  &testAuthor1,
			err:  nil,
		},
		{
			name: "author doesn't exist",
			id:   11,
			exp:  nil,
			err:  gqlmeetup.ErrNotFound,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			pgt.Runner(t, []string{"authors"}, func(t *testing.T, sdb *sqlx.DB) {
				repo := postgres.Repository{DB: sdb}
				res, err := repo.AuthorGetByID(context.Background(), tc.id)
				equals(t, tc.err, err)
				equals(t, tc.exp, res)
			})
		})
	}
}

func TestAuthorsList(t *testing.T) {
	t.Parallel()
	pgt.Runner(t, []string{"authors"}, func(t *testing.T, sdb *sqlx.DB) {
		repo := postgres.Repository{DB: sdb}
		res, err := repo.AuthorsList(context.Background())
		ok(t, err)
		equals(t, []*gqlmeetup.Author{&testAuthor1, &testAuthor2}, res)
	})
}

func TestBookGetByID(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		id   int64
		exp  *gqlmeetup.Book
		err  error
	}{
		{
			name: "book exists",
			id:   1,
			exp:  &testBook1,
			err:  nil,
		},
		{
			name: "book doesn't exist",
			id:   11,
			exp:  nil,
			err:  gqlmeetup.ErrNotFound,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			pgt.Runner(t, []string{"books"}, func(t *testing.T, sdb *sqlx.DB) {
				repo := postgres.Repository{DB: sdb}
				res, err := repo.BookGetByID(context.Background(), tc.id)
				equals(t, tc.err, err)
				equals(t, tc.exp, res)
			})
		})
	}
}

func TestBooksList(t *testing.T) {
	t.Parallel()
	pgt.Runner(t, []string{"books"}, func(t *testing.T, sdb *sqlx.DB) {
		repo := postgres.Repository{DB: sdb}
		res, err := repo.BooksList(context.Background())
		ok(t, err)
		equals(t, []*gqlmeetup.Book{&testBook1, &testBook2, &testBook3}, res)
	})
}

var (
	testAgent1 = gqlmeetup.Agent{
		ID:    1,
		Name:  "Test Agent 1",
		Email: "agent1@test.com",
	}
	testAgent2 = gqlmeetup.Agent{
		ID:    2,
		Name:  "Test Agent 2",
		Email: "agent2@test.com",
	}
	testAgent3 = gqlmeetup.Agent{
		ID:    3,
		Name:  "Test Agent 3",
		Email: "agent3@test.com",
	}
	testAuthor1 = gqlmeetup.Author{
		ID:      1,
		Name:    "Test Author 1",
		AgentID: 1,
	}
	testAuthor2 = gqlmeetup.Author{
		ID:      2,
		Name:    "Test Author 2",
		AgentID: 2,
	}
	testAuthor3 = gqlmeetup.Author{
		ID:      3,
		Name:    "Test Author 3",
		AgentID: 1,
	}
	testBook1 = gqlmeetup.Book{
		ID:    1,
		Title: "Test Book 1",
	}
	testBook2 = gqlmeetup.Book{
		ID:    2,
		Title: "Test Book 2",
	}
	testBook3 = gqlmeetup.Book{
		ID:    3,
		Title: "Test Book 3",
	}
)

const (
	agentsSetup = `
CREATE TABLE IF NOT EXISTS agents (
	id bigserial PRIMARY KEY,
	name text NOT NULL,
	email text NOT NULL UNIQUE
);
INSERT INTO agents (name, email) VALUES
('Test Agent 1', 'agent1@test.com'),
('Test Agent 2', 'agent2@test.com');
`
	authorsSetup = `
CREATE TABLE IF NOT EXISTS authors (
	id bigserial PRIMARY KEY,
	name text NOT NULL,
	agent_id bigint NOT NULL,
	FOREIGN KEY (agent_id) REFERENCES agents (id) ON DELETE RESTRICT
);
INSERT INTO authors (name, agent_id) VALUES
('Test Author 1', 1),
('Test Author 2', 2);
`
	booksSetup = `
CREATE TABLE IF NOT EXISTS books (
	id bigserial PRIMARY KEY,
	title text NOT NULL
);
INSERT INTO books (title) VALUES
('Test Book 1'),
('Test Book 2'),
('Test Book 3');
`
)

var schema pgtester.Schema = pgtester.Schema{
	"agents": pgtester.TableSchema{
		SetupSQL: agentsSetup,
	},
	"authors": pgtester.TableSchema{
		SetupSQL: authorsSetup,
		Deps:     []string{"agents"},
	},
	"books": pgtester.TableSchema{
		SetupSQL: booksSetup,
	},
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
