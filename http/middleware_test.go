package http_test

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/fwojciec/gqlmeetup"
	"github.com/fwojciec/gqlmeetup/dataloaden"
	myhttp "github.com/fwojciec/gqlmeetup/http"
	"github.com/fwojciec/gqlmeetup/jwt"
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

func TestTokenMiddleware(t *testing.T) {
	t.Parallel()
	var (
		testTime        = time.Now()
		testNow         = func() time.Time { return testTime }
		testTokenSecret = []byte("testTokenSecret")
		testUserEmail   = "user@email.com"
		testIsAdmin     = true
		testHash        = "$2a$10$49VvJ9eRiOdJp72cjss5eeP2GdRgRCA5ojhmFEBJqq5qPdO8z27Ue"
	)
	ic := &jwt.TokenService{
		Secret:               testTokenSecret,
		AccessTokenDuration:  1 * time.Second,
		RefreshTokenDuration: 5 * time.Second,
		Now:                  testNow,
	}
	tokens, err := ic.Issue(testUserEmail, testIsAdmin, testHash)
	ok(t, err)
	tests := []struct {
		name        string
		setHeader   bool
		headerKey   string
		headerValue string
		err         error
		expAp       *gqlmeetup.AccessTokenPayload
	}{
		{
			name:        "happy path",
			setHeader:   true,
			headerKey:   "Authorization",
			headerValue: "Bearer " + tokens.Access,
			err:         nil,
		},
		{
			name:        "bad token",
			setHeader:   true,
			headerKey:   "Authorization",
			headerValue: "Bearer wrong",
			err:         gqlmeetup.ErrUnauthorized,
		},
		{
			name:        "wrong header",
			setHeader:   true,
			headerKey:   "wrong",
			headerValue: "",
			err:         gqlmeetup.ErrUnauthorized,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tm := myhttp.TokenMiddleware(ic)
			req, _ := http.NewRequest("GET", "/", nil)
			if tc.setHeader {
				req.Header.Set(tc.headerKey, tc.headerValue)
			}
			handler := tm(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
				ap, err := ic.Retrieve(r.Context())
				equals(t, tc.err, err)
				if err != nil {
					t.SkipNow()
				}
				equals(t, ap.UserEmail, testUserEmail)
				equals(t, ap.IsAdmin, testIsAdmin)
			}))
			handler.ServeHTTP(nil, req)
		})
	}
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
