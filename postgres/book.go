package postgres

import (
	"context"
	"database/sql"

	"github.com/fwojciec/gqlmeetup"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

const createBookQuery = `
INSERT INTO books (title) VALUES ($1) RETURNING *;`

// BookCreate creates a book
func (r *Repository) BookCreate(ctx context.Context, data gqlmeetup.Book, authorIDs []int64) (*gqlmeetup.Book, error) {
	if len(authorIDs) < 1 {
		return nil, gqlmeetup.ErrInvalid
	}
	res := &gqlmeetup.Book{}
	return res, withTx(ctx, r.DB, func(tx *sqlx.Tx) error {
		if err := tx.GetContext(ctx, res, createBookQuery, &data.Title); err != nil {
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
SELECT * FROM books;`

// BookList lists all books.
func (r *Repository) BookList(ctx context.Context) ([]*gqlmeetup.Book, error) {
	res := make([]*gqlmeetup.Book, 0)
	if err := r.DB.SelectContext(ctx, &res, bookListQuery); err != nil {
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

const updateBookQuery = `
UPDATE books SET title=$1 WHERE id=$2 RETURNING *;`

// BookUpdate updates a book.
func (r *Repository) BookUpdate(ctx context.Context, id int64, data gqlmeetup.Book, authorIDs []int64) (*gqlmeetup.Book, error) {
	if len(authorIDs) < 1 {
		return nil, gqlmeetup.ErrInvalid
	}
	res := &gqlmeetup.Book{}
	return res, withTx(ctx, r.DB, func(tx *sqlx.Tx) error {
		if err := tx.GetContext(ctx, res, updateBookQuery, &data.Title, &id); err != nil {
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

const setBookAuthorsPrefix = `
INSERT INTO book_authors (book_id, author_id) VALUES `

func setBookAuthors(ctx context.Context, tx *sqlx.Tx, bookID int64, authorIDs []int64) error {
	suffix, args := batchify(bookID, authorIDs)
	_, err := tx.ExecContext(ctx, setBookAuthorsPrefix+suffix, args...)
	return err
}

const unsetBookAuthorsQuery = `
DELETE FROM book_authors WHERE book_id = $1;`

func unsetBookAuthors(ctx context.Context, tx *sqlx.Tx, bookID int64) error {
	_, err := tx.ExecContext(ctx, unsetBookAuthorsQuery, bookID)
	return err
}
