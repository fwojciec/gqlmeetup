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
	BookListByAuthorIDs(ctx context.Context, authorIDs []int64) ([]*Book, error)
	AuthorListByBookIDs(ctx context.Context, bookIDs []int64) ([]*Author, error)
}

// DataLoaderService provides dataloader functionality for the resolvers.
type DataLoaderService interface {
	AgentGetByID(ctx context.Context, id int64) (*Agent, error)
	AuthorListByAgentID(ctx context.Context, agentID int64) ([]*Author, error)
	AuthorListByBookID(ctx context.Context, bookID int64) ([]*Author, error)
	BookListByAuthorID(ctx context.Context, authorID int64) ([]*Book, error)
	Initialize(ctx context.Context) context.Context
}

// PasswordService performs password checking and hashing.
type PasswordService interface {
	Check(pwdHash, pwd string) error
	Hash(pwd string) (string, error)
}

// Tokens contains a pair of generated tokens and the Unix timestamp of access
// token expiration.
type Tokens struct {
	Access    string
	Refresh   string
	ExpiresAt int
}

// AccessTokenPayload is the result of a successful access token validation.
type AccessTokenPayload struct {
	UserID  int64
	IsAdmin bool
}

// RefreshTokenPayload is the result of a successful refresh token validation.
type RefreshTokenPayload struct {
	UserID int64
}

// TokenService performs token-related tasks.
type TokenService interface {
	IssueTokens(userID int64, isAdmin bool, pwdHash string) (*Tokens, error)
	DecodeRefreshToken(token string) (int64, error)
	CheckRefreshToken(token string, pwdHash string) (*RefreshTokenPayload, error)
	CheckAccessToken(token string) (*AccessTokenPayload, error)
	Retrieve(ctx context.Context) (*AccessTokenPayload, error)
	Store(context.Context, *AccessTokenPayload) context.Context
}
