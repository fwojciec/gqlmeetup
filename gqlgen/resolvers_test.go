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

// Agent Resolver -------------------------------------------------------------

func TestAgentResolverID(t *testing.T) {
	t.Parallel()
	r := &gqlgen.Resolver{}
	res, err := r.Agent().ID(context.Background(), &gqlmeetup.Agent{ID: 1337})
	ok(t, err) // should always be nil
	equals(t, "1337", res)
}

func TestAgentResolverAuthors(t *testing.T) {
	t.Parallel()
	dlMock := &mocks.DataLoaderServiceMock{
		AuthorListByAgentIDFunc: func(ctx context.Context, agentID int64) ([]*gqlmeetup.Author, error) { return nil, nil },
	}
	r := &gqlgen.Resolver{DataLoaders: dlMock}
	_, _ = r.Agent().Authors(context.Background(), &gqlmeetup.Agent{ID: 567})
	equals(t, dlMock.AuthorListByAgentIDCalls()[0].AgentID, int64(567))
}

// Author Resolver ------------------------------------------------------------

func TestAuthorResolverID(t *testing.T) {
	t.Parallel()
	r := &gqlgen.Resolver{}
	res, err := r.Author().ID(context.Background(), &gqlmeetup.Author{ID: 1337})
	ok(t, err) // should always be nil
	equals(t, "1337", res)
}

func TestAuthorResolverAgent(t *testing.T) {
	t.Parallel()
	dlMock := &mocks.DataLoaderServiceMock{
		AgentGetByIDFunc: func(ctx context.Context, id int64) (*gqlmeetup.Agent, error) { return nil, nil },
	}
	r := &gqlgen.Resolver{DataLoaders: dlMock}
	_, _ = r.Author().Agent(context.Background(), &gqlmeetup.Author{ID: 876, AgentID: 123})
	equals(t, dlMock.AgentGetByIDCalls()[0].ID, int64(123))
}

func TestAuthorResolverBooks(t *testing.T) {
	t.Parallel()
	dlMock := &mocks.DataLoaderServiceMock{
		BookListByAuthorIDFunc: func(ctx context.Context, authorID int64) ([]*gqlmeetup.Book, error) { return nil, nil },
	}
	r := &gqlgen.Resolver{DataLoaders: dlMock}
	_, _ = r.Author().Books(context.Background(), &gqlmeetup.Author{ID: 876})
	equals(t, dlMock.BookListByAuthorIDCalls()[0].AuthorID, int64(876))
}

// Book Resolver --------------------------------------------------------------

func TestBookResolverID(t *testing.T) {
	t.Parallel()
	r := &gqlgen.Resolver{}
	res, err := r.Book().ID(context.Background(), &gqlmeetup.Book{ID: 1337})
	ok(t, err) // should always be nil
	equals(t, "1337", res)
}

func TestBookResolverAuthors(t *testing.T) {
	t.Parallel()
	dlMock := &mocks.DataLoaderServiceMock{
		AuthorListByBookIDFunc: func(ctx context.Context, bookID int64) ([]*gqlmeetup.Author, error) { return nil, nil },
	}
	r := &gqlgen.Resolver{DataLoaders: dlMock}
	_, _ = r.Book().Authors(context.Background(), &gqlmeetup.Book{ID: 234})
	equals(t, dlMock.AuthorListByBookIDCalls()[0].BookID, int64(234))
}

// Mutation Resolver ----------------------------------------------------------

