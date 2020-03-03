package gqlgen_test

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/fwojciec/gqlmeetup"
	"github.com/fwojciec/gqlmeetup/gqlgen"
	"github.com/fwojciec/gqlmeetup/mocks"
)

func TestAgentResolver(t *testing.T) {
	t.Parallel()

	t.Run("ID", func(t *testing.T) {
		t.Parallel()
		r := &gqlgen.Resolver{}
		res, err := r.Agent().ID(context.Background(), &gqlmeetup.Agent{ID: 1337})
		ok(t, err) // should always be nil
		equals(t, "1337", res)
	})

	t.Run("Authors", func(t *testing.T) {
		t.Parallel()
		dlMock := &mocks.DataLoaderServiceMock{
			AuthorListByAgentIDFunc: func(ctx context.Context, agentID int64) ([]*gqlmeetup.Author, error) { return nil, nil },
		}
		r := &gqlgen.Resolver{DataLoaders: dlMock}
		_, _ = r.Agent().Authors(context.Background(), &gqlmeetup.Agent{ID: 567})
		equals(t, dlMock.AuthorListByAgentIDCalls()[0].AgentID, int64(567))
	})
}

func TestAuthorResolver(t *testing.T) {
	t.Parallel()

	t.Run("ID", func(t *testing.T) {
		t.Parallel()
		r := &gqlgen.Resolver{}
		res, err := r.Author().ID(context.Background(), &gqlmeetup.Author{ID: 1337})
		ok(t, err) // should always be nil
		equals(t, "1337", res)
	})

	t.Run("Books", func(t *testing.T) {
		t.Parallel()
		dlMock := &mocks.DataLoaderServiceMock{
			BookListByAuthorIDFunc: func(ctx context.Context, authorID int64) ([]*gqlmeetup.Book, error) { return nil, nil },
		}
		r := &gqlgen.Resolver{DataLoaders: dlMock}
		_, _ = r.Author().Books(context.Background(), &gqlmeetup.Author{ID: 876})
		equals(t, dlMock.BookListByAuthorIDCalls()[0].AuthorID, int64(876))
	})
}

func TestBookResolver(t *testing.T) {
	t.Parallel()

	t.Run("ID", func(t *testing.T) {
		t.Parallel()
		r := &gqlgen.Resolver{}
		res, err := r.Book().ID(context.Background(), &gqlmeetup.Book{ID: 1337})
		ok(t, err) // should always be nil
		equals(t, "1337", res)
	})

	t.Run("Authors", func(t *testing.T) {
		t.Parallel()
		dlMock := &mocks.DataLoaderServiceMock{
			AuthorListByBookIDFunc: func(ctx context.Context, bookID int64) ([]*gqlmeetup.Author, error) { return nil, nil },
		}
		r := &gqlgen.Resolver{DataLoaders: dlMock}
		_, _ = r.Book().Authors(context.Background(), &gqlmeetup.Book{ID: 234})
		equals(t, dlMock.AuthorListByBookIDCalls()[0].BookID, int64(234))
	})
}

