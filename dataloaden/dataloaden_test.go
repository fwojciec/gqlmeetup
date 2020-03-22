package dataloaden_test

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/fwojciec/gqlmeetup"
	"github.com/fwojciec/gqlmeetup/dataloaden"
	"github.com/fwojciec/gqlmeetup/mocks"
)

func TestAuthorListByAgentID(t *testing.T) {
	t.Parallel()
	mock := &mocks.RepositoryMock{
		AuthorListByAgentIDsFunc: func(ctx context.Context, agentIDs []int64) ([]*gqlmeetup.Author, error) {
			return []*gqlmeetup.Author{&testAuthor1, &testAuthor2, &testAuthor3}, nil
		},
	}
	dls := &dataloaden.DataLoaderService{Repository: mock}
	ctx := dls.Initialize(context.Background())
	t.Run("concurrent requests", func(t *testing.T) {
		tests := []struct {
			agentID int64
			exp     []*gqlmeetup.Author
		}{
			{2, []*gqlmeetup.Author{&testAuthor2, &testAuthor3}},
			{1, []*gqlmeetup.Author{&testAuthor1}},
		}
		for _, tc := range tests {
			tc := tc
			t.Run(fmt.Sprintf("Agent ID: %d", tc.agentID), func(t *testing.T) {
				t.Parallel()
				res, err := dls.AuthorListByAgentID(ctx, tc.agentID)
				ok(t, err)
				equals(t, tc.exp, res)
			})
		}
	})
	equals(t, 1, len(mock.AuthorListByAgentIDsCalls()))
}

func TestAgentListByIDs(t *testing.T) {
	t.Parallel()
	mock := &mocks.RepositoryMock{
		AgentListByIDsFunc: func(ctx context.Context, ids []int64) ([]*gqlmeetup.Agent, error) {
			return []*gqlmeetup.Agent{&testAgent1, &testAgent2, &testAgent3}, nil
		},
	}
	dls := dataloaden.DataLoaderService{Repository: mock}
	ctx := dls.Initialize(context.Background())
	t.Run("concurrent requests", func(t *testing.T) {
		tests := []struct {
			id  int64
			exp *gqlmeetup.Agent
		}{
			{1, &testAgent1},
			{2, &testAgent2},
			{3, &testAgent3},
		}
		for _, tc := range tests {
			tc := tc
			t.Run(fmt.Sprintf("Agent ID: %d", tc.id), func(t *testing.T) {
				t.Parallel()
				res, err := dls.AgentGetByID(ctx, tc.id)
				ok(t, err)
				equals(t, tc.exp, res)
			})
		}
	})
	equals(t, 1, len(mock.AgentListByIDsCalls()))
}

func TestBookListByAuthorID(t *testing.T) {
	t.Parallel()
	mock := &mocks.RepositoryMock{
		BookListByAuthorIDsFunc: func(ctx context.Context, authorIDs []int64) ([]*gqlmeetup.Book, error) {
			return []*gqlmeetup.Book{&testBook1, &testBook2, &testBook3}, nil
		},
	}
	dls := dataloaden.DataLoaderService{Repository: mock}
	ctx := dls.Initialize(context.Background())
	t.Run("concurrent requests", func(t *testing.T) {
		tests := []struct {
			authorID int64
			exp      []*gqlmeetup.Book
		}{
			{1, []*gqlmeetup.Book{&testBook1, &testBook3}},
			{2, []*gqlmeetup.Book{&testBook2, &testBook3}},
		}
		for _, tc := range tests {
			tc := tc
			t.Run(fmt.Sprintf("Author ID: %d", tc.authorID), func(t *testing.T) {
				t.Parallel()
				res, err := dls.BookListByAuthorID(ctx, tc.authorID)
				ok(t, err)
				equals(t, tc.exp, res)
			})
		}
	})
	equals(t, 1, len(mock.BookListByAuthorIDsCalls()))
}

func TestAuthorListByBookID(t *testing.T) {
	t.Parallel()
	mock := &mocks.RepositoryMock{
		AuthorListByBookIDsFunc: func(ctx context.Context, bookIDs []int64) ([]*gqlmeetup.Author, error) {
			return []*gqlmeetup.Author{&testAuthor1, &testAuthor2}, nil
		},
	}
	dls := dataloaden.DataLoaderService{Repository: mock}
	ctx := dls.Initialize(context.Background())
	t.Run("concurrent requests", func(t *testing.T) {
		tests := []struct {
			bookID int64
			exp    []*gqlmeetup.Author
		}{
			{1, []*gqlmeetup.Author{&testAuthor1}},
			{2, []*gqlmeetup.Author{&testAuthor2}},
			{3, []*gqlmeetup.Author{&testAuthor1, &testAuthor2}},
		}
		for _, tc := range tests {
			tc := tc
			t.Run(fmt.Sprintf("Book ID: %d", tc.bookID), func(t *testing.T) {
				t.Parallel()
				res, err := dls.AuthorListByBookID(ctx, tc.bookID)
				ok(t, err)
				equals(t, tc.exp, res)
			})
		}
	})
	equals(t, 1, len(mock.AuthorListByBookIDsCalls()))
}

func TestMiddleware(t *testing.T) {
	t.Parallel()
	repo := &mocks.RepositoryMock{
		AgentListByIDsFunc: func(ctx context.Context, ids []int64) ([]*gqlmeetup.Agent, error) { return nil, nil },
	}
	dls := &dataloaden.DataLoaderService{Repository: repo}
	req, _ := http.NewRequest("GET", "/", nil)
	handler := dls.Middleware(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		_, err := dls.AgentGetByID(r.Context(), 1)
		ok(t, err)
	}))
	handler.ServeHTTP(nil, req)
	equals(t, 1, len(repo.AgentListByIDsCalls()))
}

var (
	testAuthor1 = gqlmeetup.Author{
		ID:      1,
		Name:    "Test Author 1",
		AgentID: 1,
		BookIDs: []int64{1, 3},
	}
	testAuthor2 = gqlmeetup.Author{
		ID:      2,
		Name:    "Test Author 2",
		AgentID: 2,
		BookIDs: []int64{2, 3},
	}
	testAuthor3 = gqlmeetup.Author{
		ID:      3,
		Name:    "Test Author 3",
		AgentID: 2,
	}
	testAgent1 = gqlmeetup.Agent{
		ID:    1,
		Name:  "Test Agent 1",
		Email: "agent1@test.com",
	}
	testAgent2 = gqlmeetup.Agent{
		ID:    2,
		Name:  "Test Agent 2",
		Email: "agent2@test.com",
	}
	testAgent3 = gqlmeetup.Agent{
		ID:    3,
		Name:  "Test Agent 3",
		Email: "agent3@test.com",
	}
	testBook1 = gqlmeetup.Book{
		ID:        1,
		Title:     "Test Book 1",
		AuthorIDs: []int64{1},
	}
	testBook2 = gqlmeetup.Book{
		ID:        2,
		Title:     "Test Book 2",
		AuthorIDs: []int64{2},
	}
	testBook3 = gqlmeetup.Book{
		ID:        3,
		Title:     "Test Book 3",
		AuthorIDs: []int64{1, 2},
	}
)

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
