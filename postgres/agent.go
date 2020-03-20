package postgres

import (
	"context"
	"database/sql"

	"github.com/fwojciec/gqlmeetup"
	"github.com/lib/pq"
)

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
