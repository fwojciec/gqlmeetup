package dataloaden

//go:generate dataloaden AuthorSliceLoader int64 []*github.com/fwojciec/gqlmeetup.Author

import (
	"context"
	"time"

	"github.com/fwojciec/gqlmeetup"
)

type contextKey string

const key = contextKey("dataloaders")

// Loaders holds references to the individual dataloaders.
type loaders struct {
	AuthorsByAgentID *AuthorSliceLoader
}

// DataLoaderService implements gqlmeetup.DataLoaderService interface.
type DataLoaderService struct {
	Repository gqlmeetup.DataLoaderRepository
}

var _ gqlmeetup.DataLoaderService = (*DataLoaderService)(nil)

// Initialize initializes the dataloaders and adds them to the context.
func (s *DataLoaderService) Initialize(ctx context.Context) context.Context {
	return context.WithValue(ctx, key, &loaders{
		AuthorsByAgentID: newAuthorsByAgentID(ctx, s.Repository),
	})
}

// AuthorsListByAgentID lists Authors for a given agent ID.
func (s *DataLoaderService) AuthorsListByAgentID(ctx context.Context, agentID int64) ([]*gqlmeetup.Author, error) {
	l, err := s.retrieve(ctx)
	if err != nil {
		return nil, err
	}
	return l.AuthorsByAgentID.Load(agentID)
}

func (s *DataLoaderService) retrieve(ctx context.Context) (*loaders, error) {
	l, ok := ctx.Value(key).(*loaders)
	if !ok {
		return nil, gqlmeetup.ErrInvalid
	}
	return l, nil
}

func newAuthorsByAgentID(ctx context.Context, repo gqlmeetup.DataLoaderRepository) *AuthorSliceLoader {
	return NewAuthorSliceLoader(AuthorSliceLoaderConfig{
		MaxBatch: 100,
		Wait:     5 * time.Millisecond,
		Fetch: func(agentIDs []int64) ([][]*gqlmeetup.Author, []error) {
			res, err := repo.AuthorsListByAgentIDs(ctx, agentIDs)
			if err != nil {
				return nil, []error{err}
			}
			// group
			groupByAgentID := make(map[int64][]*gqlmeetup.Author, len(agentIDs))
			for _, r := range res {
				groupByAgentID[r.AgentID] = append(groupByAgentID[r.AgentID], r)
			}
			// order
			result := make([][]*gqlmeetup.Author, len(agentIDs))
			for i, agentID := range agentIDs {
				result[i] = groupByAgentID[agentID]
			}
			return result, nil
		},
	})
}
