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
	AgentDelete(ctx context.Context, id int64) (*Agent, error)
	AgentGetByID(ctx context.Context, id int64) (*Agent, error)
	AgentList(ctx context.Context) ([]*Agent, error)
	AgentUpdate(ctx context.Context, id int64, data Agent) (*Agent, error)
	AuthorCreate(ctx context.Context, data Author) (*Author, error)
	AuthorDelete(ctx context.Context, id int64) (*Author, error)
	AuthorGetByID(ctx context.Context, id int64) (*Author, error)
	AuthorList(ctx context.Context) ([]*Author, error)
	AuthorUpdate(ctx context.Context, id int64, data Author) (*Author, error)
	BookCreate(ctx context.Context, data Book, authorIDs []int64) (*Book, error)
	BookDelete(ctx context.Context, id int64) (*Book, error)
	BookGetByID(ctx context.Context, id int64) (*Book, error)
	BookList(ctx context.Context) ([]*Book, error)
	BookUpdate(ctx context.Context, id int64, data Book, authorIDs []int64) (*Book, error)
}

// DataLoaderRepository represents database functionality used by the dataloader
// service.
type DataLoaderRepository interface {
	AuthorListByAgentIDs(ctx context.Context, agentIDs []int64) ([]*Author, error)
	AgentListByIDs(ctx context.Context, ids []int64) ([]*Agent, error)
}

// DataLoaderService provides dataloader functionality for the resolvers.
type DataLoaderService interface {
	Initialize(ctx context.Context) context.Context
	AuthorListByAgentID(ctx context.Context, agentID int64) ([]*Author, error)
	AgentGetByID(ctx context.Context, id int64) (*Agent, error)
}