func TestMutationResolverLogin(t *testing.T) {
	t.Parallel()

	testUser := &gqlmeetup.User{
		Email:    "test@email.com",
		Password: "hash",
		Admin:    true,
	}

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		repoMock := &mocks.RepositoryMock{
			UserGetByEmailFunc: func(ctx context.Context, email string) (*gqlmeetup.User, error) { return testUser, nil },
		}
		pwdMock := &mocks.PasswordServiceMock{
			CheckFunc: func(pwdHash, pwd string) error { return nil },
		}
		sessionMock := &mocks.SessionServiceMock{
			LoginFunc: func(ctx context.Context, user *gqlmeetup.User) error { return nil },
		}
		r := &gqlgen.Resolver{
			Repository: repoMock,
			Password:   pwdMock,
			Session:    sessionMock,
		}
		_, _ = r.Mutation().Login(context.Background(), "test@email.com", "password")

		repoCall := repoMock.UserGetByEmailCalls()[0]
		equals(t, "test@email.com", repoCall.Email)

		pwdCall := pwdMock.CheckCalls()[0]
		equals(t, "password", pwdCall.Pwd)
		equals(t, "hash", pwdCall.PwdHash)

		sessionCall := sessionMock.LoginCalls()[0]
		equals(t, testUser, sessionCall.User)
	})

	t.Run("user not found in the db", func(t *testing.T) {
		t.Parallel()
		repoMock := &mocks.RepositoryMock{
			UserGetByEmailFunc: func(ctx context.Context, email string) (*gqlmeetup.User, error) { return nil, gqlmeetup.ErrNotFound },
		}
		r := &gqlgen.Resolver{Repository: repoMock}
		_, err := r.Mutation().Login(context.Background(), "", "")
		equals(t, gqlmeetup.ErrUnauthorized, err)
	})

	t.Run("wrong password", func(t *testing.T) {
		t.Parallel()
		repoMock := &mocks.RepositoryMock{
			UserGetByEmailFunc: func(ctx context.Context, email string) (*gqlmeetup.User, error) { return testUser, nil },
		}
		pwdMock := &mocks.PasswordServiceMock{
			CheckFunc: func(pwdHash, pwd string) error { return gqlmeetup.ErrPwdCheck },
		}
		r := &gqlgen.Resolver{
			Repository: repoMock,
			Password:   pwdMock,
		}
		_, err := r.Mutation().Login(context.Background(), "", "")
		equals(t, gqlmeetup.ErrUnauthorized, err)
	})
}

func TestMutationResolverLogout(t *testing.T) {
	t.Parallel()
	testError := errors.New("test error")
	tests := []struct {
		name        string
		logoutError error
		expError    error
		exp         bool
	}{
		{"successful", nil, nil, true},
		{"failed", testError, testError, false},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sessionMock := &mocks.SessionServiceMock{
				LogoutFunc: func(ctx context.Context) error { return tc.logoutError },
			}
			r := &gqlgen.Resolver{Session: sessionMock}
			res, err := r.Mutation().Logout(context.Background())
			equals(t, tc.exp, res)
			equals(t, tc.expError, err)
		})
	}
}

func TestMutationResolverAgentCreate(t *testing.T) {
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
}

func TestMutationResolverAgentDelete(t *testing.T) {
	t.Parallel()
	repoMock := &mocks.RepositoryMock{
		AgentDeleteFunc: func(ctx context.Context, id int64) (*gqlmeetup.Agent, error) { return nil, nil },
	}
	r := &gqlgen.Resolver{Repository: repoMock}
	_, _ = r.Mutation().AgentDelete(context.Background(), "234")
	equals(t, repoMock.AgentDeleteCalls()[0].ID, int64(234))
}

func TestMutationResolverAgentUpdate(t *testing.T) {
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
}

func TestMutationResolverAuthorCreate(t *testing.T) {
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
}

func TestMutationResolverAuthorDelete(t *testing.T) {
	t.Parallel()
	repoMock := &mocks.RepositoryMock{
		AuthorDeleteFunc: func(ctx context.Context, id int64) (*gqlmeetup.Author, error) { return nil, nil },
	}
	r := &gqlgen.Resolver{Repository: repoMock}
	_, _ = r.Mutation().AuthorDelete(context.Background(), "234")
	equals(t, repoMock.AuthorDeleteCalls()[0].ID, int64(234))
}

func TestMutationResolverAuthorUpdate(t *testing.T) {
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
}

func TestMutationResolverBookCreate(t *testing.T) {
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
}

func TestMutationResolverBookDelete(t *testing.T) {
	t.Parallel()
	repoMock := &mocks.RepositoryMock{
		BookDeleteFunc: func(ctx context.Context, id int64) (*gqlmeetup.Book, error) { return nil, nil },
	}
	r := &gqlgen.Resolver{Repository: repoMock}
	_, _ = r.Mutation().BookDelete(context.Background(), "234")
	equals(t, repoMock.BookDeleteCalls()[0].ID, int64(234))
}

func TestMutationResolverBookUpdate(t *testing.T) {
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
}

// Query Resolver -------------------------------------------------------------