func TestMutationResolver(t *testing.T) {
	t.Parallel()

	t.Run("AgentCreate", func(t *testing.T) {
		t.Parallel()
		repoMock := &mocks.RepositoryMock{
			AgentCreateFunc: func(ctx context.Context, data gqlmeetup.Agent) (*gqlmeetup.Agent, error) { return nil, nil },
		}
		r := &gqlgen.Resolver{Repository: repoMock}
		_, _ = r.Mutation().AgentCreate(context.Background(), gqlgen.AgentInput{
			Email: "test@email.com",
			Name:  "Test Name",
		})
		equals(t, repoMock.AgentCreateCalls()[0].Data, gqlmeetup.Agent{
			Email: "test@email.com",
			Name:  "Test Name",
		})
	})

	t.Run("AgentDelete", func(t *testing.T) {
		t.Parallel()
		repoMock := &mocks.RepositoryMock{
			AgentDeleteFunc: func(ctx context.Context, id int64) (*gqlmeetup.Agent, error) { return nil, nil },
		}
		r := &gqlgen.Resolver{Repository: repoMock}
		_, _ = r.Mutation().AgentDelete(context.Background(), "234")
		equals(t, repoMock.AgentDeleteCalls()[0].ID, int64(234))
	})

	t.Run("AgentUpdate", func(t *testing.T) {
		t.Parallel()
		repoMock := &mocks.RepositoryMock{
			AgentUpdateFunc: func(ctx context.Context, id int64, data gqlmeetup.Agent) (*gqlmeetup.Agent, error) { return nil, nil },
		}
		r := &gqlgen.Resolver{Repository: repoMock}
		_, _ = r.Mutation().AgentUpdate(context.Background(), "234", gqlgen.AgentInput{
			Email: "test@email.com",
			Name:  "test name",
		})
		equals(t, repoMock.AgentUpdateCalls()[0].ID, int64(234))
		equals(t, repoMock.AgentUpdateCalls()[0].Data, gqlmeetup.Agent{
			Email: "test@email.com",
			Name:  "test name",
		})
	})

	t.Run("AuthorCreate", func(t *testing.T) {
		t.Parallel()
		repoMock := &mocks.RepositoryMock{
			AuthorCreateFunc: func(ctx context.Context, data gqlmeetup.Author) (*gqlmeetup.Author, error) { return nil, nil },
		}
		r := &gqlgen.Resolver{Repository: repoMock}
		_, _ = r.Mutation().AuthorCreate(context.Background(), gqlgen.AuthorInput{
			Name:    "Test Name",
			AgentID: "12",
		})
		equals(t, repoMock.AuthorCreateCalls()[0].Data, gqlmeetup.Author{
			Name:    "Test Name",
			AgentID: 12,
		})
	})

	t.Run("AuthorDelete", func(t *testing.T) {
		t.Parallel()
		repoMock := &mocks.RepositoryMock{
			AuthorDeleteFunc: func(ctx context.Context, id int64) (*gqlmeetup.Author, error) { return nil, nil },
		}
		r := &gqlgen.Resolver{Repository: repoMock}
		_, _ = r.Mutation().AuthorDelete(context.Background(), "234")
		equals(t, repoMock.AuthorDeleteCalls()[0].ID, int64(234))
	})

	t.Run("AuthorUpdate", func(t *testing.T) {
		t.Parallel()
		repoMock := &mocks.RepositoryMock{
			AuthorUpdateFunc: func(ctx context.Context, id int64, data gqlmeetup.Author) (*gqlmeetup.Author, error) { return nil, nil },
		}
		r := &gqlgen.Resolver{Repository: repoMock}
		_, _ = r.Mutation().AuthorUpdate(context.Background(), "234", gqlgen.AuthorInput{
			Name:    "test name",
			AgentID: "567",
		})
		equals(t, repoMock.AuthorUpdateCalls()[0].ID, int64(234))
		equals(t, repoMock.AuthorUpdateCalls()[0].Data, gqlmeetup.Author{
			Name:    "test name",
			AgentID: 567,
		})
	})

	t.Run("BookCreate", func(t *testing.T) {
		t.Parallel()
		repoMock := &mocks.RepositoryMock{
			BookCreateFunc: func(ctx context.Context, data gqlmeetup.Book, authorIDs []int64) (*gqlmeetup.Book, error) {
				return nil, nil
			},
		}
		r := &gqlgen.Resolver{Repository: repoMock}
		_, _ = r.Mutation().BookCreate(context.Background(), gqlgen.BookInput{
			Title:     "Test Title",
			AuthorIDs: []string{"123", "234"},
		})
		equals(t, repoMock.BookCreateCalls()[0].Data, gqlmeetup.Book{
			Title: "Test Title",
		})
		equals(t, repoMock.BookCreateCalls()[0].AuthorIDs, []int64{123, 234})
	})

	t.Run("BookDelete", func(t *testing.T) {
		t.Parallel()
		repoMock := &mocks.RepositoryMock{
			BookDeleteFunc: func(ctx context.Context, id int64) (*gqlmeetup.Book, error) { return nil, nil },
		}
		r := &gqlgen.Resolver{Repository: repoMock}
		_, _ = r.Mutation().BookDelete(context.Background(), "234")
		equals(t, repoMock.BookDeleteCalls()[0].ID, int64(234))
	})

	t.Run("BookUpdate", func(t *testing.T) {
		t.Parallel()
		repoMock := &mocks.RepositoryMock{
			BookUpdateFunc: func(ctx context.Context, id int64, data gqlmeetup.Book, authorIDs []int64) (*gqlmeetup.Book, error) {
				return nil, nil
			},
		}
		r := &gqlgen.Resolver{Repository: repoMock}
		_, _ = r.Mutation().BookUpdate(context.Background(), "234", gqlgen.BookInput{
			Title:     "Test Title",
			AuthorIDs: []string{"123", "234"},
		})
		equals(t, repoMock.BookUpdateCalls()[0].ID, int64(234))
		equals(t, repoMock.BookUpdateCalls()[0].Data, gqlmeetup.Book{
			Title: "Test Title",
		})
	})
}

