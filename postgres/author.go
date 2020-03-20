package postgres

import (
	"context"
	"database/sql"

	"github.com/fwojciec/gqlmeetup"
	"github.com/lib/pq"
)

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
