package gqlmeetup

import (
	"context"
	"net/http"
)

// User is a user of the website.
type User struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Admin    bool   `json:"admin,omitempty"`
}

// Agent is an agent working in the agency.
type Agent struct {
	ID    int64
	Name  string
	Email string
}

// Author is a writer represented by the agency.
type Author struct {
	ID      int64
	Name    string
	AgentID int64 `json:"agent_id" db:"agent_id"`
	BookIDs []int64
}

// Book is a book written by an author.
type Book struct {
	ID        int64
	Title     string
	AuthorIDs []int64
}

// BookAuthor is an associative table between books and authors.
type BookAuthor struct {
	ID       int64
	BookID   int64 `json:"book_id" db:"book_id"`
	AuthorID int64 `json:"author_id" db:"author_id"`
}

// Repository represents database functionality.
type Repository interface {
	// Agent
	AgentCreate(ctx context.Context, data Agent) (*Agent, error)
	AgentDelete(ctx context.Context, id int64) (*Agent, error)
	AgentGetByID(ctx context.Context, id int64) (*Agent, error)
	AgentList(ctx context.Context) ([]*Agent, error)
	AgentListByIDs(ctx context.Context, ids []int64) ([]*Agent, error)
	AgentUpdate(ctx context.Context, id int64, data Agent) (*Agent, error)

	// Author
	AuthorCreate(ctx context.Context, data Author) (*Author, error)
	AuthorDelete(ctx context.Context, id int64) (*Author, error)
	AuthorGetByID(ctx context.Context, id int64) (*Author, error)
	AuthorList(ctx context.Context) ([]*Author, error)
	AuthorListByAgentIDs(ctx context.Context, agentIDs []int64) ([]*Author, error)
	AuthorListByBookIDs(ctx context.Context, bookIDs []int64) ([]*Author, error)
	AuthorUpdate(ctx context.Context, id int64, data Author) (*Author, error)

	// Book
	BookCreate(ctx context.Context, data Book, authorIDs []int64) (*Book, error)
	BookDelete(ctx context.Context, id int64) (*Book, error)
	BookGetByID(ctx context.Context, id int64) (*Book, error)
	BookList(ctx context.Context) ([]*Book, error)
	BookListByAuthorIDs(ctx context.Context, authorIDs []int64) ([]*Book, error)
	BookUpdate(ctx context.Context, id int64, data Book, authorIDs []int64) (*Book, error)

	// User
	UserCreate(data User) error
	UserGetByEmail(ctx context.Context, email string) (*User, error)
}

// DataLoaderService provides dataloader functionality for the resolvers.
type DataLoaderService interface {
	AgentGetByID(ctx context.Context, id int64) (*Agent, error)
	AuthorListByAgentID(ctx context.Context, agentID int64) ([]*Author, error)
	AuthorListByBookID(ctx context.Context, bookID int64) ([]*Author, error)
	BookListByAuthorID(ctx context.Context, authorID int64) ([]*Book, error)
	Initialize(ctx context.Context) context.Context
	Middleware(http.Handler) http.Handler
}

// PasswordService performs password checking and hashing.
type PasswordService interface {
	Check(pwdHash, pwd string) error
	Hash(pwd string) (string, error)
}

// SessionService manages the session.
type SessionService interface {
	Login(ctx context.Context, user *User) error
	Logout(ctx context.Context) error
	GetUser(ctx context.Context) *User
	Middleware(http.Handler) http.Handler
}
