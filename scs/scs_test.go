package scs_test

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/fwojciec/gqlmeetup"
	"github.com/fwojciec/gqlmeetup/scs"
)

func TestSessionManager(t *testing.T) {
	t.Parallel()
	ss := scs.New()
	sm := ss.Middleware()
	mux := http.NewServeMux()
	mux.HandleFunc("/login", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := ss.Login(r.Context(), &gqlmeetup.User{Email: "testEmail", Admin: true, Password: "something"})
		ok(t, err)
	}))
	mux.HandleFunc("/check1", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := ss.GetUser(r.Context())
		equals(t, "testEmail", user.Email)
		equals(t, "", user.Password)
		equals(t, true, user.Admin)
	}))
	mux.HandleFunc("/logout", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := ss.Logout(r.Context())
		ok(t, err)
	}))
	mux.HandleFunc("/check2", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := ss.GetUser(r.Context())
		equals(t, (*gqlmeetup.User)(nil), user)
	}))
	ts := newTestServer(t, sm(mux))
	defer ts.Close()
	_, _ = ts.Client().Get(ts.URL + "/login")
	_, _ = ts.Client().Get(ts.URL + "/check1")
	_, _ = ts.Client().Get(ts.URL + "/logout")
	_, _ = ts.Client().Get(ts.URL + "/check2")
}

func newTestServer(t *testing.T, h http.Handler) *httptest.Server {
	ts := httptest.NewTLSServer(h)
	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Fatal(err)
	}
	ts.Client().Jar = jar
	return ts
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
