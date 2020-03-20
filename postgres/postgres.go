package postgres

import (
	"bytes"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // required
)

// Open is the same as sqlx.Open, but assumes a PostgreSQL database.
func Open(dataSourceName string) (*sqlx.DB, error) {
	return sqlx.Open("postgres", dataSourceName)
}

// batchify is a helper function for batch inserts. It can be used to create
// entries is associative tables (for example book_authors, etc.). It generates
// a query prefix with placeholders and a slice of arguments that correspond to
// these placeholders.
func batchify(pid int64, cids []int64) (string, []interface{}) {
	args := make([]interface{}, len(cids)*2)
	var buf bytes.Buffer
	for i := 1; i <= len(cids); i++ {
		fmt.Fprintf(&buf, "($%d, $%d)", i*2-1, i*2)
		args[i*2-2] = pid
		args[i*2-1] = cids[i-1]
		if i < len(cids) {
			buf.WriteString(", ")
		}
	}
	return buf.String(), args
}

// withTx helps with transactions.
func withTx(ctx context.Context, db *sqlx.DB, txFn func(tx *sqlx.Tx) error) error {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	err = txFn(tx)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			err = fmt.Errorf("tx failed: %v, unable to rollback: %v", err, rbErr)
		}
	} else {
		err = tx.Commit()
	}
	return err
}