func TestQueryResolverMe(t *testing.T) {
	t.Parallel()
	testUser := &gqlmeetup.User{
		Email: "test@email.com",
		Admin: true,
	}
	tests := []struct {
		name     string
		loggedIn bool
		expUser  *gqlmeetup.User
		expErr   error
	}{
		{"logged in", true, testUser, nil},
		{"logged out", false, nil, gqlmeetup.ErrNotFound},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sessionMock := &mocks.SessionServiceMock{
				GetUserFunc: func(ctx context.Context) *gqlmeetup.User {
					if tc.loggedIn {
						return testUser
					}
					return nil
				},
			}
			r := &gqlgen.Resolver{Session: sessionMock}
			res, err := r.Query().Me(context.Background())
			equals(t, tc.expErr, err)
			equals(t, tc.expUser, res)
		})
	}
}

func TestQueryResolverAgent(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		id      string
		repoErr error
		expErr  error
		exp     int64
	}{
		{"exists", "123", nil, nil, 123},
		{"not found", "22", errors.New("test error"), nil, 22},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			repoMock := &mocks.RepositoryMock{
				AgentGetByIDFunc: func(ctx context.Context, id int64) (*gqlmeetup.Agent, error) { return nil, tc.repoErr },
			}
			r := &gqlgen.Resolver{Repository: repoMock}
			_, err := r.Query().Agent(context.Background(), tc.id)
			equals(t, tc.repoErr, err)
			equals(t, repoMock.AgentGetByIDCalls()[0].ID, tc.exp)
		})
	}
}

func TestQueryResolverAgents(t *testing.T) {
	t.Parallel()
	repoMock := &mocks.RepositoryMock{
		AgentListFunc: func(ctx context.Context) ([]*gqlmeetup.Agent, error) { return nil, nil },
	}
	r := &gqlgen.Resolver{Repository: repoMock}
	_, err := r.Query().Agents(context.Background())
	ok(t, err)
	equals(t, len(repoMock.AgentListCalls()), 1)
}

func TestQueryResolverAuthor(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		id      string
		repoErr error
		expErr  error
		exp     int64
	}{
		{"exists", "123", nil, nil, 123},
		{"not found", "22", errors.New("test error"), nil, 22},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			repoMock := &mocks.RepositoryMock{
				AuthorGetByIDFunc: func(ctx context.Context, id int64) (*gqlmeetup.Author, error) { return nil, tc.repoErr },
			}
			r := &gqlgen.Resolver{Repository: repoMock}
			_, err := r.Query().Author(context.Background(), tc.id)
			equals(t, tc.repoErr, err)
			equals(t, repoMock.AuthorGetByIDCalls()[0].ID, tc.exp)
		})
	}
}

func TestQueryResolverAuthors(t *testing.T) {
	t.Parallel()
	repoMock := &mocks.RepositoryMock{
		AuthorListFunc: func(ctx context.Context) ([]*gqlmeetup.Author, error) { return nil, nil },
	}
	r := &gqlgen.Resolver{Repository: repoMock}
	_, err := r.Query().Authors(context.Background())
	ok(t, err)
	equals(t, len(repoMock.AuthorListCalls()), 1)
}

func TestQueryResolverBook(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		id      string
		repoErr error
		expErr  error
		exp     int64
	}{
		{"book exists", "123", nil, nil, 123},
		{"book not found", "22", errors.New("test error"), nil, 22},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			repoMock := &mocks.RepositoryMock{
				BookGetByIDFunc: func(ctx context.Context, id int64) (*gqlmeetup.Book, error) { return nil, tc.repoErr },
			}
			r := &gqlgen.Resolver{Repository: repoMock}
			_, err := r.Query().Book(context.Background(), tc.id)
			equals(t, tc.repoErr, err)
			equals(t, repoMock.BookGetByIDCalls()[0].ID, tc.exp)
		})
	}
}

func TestQueryResolverBooks(t *testing.T) {
	t.Parallel()
	repoMock := &mocks.RepositoryMock{
		BookListFunc: func(ctx context.Context, limit, offset *int) ([]*gqlmeetup.Book, error) { return nil, nil },
	}
	r := &gqlgen.Resolver{Repository: repoMock}
	_, err := r.Query().Books(context.Background(), intToPtr(2), nil)
	ok(t, err)
	equals(t, repoMock.BookListCalls()[0].Limit, intToPtr(2))
	equals(t, repoMock.BookListCalls()[0].Offset, (*int)(nil))
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

func intToPtr(i int) *int {
	return &i
}
