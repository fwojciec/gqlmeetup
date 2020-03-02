package gqlgen

import (
	"context"
	"strconv"

	"github.com/fwojciec/gqlmeetup"
)

// Resolver resolves a graphql query.
type Resolver struct {
	Repository gqlmeetup.Repository
}

func (r *agentResolver) ID(ctx context.Context, obj *gqlmeetup.Agent) (string, error) {
	return int64ToString(obj.ID), nil
}

func (r *agentResolver) Authors(ctx context.Context, obj *gqlmeetup.Agent) ([]*gqlmeetup.Author, error) {
	panic("not implemented")
}

func (r *authorResolver) ID(ctx context.Context, obj *gqlmeetup.Author) (string, error) {
	return int64ToString(obj.ID), nil
}

func (r *authorResolver) Agent(ctx context.Context, obj *gqlmeetup.Author) (*gqlmeetup.Agent, error) {
	panic("not implemented")
}

func (r *authorResolver) Books(ctx context.Context, obj *gqlmeetup.Author) ([]*gqlmeetup.Book, error) {
	panic("not implemented")
}

func (r *bookResolver) ID(ctx context.Context, obj *gqlmeetup.Book) (string, error) {
	return int64ToString(obj.ID), nil
}

func (r *bookResolver) Authors(ctx context.Context, obj *gqlmeetup.Book) ([]*gqlmeetup.Author, error) {
	panic("not implemented")
}

func (r *mutationResolver) AgentCreate(ctx context.Context, data AgentInput) (*gqlmeetup.Agent, error) {
	panic("not implemented")
}

func (r *mutationResolver) AgentUpdate(ctx context.Context, id string, data AgentInput) (*gqlmeetup.Agent, error) {
	panic("not implemented")
}

func (r *mutationResolver) AgentDelete(ctx context.Context, id string) (*gqlmeetup.Agent, error) {
	panic("not implemented")
}

func (r *mutationResolver) AuthorCreate(ctx context.Context, data AuthorInput) (*gqlmeetup.Author, error) {
	panic("not implemented")
}

func (r *mutationResolver) AuthorUpdate(ctx context.Context, id string, data AuthorInput) (*gqlmeetup.Author, error) {
	panic("not implemented")
}

func (r *mutationResolver) AuthorDelete(ctx context.Context, id string) (*gqlmeetup.Author, error) {
	panic("not implemented")
}

func (r *mutationResolver) BookCreate(ctx context.Context, data BookInput) (*gqlmeetup.Book, error) {
	panic("not implemented")
}

func (r *mutationResolver) BookUpdate(ctx context.Context, id string, data BookInput) (*gqlmeetup.Book, error) {
	panic("not implemented")
}

func (r *mutationResolver) BookDelete(ctx context.Context, id string) (*gqlmeetup.Book, error) {
	panic("not implemented")
}

func (r *queryResolver) Agent(ctx context.Context, id string) (*gqlmeetup.Agent, error) {
	agentID, err := stringToint64(id)
	if err != nil {
		return nil, err
	}
	agent, err := r.Repository.AgentGetByID(ctx, agentID)
	if err != nil {
		return nil, err
	}
	return agent, nil
}

func (r *queryResolver) Agents(ctx context.Context) ([]*gqlmeetup.Agent, error) {
	agents, err := r.Repository.AgentsList(ctx)
	if err != nil {
		return nil, err
	}
	return agents, nil
}

func (r *queryResolver) Author(ctx context.Context, id string) (*gqlmeetup.Author, error) {
	authorID, err := stringToint64(id)
	if err != nil {
		return nil, err
	}
	author, err := r.Repository.AuthorGetByID(ctx, authorID)
	if err != nil {
		return nil, err
	}
	return author, nil
}

func (r *queryResolver) Authors(ctx context.Context) ([]*gqlmeetup.Author, error) {
	authors, err := r.Repository.AuthorsList(ctx)
	if err != nil {
		return nil, err
	}
	return authors, nil
}

func (r *queryResolver) Book(ctx context.Context, id string) (*gqlmeetup.Book, error) {
	bookID, err := stringToint64(id)
	if err != nil {
		return nil, err
	}
	book, err := r.Repository.BookGetByID(ctx, bookID)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (r *queryResolver) Books(ctx context.Context) ([]*gqlmeetup.Book, error) {
	books, err := r.Repository.BooksList(ctx)
	if err != nil {
		return nil, err
	}
	return books, nil
}

// Agent returns an implmentation of the AgentResolver interface.
func (r *Resolver) Agent() AgentResolver { return &agentResolver{r} }

// Author returns an implmentation of the AuthorResolver interface.
func (r *Resolver) Author() AuthorResolver { return &authorResolver{r} }

// Book returns an implmentation of the BookResolver interface.
func (r *Resolver) Book() BookResolver { return &bookResolver{r} }

// Mutation returns an implmentation of the MutationResolver interface.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns an implmentation of the QueryResolver interface.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type agentResolver struct{ *Resolver }
type authorResolver struct{ *Resolver }
type bookResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

func int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func stringToint64(s string) (int64, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}