func TestQueryResolver(t *testing.T) {
	t.Parallel()

	t.Run("Agent", func(t *testing.T) {
		t.Parallel()
		tests := []struct {
			name    string
			id      string
			repoErr error
			expErr  error
			exp     int64
		}{
			{
				name:    "agent exists",
				id:      "123",
				repoErr: nil,
				expErr:  nil,
				exp:     123,
			},
			{
				name:    "agent not found",
				id:      "22",
				repoErr: errors.New("test error"),
				exp:     22,
			},
		}
		for _, tc := range tests {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				repoMock := &mocks.RepositoryMock{
					AgentGetByIDFunc: func(ctx context.Context, id int64) (*gqlmeetup.Agent, error) {
						return nil, tc.repoErr
					},
				}
				r := &gqlgen.Resolver{
					Repository: repoMock,
				}
				_, err := r.Query().Agent(context.Background(), tc.id)
				equals(t, tc.repoErr, err)
				equals(t, repoMock.AgentGetByIDCalls()[0].ID, tc.exp)
			})
		}
	})

	t.Run("Agents", func(t *testing.T) {
		t.Parallel()
		repoMock := &mocks.RepositoryMock{
			AgentListFunc: func(ctx context.Context) ([]*gqlmeetup.Agent, error) {
				return nil, nil
			},
		}
		r := &gqlgen.Resolver{
			Repository: repoMock,
		}
		_, err := r.Query().Agents(context.Background())
		ok(t, err)
		equals(t, len(repoMock.AgentListCalls()), 1)
	})

	t.Run("Author", func(t *testing.T) {
		t.Parallel()
		tests := []struct {
			name    string
			id      string
			repoErr error
			expErr  error
			exp     int64
		}{
			{
				name:    "author exists",
				id:      "123",
				repoErr: nil,
				expErr:  nil,
				exp:     123,
			},
			{
				name:    "author not found",
				id:      "22",
				repoErr: errors.New("test error"),
				exp:     22,
			},
		}
		for _, tc := range tests {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				repoMock := &mocks.RepositoryMock{
					AuthorGetByIDFunc: func(ctx context.Context, id int64) (*gqlmeetup.Author, error) {
						return nil, tc.repoErr
					},
				}
				r := &gqlgen.Resolver{
					Repository: repoMock,
				}
				_, err := r.Query().Author(context.Background(), tc.id)
				equals(t, tc.repoErr, err)
				equals(t, repoMock.AuthorGetByIDCalls()[0].ID, tc.exp)
			})
		}
	})

	t.Run("Authors", func(t *testing.T) {
		t.Parallel()
		repoMock := &mocks.RepositoryMock{
			AuthorListFunc: func(ctx context.Context) ([]*gqlmeetup.Author, error) {
				return nil, nil
			},
		}
		r := &gqlgen.Resolver{
			Repository: repoMock,
		}
		_, err := r.Query().Authors(context.Background())
		ok(t, err)
		equals(t, len(repoMock.AuthorListCalls()), 1)
	})

	t.Run("Book", func(t *testing.T) {
		t.Parallel()
		tests := []struct {
			name    string
			id      string
			repoErr error
			expErr  error
			exp     int64
		}{
			{
				name:    "book exists",
				id:      "123",
				repoErr: nil,
				expErr:  nil,
				exp:     123,
			},
			{
				name:    "book not found",
				id:      "22",
				repoErr: errors.New("test error"),
				exp:     22,
			},
		}
		for _, tc := range tests {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				repoMock := &mocks.RepositoryMock{
					BookGetByIDFunc: func(ctx context.Context, id int64) (*gqlmeetup.Book, error) {
						return nil, tc.repoErr
					},
				}
				r := &gqlgen.Resolver{
					Repository: repoMock,
				}
				_, err := r.Query().Book(context.Background(), tc.id)
				equals(t, tc.repoErr, err)
				equals(t, repoMock.BookGetByIDCalls()[0].ID, tc.exp)
			})
		}
	})

	t.Run("Books", func(t *testing.T) {
		t.Parallel()
		repoMock := &mocks.RepositoryMock{
			BookListFunc: func(ctx context.Context) ([]*gqlmeetup.Book, error) {
				return nil, nil
			},
		}
		r := &gqlgen.Resolver{
			Repository: repoMock,
		}
		_, err := r.Query().Books(context.Background())
		ok(t, err)
		equals(t, len(repoMock.BookListCalls()), 1)
	})
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
