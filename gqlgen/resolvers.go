package gqlgen

import (
	"context"
	"strconv"

	"github.com/fwojciec/gqlmeetup"
)

// Resolver resolves a graphql query.
type Resolver struct {
	Repository  gqlmeetup.Repository
	DataLoaders gqlmeetup.DataLoaderService
	Password    gqlmeetup.PasswordService
	Tokens      gqlmeetup.TokenService
}

func (r *agentResolver) ID(ctx context.Context, obj *gqlmeetup.Agent) (string, error) {
	return int64ToString(obj.ID), nil
}

func (r *agentResolver) Authors(ctx context.Context, obj *gqlmeetup.Agent) ([]*gqlmeetup.Author, error) {
	return r.DataLoaders.AuthorListByAgentID(ctx, obj.ID)
}

func (r *authorResolver) ID(ctx context.Context, obj *gqlmeetup.Author) (string, error) {
	return int64ToString(obj.ID), nil
}

func (r *authorResolver) Agent(ctx context.Context, obj *gqlmeetup.Author) (*gqlmeetup.Agent, error) {
	return r.DataLoaders.AgentGetByID(ctx, obj.AgentID)
}

func (r *authorResolver) Books(ctx context.Context, obj *gqlmeetup.Author) ([]*gqlmeetup.Book, error) {
	return r.DataLoaders.BookListByAuthorID(ctx, obj.ID)
}

func (r *bookResolver) ID(ctx context.Context, obj *gqlmeetup.Book) (string, error) {
	return int64ToString(obj.ID), nil
}

func (r *bookResolver) Authors(ctx context.Context, obj *gqlmeetup.Book) ([]*gqlmeetup.Author, error) {
	return r.DataLoaders.AuthorListByBookID(ctx, obj.ID)
}

func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*gqlmeetup.Tokens, error) {
	user, err := r.Repository.UserGetByEmail(ctx, email)
	if err != nil {
		if err == gqlmeetup.ErrNotFound {
			return nil, gqlmeetup.ErrUnauthorized
		}
		return nil, err
	}
	if err := r.Password.Check(user.Password, password); err != nil {
		return nil, gqlmeetup.ErrUnauthorized
	}
	tokens, err := r.Tokens.Issue(user.Email, user.Admin, user.Password)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

func (r *mutationResolver) Refresh(ctx context.Context, token string) (*gqlmeetup.Tokens, error) {
	email, err := r.Tokens.DecodeRefreshToken(token)
	if err != nil {
		return nil, gqlmeetup.ErrUnauthorized
	}
	user, err := r.Repository.UserGetByEmail(ctx, email)
	if err != nil {
		if err == gqlmeetup.ErrNotFound {
			return nil, gqlmeetup.ErrUnauthorized
		}
		return nil, err
	}
	_, err = r.Tokens.CheckRefreshToken(token, user.Password)
	if err != nil {
		return nil, gqlmeetup.ErrUnauthorized
	}
	tokens, err := r.Tokens.Issue(user.Email, user.Admin, user.Password)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

func (r *mutationResolver) AgentCreate(ctx context.Context, data AgentInput) (*gqlmeetup.Agent, error) {
	res, err := r.Repository.AgentCreate(ctx, gqlmeetup.Agent{
		Email: data.Email,
		Name:  data.Name,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) AgentDelete(ctx context.Context, id string) (*gqlmeetup.Agent, error) {
	agentID, err := stringToInt64(id)
	if err != nil {
		return nil, err
	}
	res, err := r.Repository.AgentDelete(ctx, agentID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) AgentUpdate(ctx context.Context, id string, data AgentInput) (*gqlmeetup.Agent, error) {
	agentID, err := stringToInt64(id)
	if err != nil {
		return nil, err
	}
	res, err := r.Repository.AgentUpdate(ctx, agentID, gqlmeetup.Agent{
		Email: data.Email,
		Name:  data.Name,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) AuthorCreate(ctx context.Context, data AuthorInput) (*gqlmeetup.Author, error) {
	agentID, err := stringToInt64(data.AgentID)
	if err != nil {
		return nil, err
	}
	res, err := r.Repository.AuthorCreate(ctx, gqlmeetup.Author{
		Name:    data.Name,
		AgentID: agentID,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) AuthorDelete(ctx context.Context, id string) (*gqlmeetup.Author, error) {
	authorID, err := stringToInt64(id)
	if err != nil {
		return nil, err
	}
	res, err := r.Repository.AuthorDelete(ctx, authorID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) AuthorUpdate(ctx context.Context, id string, data AuthorInput) (*gqlmeetup.Author, error) {
	authorID, err := stringToInt64(id)
	if err != nil {
		return nil, err
	}
	agentID, err := stringToInt64(data.AgentID)
	if err != nil {
		return nil, err
	}
	res, err := r.Repository.AuthorUpdate(ctx, authorID, gqlmeetup.Author{
		Name:    data.Name,
		AgentID: agentID,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) BookCreate(ctx context.Context, data BookInput) (*gqlmeetup.Book, error) {
	authorIDs, err := stringSliceToInt64Slice(data.AuthorIDs)
	if err != nil {
		return nil, err
	}
	res, err := r.Repository.BookCreate(ctx, gqlmeetup.Book{Title: data.Title}, authorIDs)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) BookUpdate(ctx context.Context, id string, data BookInput) (*gqlmeetup.Book, error) {
	bookID, err := stringToInt64(id)
	if err != nil {
		return nil, err
	}
	authorIDs, err := stringSliceToInt64Slice(data.AuthorIDs)
	if err != nil {
		return nil, err
	}
	res, err := r.Repository.BookUpdate(ctx, bookID, gqlmeetup.Book{Title: data.Title}, authorIDs)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) BookDelete(ctx context.Context, id string) (*gqlmeetup.Book, error) {
	bookID, err := stringToInt64(id)
	if err != nil {
		return nil, err
	}
	res, err := r.Repository.BookDelete(ctx, bookID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *queryResolver) Agent(ctx context.Context, id string) (*gqlmeetup.Agent, error) {
	agentID, err := stringToInt64(id)
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
	agents, err := r.Repository.AgentList(ctx)
	if err != nil {
		return nil, err
	}
	return agents, nil
}

func (r *queryResolver) Author(ctx context.Context, id string) (*gqlmeetup.Author, error) {
	authorID, err := stringToInt64(id)
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
	authors, err := r.Repository.AuthorList(ctx)
	if err != nil {
		return nil, err
	}
	return authors, nil
}

func (r *queryResolver) Book(ctx context.Context, id string) (*gqlmeetup.Book, error) {
	bookID, err := stringToInt64(id)
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
	books, err := r.Repository.BookList(ctx)
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
type userResolver struct{ *Resolver }

func int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func stringToInt64(s string) (int64, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func stringSliceToInt64Slice(ss []string) ([]int64, error) {
	res := make([]int64, len(ss))
	for i, s := range ss {
		id, err := stringToInt64(s)
		if err != nil {
			return nil, err
		}
		res[i] = id
	}
	return res, nil
}
