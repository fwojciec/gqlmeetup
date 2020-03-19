package http_test

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
	myhttp "github.com/fwojciec/gqlmeetup/http"
	"github.com/fwojciec/gqlmeetup/mocks"
)

func TestDataLoaderMiddleware(t *testing.T) {
	t.Parallel()
	dlRepo := &mocks.DataLoaderRepositoryMock{
		AgentListByIDsFunc: func(ctx context.Context, ids []int64) ([]*gqlmeetup.Agent, error) { return nil, nil },
	}
	dls := &dataloaden.DataLoaderService{Repository: dlRepo}
	dm := myhttp.DataloaderMiddleware(dls)
	req, _ := http.NewRequest("GET", "/", nil)
	handler := dm(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		_, err := dls.AgentGetByID(r.Context(), 1)
		ok(t, err)
	}))
	handler.ServeHTTP(nil, req)
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
