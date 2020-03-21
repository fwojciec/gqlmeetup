package postgres

import (
	"context"
	"database/sql"

	"github.com/fwojciec/gqlmeetup"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

// Repository implements the gqlmeetup.Repository interface.
type Repository struct {
	DB *sqlx.DB
}

var _ gqlmeetup.Repository = (*Repository)(nil)

// Agent -----------------------------------------------------------------------

const agentCreateQuery = `
INSERT INTO agents (name, email) VALUES ($1, $2) RETURNING *;`

// AgentCreate creates an agent.
func (r *Repository) AgentCreate(ctx context.Context, data gqlmeetup.Agent) (*gqlmeetup.Agent, error) {
	res := &gqlmeetup.Agent{}
	if err := r.DB.GetContext(ctx, res, agentCreateQuery, &data.Name, &data.Email); err != nil {
		return nil, err
	}
	return res, nil
}

const agentDeleteQuery = `
DELETE FROM agents WHERE id = $1 RETURNING *;`

// AgentDelete deletes an agent.
func (r *Repository) AgentDelete(ctx context.Context, id int64) (*gqlmeetup.Agent, error) {
	res := &gqlmeetup.Agent{}
	if err := r.DB.GetContext(ctx, res, agentDeleteQuery, &id); err != nil {
		return nil, err
	}
	return res, nil
}

const agentGetByIDQuery = `
SELECT * FROM agents WHERE id = $1;`

// AgentGetByID gets an agent by ID.
func (r *Repository) AgentGetByID(ctx context.Context, id int64) (*gqlmeetup.Agent, error) {
	res := &gqlmeetup.Agent{}
	if err := r.DB.GetContext(ctx, res, agentGetByIDQuery, &id); err != nil {
		if err == sql.ErrNoRows {
			return nil, gqlmeetup.ErrNotFound
		}
		return nil, err
	}
	return res, nil
}

const agentListQuery = `
SELECT * FROM agents;`

// AgentList lists all agents.
func (r *Repository) AgentList(ctx context.Context) ([]*gqlmeetup.Agent, error) {
	res := make([]*gqlmeetup.Agent, 0)
	if err := r.DB.SelectContext(ctx, &res, agentListQuery); err != nil {
		return nil, err
	}
	return res, nil
}

const agentListByIDsQuery = `
SELECT * FROM agents WHERE id = ANY($1::bigint[]);
`

// AgentListByIDs lists agents for a list of matching ids.
func (r *Repository) AgentListByIDs(ctx context.Context, ids []int64) ([]*gqlmeetup.Agent, error) {
	res := make([]*gqlmeetup.Agent, 0)
	if err := r.DB.SelectContext(ctx, &res, agentListByIDsQuery, pq.Array(ids)); err != nil {
		return nil, err
	}
	return res, nil
}

const agentUpdateQuery = `
UPDATE agents SET name=$1, email=$2 WHERE id=$3 RETURNING *;`

// AgentUpdate updates an agent.
func (r *Repository) AgentUpdate(ctx context.Context, id int64, data gqlmeetup.Agent) (*gqlmeetup.Agent, error) {
	res := &gqlmeetup.Agent{}
	if err := r.DB.GetContext(ctx, res, agentUpdateQuery, &data.Name, &data.Email, &id); err != nil {
		return nil, err
	}
	return res, nil
}

// Author ----------------------------------------------------------------------

const authorCreateQuery = `
INSERT INTO authors (name, agent_id) VALUES ($1, $2) RETURNING *;`

// AuthorCreate creates an author.
func (r *Repository) AuthorCreate(ctx context.Context, data gqlmeetup.Author) (*gqlmeetup.Author, error) {
	res := &gqlmeetup.Author{}
	if err := r.DB.GetContext(ctx, res, authorCreateQuery, &data.Name, &data.AgentID); err != nil {
		return nil, err
	}
	return res, nil
}

const authorDeleteQuery = `
DELETE FROM authors WHERE id = $1 RETURNING *;`

// AuthorDelete deletes an author.
func (r *Repository) AuthorDelete(ctx context.Context, id int64) (*gqlmeetup.Author, error) {
	res := &gqlmeetup.Author{}
	if err := r.DB.GetContext(ctx, res, authorDeleteQuery, &id); err != nil {
		return nil, err
	}
	return res, nil
}

const authorGetByIDQuery = `
SELECT * FROM authors WHERE id = $1;`

// AuthorGetByID gets an agent by ID.
func (r *Repository) AuthorGetByID(ctx context.Context, id int64) (*gqlmeetup.Author, error) {
	res := &gqlmeetup.Author{}
	if err := r.DB.GetContext(ctx, res, authorGetByIDQuery, &id); err != nil {
		if err == sql.ErrNoRows {
			return nil, gqlmeetup.ErrNotFound
		}
		return nil, err
	}
	return res, nil
}

const authorListQuery = `
SELECT * FROM authors;`

// AuthorList lists all agents.
func (r *Repository) AuthorList(ctx context.Context) ([]*gqlmeetup.Author, error) {
	res := make([]*gqlmeetup.Author, 0)
	if err := r.DB.SelectContext(ctx, &res, authorListQuery); err != nil {
		return nil, err
	}
	return res, nil
}

const authorListByAgentIDsQuery = `
SELECT authors.* FROM authors, agents
WHERE authors.agent_id = agents.id AND agents.id = ANY($1::bigint[]);
`

// AuthorListByAgentIDs lists authors for a list of matching agent ids.
func (r *Repository) AuthorListByAgentIDs(ctx context.Context, agentIDs []int64) ([]*gqlmeetup.Author, error) {
	res := make([]*gqlmeetup.Author, 0)
	if err := r.DB.SelectContext(ctx, &res, authorListByAgentIDsQuery, pq.Array(agentIDs)); err != nil {
		return nil, err
	}
	return res, nil
}

const authorListByBookIDsQuery = `
SELECT DISTINCT authors.*, ARRAY_AGG(book_authors.book_id ORDER BY book_authors.book_id) AS book_ids
FROM authors, book_authors
WHERE authors.id = book_authors.author_id AND book_authors.book_id = ANY($1::int[])
GROUP BY authors.id;`

type authorListByBookIDsRow struct {
	ID      int64
	Name    string
	AgentID int64         `json:"agent_id" db:"agent_id"`
	BookIDs pq.Int64Array `db:"book_ids"`
}

// AuthorListByBookIDs lists authors for a list of matching book ids.
func (r *Repository) AuthorListByBookIDs(ctx context.Context, bookIDs []int64) ([]*gqlmeetup.Author, error) {
	rows := make([]*authorListByBookIDsRow, 0)
	if err := r.DB.SelectContext(ctx, &rows, authorListByBookIDsQuery, pq.Array(bookIDs)); err != nil {
		return nil, err
	}
	res := make([]*gqlmeetup.Author, len(rows))
	for i, r := range rows {
		res[i] = &gqlmeetup.Author{
			ID:      r.ID,
			Name:    r.Name,
			AgentID: r.AgentID,
			BookIDs: r.BookIDs,
		}
	}
	return res, nil
}

const authorUpdateQuery = `
UPDATE authors SET name=$1, agent_id=$2 WHERE id=$3 RETURNING *;`

// AuthorUpdate updates an author.
func (r *Repository) AuthorUpdate(ctx context.Context, id int64, data gqlmeetup.Author) (*gqlmeetup.Author, error) {
	res := &gqlmeetup.Author{}
	if err := r.DB.GetContext(ctx, res, authorUpdateQuery, &data.Name, &data.AgentID, &id); err != nil {
		return nil, err
	}
	return res, nil
}

// Book ------------------------------------------------------------------------

const bookCreateQuery = `
INSERT INTO books (title) VALUES ($1) RETURNING *;`

// BookCreate creates a book
func (r *Repository) BookCreate(ctx context.Context, data gqlmeetup.Book, authorIDs []int64) (*gqlmeetup.Book, error) {
	if len(authorIDs) < 1 {
		return nil, gqlmeetup.ErrInvalid
	}
	res := &gqlmeetup.Book{}
	return res, withTx(ctx, r.DB, func(tx *sqlx.Tx) error {
		if err := tx.GetContext(ctx, res, bookCreateQuery, &data.Title); err != nil {
			return err
		}
		if err := setBookAuthors(ctx, tx, res.ID, authorIDs); err != nil {
			return err
		}
		return nil
	})
}

const bookDeleteQuery = `
DELETE FROM books WHERE id = $1 RETURNING *;`

// BookDelete deletes an book.
func (r *Repository) BookDelete(ctx context.Context, id int64) (*gqlmeetup.Book, error) {
	res := &gqlmeetup.Book{}
	if err := r.DB.GetContext(ctx, res, bookDeleteQuery, &id); err != nil {
		return nil, err
	}
	return res, nil
}

const bookGetByIDQuery = `
SELECT * FROM books WHERE id = $1;`

// BookGetByID gets a book by ID.
func (r *Repository) BookGetByID(ctx context.Context, id int64) (*gqlmeetup.Book, error) {
	res := &gqlmeetup.Book{}
	if err := r.DB.GetContext(ctx, res, bookGetByIDQuery, &id); err != nil {
		if err == sql.ErrNoRows {
			return nil, gqlmeetup.ErrNotFound
		}
		return nil, err
	}
	return res, nil
}

const bookListQuery = `
SELECT * FROM books ORDER BY title LIMIT $1 OFFSET $2;`

// BookList lists all books.
func (r *Repository) BookList(ctx context.Context, limit *int, offset *int) ([]*gqlmeetup.Book, error) {
	res := make([]*gqlmeetup.Book, 0)
	if err := r.DB.SelectContext(ctx, &res, bookListQuery, limit, offset); err != nil {
		return nil, err
	}
	return res, nil
}

const bookListByAuthorIDsQuery = `
SELECT DISTINCT books.*, ARRAY_AGG(book_authors.author_id ORDER BY book_authors.author_id) AS author_ids
FROM books, book_authors
WHERE books.id = book_authors.book_id AND book_authors.author_id = ANY($1::int[])
GROUP BY books.id;`

type bookListByAuthorIDsRow struct {
	ID        int64
	Title     string
	AuthorIDs pq.Int64Array `db:"author_ids"`
}

// BookListByAuthorIDs lists books for a list of matching author ids.
func (r *Repository) BookListByAuthorIDs(ctx context.Context, authorIDs []int64) ([]*gqlmeetup.Book, error) {
	rows := make([]*bookListByAuthorIDsRow, 0)
	if err := r.DB.SelectContext(ctx, &rows, bookListByAuthorIDsQuery, pq.Array(authorIDs)); err != nil {
		return nil, err
	}
	res := make([]*gqlmeetup.Book, len(rows))
	for i, r := range rows {
		res[i] = &gqlmeetup.Book{
			ID:        r.ID,
			Title:     r.Title,
			AuthorIDs: r.AuthorIDs,
		}
	}
	return res, nil
}

const bookUpdateQuery = `
UPDATE books SET title=$1 WHERE id=$2 RETURNING *;`

// BookUpdate updates a book.
func (r *Repository) BookUpdate(ctx context.Context, id int64, data gqlmeetup.Book, authorIDs []int64) (*gqlmeetup.Book, error) {
	if len(authorIDs) < 1 {
		return nil, gqlmeetup.ErrInvalid
	}
	res := &gqlmeetup.Book{}
	return res, withTx(ctx, r.DB, func(tx *sqlx.Tx) error {
		if err := tx.GetContext(ctx, res, bookUpdateQuery, &data.Title, &id); err != nil {
			return err
		}
		if err := unsetBookAuthors(ctx, tx, id); err != nil {
			return nil
		}
		if err := setBookAuthors(ctx, tx, id, authorIDs); err != nil {
			return err
		}
		return nil
	})
}

const bookSetAuthorsPrefix = `
INSERT INTO book_authors (book_id, author_id) VALUES `

func setBookAuthors(ctx context.Context, tx *sqlx.Tx, bookID int64, authorIDs []int64) error {
	suffix, args := batchify(bookID, authorIDs)
	_, err := tx.ExecContext(ctx, bookSetAuthorsPrefix+suffix, args...)
	return err
}

const bookUnsetAuthorsQuery = `
DELETE FROM book_authors WHERE book_id = $1;`

func unsetBookAuthors(ctx context.Context, tx *sqlx.Tx, bookID int64) error {
	_, err := tx.ExecContext(ctx, bookUnsetAuthorsQuery, bookID)
	return err
}

// User ------------------------------------------------------------------------

const userCreateQuery = `
INSERT INTO users (email, password, admin) VALUES ($1, $2, $3);`

// UserCreate creates a user.
func (r *Repository) UserCreate(data gqlmeetup.User) error {
	if _, err := r.DB.Exec(userCreateQuery, &data.Email, &data.Password, &data.Admin); err != nil {
		return err
	}
	return nil
}

const userGetByEmailQuery = `
SELECT * FROM users WHERE email = $1`

// UserGetByEmail gets a user by email address.
func (r *Repository) UserGetByEmail(ctx context.Context, email string) (*gqlmeetup.User, error) {
	res := &gqlmeetup.User{}
	if err := r.DB.GetContext(ctx, res, userGetByEmailQuery, email); err != nil {
		if err == sql.ErrNoRows {
			return nil, gqlmeetup.ErrNotFound
		}
		return nil, err
	}
	return res, nil
}
