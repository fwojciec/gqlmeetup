package dataloaden

//go:generate dataloaden AgentLoader int64 *github.com/fwojciec/gqlmeetup.Agent
//go:generate dataloaden AuthorSliceLoader int64 []*github.com/fwojciec/gqlmeetup.Author
//go:generate dataloaden BookSliceLoader int64 []*github.com/fwojciec/gqlmeetup.Book

import (
	"context"
	"net/http"
	"time"

	"github.com/fwojciec/gqlmeetup"
)

type contextKey string

const key = contextKey("dataloaders")

// Loaders holds references to the individual dataloaders.
type loaders struct {
	AgentByID        *AgentLoader
	AuthorsByAgentID *AuthorSliceLoader
	AuthorsByBookID  *AuthorSliceLoader
	BooksByAuthorID  *BookSliceLoader
}

// DataLoaderService implements gqlmeetup.DataLoaderService interface.
type DataLoaderService struct {
	Repository gqlmeetup.Repository
}

var _ gqlmeetup.DataLoaderService = (*DataLoaderService)(nil)

// Initialize initializes the dataloaders and adds them to the context.
func (s *DataLoaderService) Initialize(ctx context.Context) context.Context {
	return context.WithValue(ctx, key, &loaders{
		AgentByID:        newAgentByID(ctx, s.Repository),
		AuthorsByAgentID: newAuthorsByAgentID(ctx, s.Repository),
		AuthorsByBookID:  newAuthorsByBookID(ctx, s.Repository),
		BooksByAuthorID:  newBooksByAuthorID(ctx, s.Repository),
	})
}

// AgentGetByID gets an agent by ID.
func (s *DataLoaderService) AgentGetByID(ctx context.Context, id int64) (*gqlmeetup.Agent, error) {
	l, err := s.retrieve(ctx)
	if err != nil {
		return nil, err
	}
	return l.AgentByID.Load(id)
}

// AuthorListByAgentID lists Authors for a given agent ID.
func (s *DataLoaderService) AuthorListByAgentID(ctx context.Context, agentID int64) ([]*gqlmeetup.Author, error) {
	l, err := s.retrieve(ctx)
	if err != nil {
		return nil, err
	}
	return l.AuthorsByAgentID.Load(agentID)
}

// AuthorListByBookID lists Authors for a given agent ID.
func (s *DataLoaderService) AuthorListByBookID(ctx context.Context, bookID int64) ([]*gqlmeetup.Author, error) {
	l, err := s.retrieve(ctx)
	if err != nil {
		return nil, err
	}
	return l.AuthorsByBookID.Load(bookID)
}

// BookListByAuthorID lists Books for a given author ID.
func (s *DataLoaderService) BookListByAuthorID(ctx context.Context, authorID int64) ([]*gqlmeetup.Book, error) {
	l, err := s.retrieve(ctx)
	if err != nil {
		return nil, err
	}
	return l.BooksByAuthorID.Load(authorID)
}

// Middleware returns the HTTP middleware that enables per-request dataloader functionality.
func (s *DataLoaderService) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		nextCtx := s.Initialize(ctx)
		r = r.WithContext(nextCtx)
		next.ServeHTTP(w, r)
	})
}

func (s *DataLoaderService) retrieve(ctx context.Context) (*loaders, error) {
	l, ok := ctx.Value(key).(*loaders)
	if !ok {
		return nil, gqlmeetup.ErrInvalid
	}
	return l, nil
}

func newAgentByID(ctx context.Context, repo gqlmeetup.Repository) *AgentLoader {
	return NewAgentLoader(AgentLoaderConfig{
		MaxBatch: 100,
		Wait:     5 * time.Millisecond,
		Fetch: func(ids []int64) ([]*gqlmeetup.Agent, []error) {
			res, err := repo.AgentListByIDs(ctx, ids)
			if err != nil {
				return nil, []error{err}
			}
			groupByID := make(map[int64]*gqlmeetup.Agent, len(ids))
			for _, r := range res {
				groupByID[r.ID] = r
			}
			result := make([]*gqlmeetup.Agent, len(ids))
			for i, id := range ids {
				result[i] = groupByID[id]
			}
			return result, nil
		},
	})
}

func newAuthorsByAgentID(ctx context.Context, repo gqlmeetup.Repository) *AuthorSliceLoader {
	return NewAuthorSliceLoader(AuthorSliceLoaderConfig{
		MaxBatch: 100,
		Wait:     5 * time.Millisecond,
		Fetch: func(agentIDs []int64) ([][]*gqlmeetup.Author, []error) {
			res, err := repo.AuthorListByAgentIDs(ctx, agentIDs)
			if err != nil {
				return nil, []error{err}
			}
			groupByAgentID := make(map[int64][]*gqlmeetup.Author, len(agentIDs))
			for _, r := range res {
				groupByAgentID[r.AgentID] = append(groupByAgentID[r.AgentID], r)
			}
			result := make([][]*gqlmeetup.Author, len(agentIDs))
			for i, agentID := range agentIDs {
				result[i] = groupByAgentID[agentID]
			}
			return result, nil
		},
	})
}

func newAuthorsByBookID(ctx context.Context, repo gqlmeetup.Repository) *AuthorSliceLoader {
	return NewAuthorSliceLoader(AuthorSliceLoaderConfig{
		MaxBatch: 100,
		Wait:     5 * time.Millisecond,
		Fetch: func(bookIDs []int64) ([][]*gqlmeetup.Author, []error) {
			res, err := repo.AuthorListByBookIDs(ctx, bookIDs)
			if err != nil {
				return nil, []error{err}
			}
			groupByBookID := make(map[int64][]*gqlmeetup.Author, len(bookIDs))
			for _, r := range res {
				for _, bookID := range r.BookIDs {
					groupByBookID[bookID] = append(groupByBookID[bookID], r)
				}
			}
			result := make([][]*gqlmeetup.Author, len(bookIDs))
			for i, bookID := range bookIDs {
				result[i] = groupByBookID[bookID]
			}
			return result, nil
		},
	})
}

func newBooksByAuthorID(ctx context.Context, repo gqlmeetup.Repository) *BookSliceLoader {
	return NewBookSliceLoader(BookSliceLoaderConfig{
		MaxBatch: 100,
		Wait:     5 * time.Millisecond,
		Fetch: func(authorIDs []int64) ([][]*gqlmeetup.Book, []error) {
			res, err := repo.BookListByAuthorIDs(ctx, authorIDs)
			if err != nil {
				return nil, []error{err}
			}
			groupByAuthorID := make(map[int64][]*gqlmeetup.Book, len(authorIDs))
			for _, r := range res {
				for _, authorID := range r.AuthorIDs {
					groupByAuthorID[authorID] = append(groupByAuthorID[authorID], r)
				}
			}
			result := make([][]*gqlmeetup.Book, len(authorIDs))
			for i, authorID := range authorIDs {
				result[i] = groupByAuthorID[authorID]
			}
			return result, nil
		},
	})
}
