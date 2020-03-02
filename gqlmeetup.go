package gqlmeetup

import "context"

// Agent is an employee of the agency.
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
}

// Book is a book written by an author.
type Book struct {
	ID    int64
	Title string
}

// BookAuthor is an associative table between books and authors.
type BookAuthor struct {
	ID       int64
	BookID   int64 `json:"book_id" db:"book_id"`
	AuthorID int64 `json:"author_id" db:"author_id"`
}

// Repository represents database functionality.
type Repository interface {
	AgentCreate(ctx context.Context, data Agent) (*Agent, error)
	AgentGetByID(ctx context.Context, id int64) (*Agent, error)
	AgentsList(ctx context.Context) ([]*Agent, error)
	AuthorCreate(ctx context.Context, data Author) (*Author, error)
	AuthorGetByID(ctx context.Context, id int64) (*Author, error)
	AuthorsList(ctx context.Context) ([]*Author, error)
	BookGetByID(ctx context.Context, id int64) (*Book, error)
	BooksList(ctx context.Context) ([]*Book, error)
}
