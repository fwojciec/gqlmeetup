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

const authorsListByAgentIDsQuery = `
SELECT authors.* FROM authors, agents
WHERE authors.agent_id = agents.id AND agents.id = ANY($1::bigint[]);
`

// AuthorsListByAgentIDs lists authors for a list of matching agent ids.
func (r *DataLoaderRepository) AuthorsListByAgentIDs(ctx context.Context, agentIDs []int64) ([]*gqlmeetup.Author, error) {
	res := make([]*gqlmeetup.Author, 0)
	if err := r.DB.SelectContext(ctx, &res, authorsListByAgentIDsQuery, pq.Array(agentIDs)); err != nil {
		return nil, err
	}
	return res, nil
}
