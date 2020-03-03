package postgres

import (
	"context"

	"github.com/fwojciec/gqlmeetup"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

// DataLoaderRepository implements the gqlmeetup.DataLoaderRepository interface.
type DataLoaderRepository struct {
	DB *sqlx.DB
}

var _ gqlmeetup.DataLoaderRepository = (*DataLoaderRepository)(nil)

const authorListByAgentIDsQuery = `
SELECT authors.* FROM authors, agents
WHERE authors.agent_id = agents.id AND agents.id = ANY($1::bigint[]);
`

// AuthorListByAgentIDs lists authors for a list of matching agent ids.
func (r *DataLoaderRepository) AuthorListByAgentIDs(ctx context.Context, agentIDs []int64) ([]*gqlmeetup.Author, error) {
	res := make([]*gqlmeetup.Author, 0)
	if err := r.DB.SelectContext(ctx, &res, authorListByAgentIDsQuery, pq.Array(agentIDs)); err != nil {
		return nil, err
	}
	return res, nil
}

const agentListByIDsQuery = `
SELECT * FROM agents WHERE id = ANY($1::bigint[]);
`

// AgentListByIDs lists agents for a list of matching ids.
func (r *DataLoaderRepository) AgentListByIDs(ctx context.Context, ids []int64) ([]*gqlmeetup.Agent, error) {
	res := make([]*gqlmeetup.Agent, 0)
	if err := r.DB.SelectContext(ctx, &res, agentListByIDsQuery, pq.Array(ids)); err != nil {
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
func (r *DataLoaderRepository) BookListByAuthorIDs(ctx context.Context, authorIDs []int64) ([]*gqlmeetup.Book, error) {
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
func (r *DataLoaderRepository) AuthorListByBookIDs(ctx context.Context, bookIDs []int64) ([]*gqlmeetup.Author, error) {
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
