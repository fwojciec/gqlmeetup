package dataloaden_test

import (
	"context"
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/fwojciec/gqlmeetup"
	"github.com/fwojciec/gqlmeetup/dataloaden"
	"github.com/fwojciec/gqlmeetup/mocks"
)

func TestAuthorsListByAgentID(t *testing.T) {
	t.Parallel()
	mock := &mocks.DataLoaderRepositoryMock{
		AuthorsListByAgentIDsFunc: func(ctx context.Context, agentIDs []int64) ([]*gqlmeetup.Author, error) {
			return []*gqlmeetup.Author{&testAuthor1, &testAuthor2, &testAuthor3}, nil
		},
	}
	dls := dataloaden.DataLoaderService{Repository: mock}
	ctx := dls.Initialize(context.Background())
	t.Run("concurrent requests", func(t *testing.T) {
		t.Parallel()
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
				res, err := dls.AuthorsListByAgentID(ctx, tc.agentID)
				ok(t, err)
				equals(t, tc.exp, res)
			})
		}
	})
}

var (
	testAuthor1 = gqlmeetup.Author{
		ID:      1,
		Name:    "Test Author 1",
		AgentID: 1,
	}
	testAuthor2 = gqlmeetup.Author{
		ID:      2,
		Name:    "Test Author 2",
		AgentID: 2,
	}
	testAuthor3 = gqlmeetup.Author{
		ID:      3,
		Name:    "Test Author 3",
		AgentID: 2,
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
