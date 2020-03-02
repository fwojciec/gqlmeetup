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
}

func TestMutationResolver(t *testing.T) {
	t.Parallel()

	t.Run("AgentCreate", func(t *testing.T) {
		t.Parallel()
		tests := []struct {
			name string
			data gqlgen.AgentInput
			err  error
			exp  gqlmeetup.Agent
		}{
			{
				name: "successful create",
				data: gqlgen.AgentInput{
					Email: "test@email.com",
					Name:  "Test Name",
				},
				err: nil,
				exp: gqlmeetup.Agent{
					Email: "test@email.com",
					Name:  "Test Name",
				},
			},
			{
				name: "error",
				data: gqlgen.AgentInput{},
				err:  errors.New("text error"),
				exp:  gqlmeetup.Agent{},
			},
		}

		for _, tc := range tests {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				repoMock := &mocks.RepositoryMock{
					AgentCreateFunc: func(ctx context.Context, data gqlmeetup.Agent) (*gqlmeetup.Agent, error) {
						return nil, tc.err
					},
				}
				r := &gqlgen.Resolver{
					Repository: repoMock,
				}
				_, err := r.Mutation().AgentCreate(context.Background(), tc.data)
				equals(t, tc.err, err)
				equals(t, repoMock.AgentCreateCalls()[0].Data, tc.exp)
			})
		}
	})

	t.Run("AgentDelete", func(t *testing.T) {
		t.Parallel()
		repoMock := &mocks.RepositoryMock{
			AgentDeleteFunc: func(ctx context.Context, id int64) (*gqlmeetup.Agent, error) {
				return nil, nil
			},
		}
		r := &gqlgen.Resolver{
			Repository: repoMock,
		}
		_, _ = r.Mutation().AgentDelete(context.Background(), "234")
		equals(t, repoMock.AgentDeleteCalls()[0].ID, int64(234))
	})

	t.Run("AgentUpdate", func(t *testing.T) {
		t.Parallel()
		repoMock := &mocks.RepositoryMock{
			AgentUpdateFunc: func(ctx context.Context, id int64, data gqlmeetup.Agent) (*gqlmeetup.Agent, error) {
				return nil, nil
			},
		}
		r := &gqlgen.Resolver{
			Repository: repoMock,
		}
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
		tests := []struct {
			name string
			data gqlgen.AuthorInput
			err  error
			exp  gqlmeetup.Author
		}{
			{
				name: "successful create",
				data: gqlgen.AuthorInput{
					Name:    "Test Name",
					AgentID: "12",
				},
				err: nil,
				exp: gqlmeetup.Author{
					Name:    "Test Name",
					AgentID: 12,
				},
			},
			{
				name: "error",
				data: gqlgen.AuthorInput{AgentID: "1"},
				err:  errors.New("text error"),
				exp:  gqlmeetup.Author{AgentID: 1},
			},
		}
		for _, tc := range tests {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				repoMock := &mocks.RepositoryMock{
					AuthorCreateFunc: func(ctx context.Context, data gqlmeetup.Author) (*gqlmeetup.Author, error) {
						return nil, tc.err
					},
				}
				r := &gqlgen.Resolver{
					Repository: repoMock,
				}
				_, err := r.Mutation().AuthorCreate(context.Background(), tc.data)
				equals(t, tc.err, err)
				equals(t, repoMock.AuthorCreateCalls()[0].Data, tc.exp)
			})
		}
	})

	t.Run("AuthorDelete", func(t *testing.T) {
		t.Parallel()
		repoMock := &mocks.RepositoryMock{
			AuthorDeleteFunc: func(ctx context.Context, id int64) (*gqlmeetup.Author, error) {
				return nil, nil
			},
		}
		r := &gqlgen.Resolver{
			Repository: repoMock,
		}
		_, _ = r.Mutation().AuthorDelete(context.Background(), "234")
		equals(t, repoMock.AuthorDeleteCalls()[0].ID, int64(234))
	})

	t.Run("AuthorUpdate", func(t *testing.T) {
		t.Parallel()
		repoMock := &mocks.RepositoryMock{
			AuthorUpdateFunc: func(ctx context.Context, id int64, data gqlmeetup.Author) (*gqlmeetup.Author, error) {
				return nil, nil
			},
		}
		r := &gqlgen.Resolver{
			Repository: repoMock,
		}
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
			AgentsListFunc: func(ctx context.Context) ([]*gqlmeetup.Agent, error) {
				return nil, nil
			},
		}
		r := &gqlgen.Resolver{
			Repository: repoMock,
		}
		_, err := r.Query().Agents(context.Background())
		ok(t, err)
		equals(t, len(repoMock.AgentsListCalls()), 1)
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
			AuthorsListFunc: func(ctx context.Context) ([]*gqlmeetup.Author, error) {
				return nil, nil
			},
		}
		r := &gqlgen.Resolver{
			Repository: repoMock,
		}
		_, err := r.Query().Authors(context.Background())
		ok(t, err)
		equals(t, len(repoMock.AuthorsListCalls()), 1)
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
			BooksListFunc: func(ctx context.Context) ([]*gqlmeetup.Book, error) {
				return nil, nil
			},
		}
		r := &gqlgen.Resolver{
			Repository: repoMock,
		}
		_, err := r.Query().Books(context.Background())
		ok(t, err)
		equals(t, len(repoMock.BooksListCalls()), 1)
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
