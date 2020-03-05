package postgres_test

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/fwojciec/gqlmeetup"
	"github.com/fwojciec/pgtester"
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
	testAgentCreate = gqlmeetup.Agent{
		ID:    3,
		Name:  "Test Agent 3",
		Email: "agent3@test.com",
	}
	testAgentUpdate = gqlmeetup.Agent{
		ID:    1,
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
	testAuthorCreate = gqlmeetup.Author{
		ID:      3,
		Name:    "Test Author 3",
		AgentID: 1,
	}
	testAuthorUpdate = gqlmeetup.Author{
		ID:      1,
		Name:    "Test Author 3",
		AgentID: 2,
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
	testBookCreate = gqlmeetup.Book{
		ID:    4,
		Title: "Test Book 4",
	}
	testBookUpdate = gqlmeetup.Book{
		ID:    1,
		Title: "Test Book 5",
	}
	testUser1 = gqlmeetup.User{
		Email:    "user1@email.com",
		Password: "password1",
		Admin:    true,
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
	bookAuthorsSetup = `
CREATE TABLE IF NOT EXISTS book_authors (
	id bigserial PRIMARY KEY,
	book_id bigint NOT NULL,
	author_id bigint NOT NULL,
	FOREIGN KEY (book_id) REFERENCES books (id) ON DELETE CASCADE,
	FOREIGN KEY (author_id) REFERENCES authors (id) ON DELETE CASCADE,
	UNIQUE (book_id, author_id)
);
INSERT INTO book_authors (book_id, author_id) VALUES
(1, 1),
(2, 2),
(3, 1),
(3, 2);
`
	usersSetup = `
CREATE TABLE IF NOT EXISTS users (
	email varchar(254) PRIMARY KEY,
	password varchar(60) NOT NULL,
	admin boolean
);
INSERT INTO users (email, password, admin) VALUES
('user1@email.com', 'password1', true),
('user2@email.com', 'password2', false);
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
	"book_authors": pgtester.TableSchema{
		SetupSQL: bookAuthorsSetup,
		Deps:     []string{"books", "authors"},
	},
	"users": pgtester.TableSchema{
		SetupSQL: usersSetup,
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
