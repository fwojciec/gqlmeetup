package jwt_test

import (
	"context"
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/fwojciec/gqlmeetup"
	"github.com/fwojciec/gqlmeetup/jwt"
)

var (
	testTime        = time.Now()
	testNow         = func() time.Time { return testTime }
	testTokenSecret = []byte("testTokenSecret")
	testATDuration  = 1 * time.Second
	testRTDuration  = 5 * time.Second
	testUserID      = int64(99)
	testIsAdmin     = true
	testHash        = "$2a$10$49VvJ9eRiOdJp72cjss5eeP2GdRgRCA5ojhmFEBJqq5qPdO8z27Ue"
	testTokenWrong1 = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.POstGetfAytaZS82wHcjoTyoqhMyxXiWdR7Nn7A29DNSl0EiXLdwJ6xC6AfgZWF1bOsS_TuYI3OG85AmiExREkrS6tDfTQ2B3WXlrr-wp5AokiRbz3_oB4OxG-W9KcEEbDRcZc0nH3L7LzYptiy1PtAylQGxHTWZXtGz4ht0bAecBgmpdgXMguEIcoqPJ1n3pIWk_dUZegpqx0Lka21H6XxUTxiy8OcaarA8zdnPUnV6AmNP3ecFawIFYdvJB_cm-GvpCSbr8G8y_Mllj8f4x9nBH8pQux89_6gUY618iYv7tuPWBFfEbLxtF2pZS6YC1aSfLQxeNe8djT9YjpvRZA"
	testATPayload   = &gqlmeetup.AccessTokenPayload{UserID: testUserID, IsAdmin: testIsAdmin}
)

func TestIssueCheckAndDecodeTokens(t *testing.T) {
	t.Parallel()

	var (
		testAccessToken  string
		testRefreshToken string
	)

	ic := &jwt.TokenService{
		Secret:               testTokenSecret,
		AccessTokenDuration:  testATDuration,
		RefreshTokenDuration: testRTDuration,
		Now:                  testNow,
	}

	t.Run("issue", func(t *testing.T) {
		tokens, err := ic.IssueTokens(testUserID, testIsAdmin, testHash)
		ok(t, err)
		testAccessToken = tokens.Access
		testRefreshToken = tokens.Refresh
		equals(t, tokens.ExpiresAt, int(testTime.Add(testATDuration).Unix()))
	})

	t.Run("check access token", func(t *testing.T) {
		ap, err := ic.CheckAccessToken(testAccessToken)
		ok(t, err)
		equals(t, ap.UserID, testUserID)
		equals(t, ap.IsAdmin, testIsAdmin)
	})

	t.Run("check refresh token", func(t *testing.T) {
		rt, err := ic.CheckRefreshToken(testRefreshToken, testHash)
		ok(t, err)
		equals(t, rt.UserID, testUserID)
	})

	t.Run("check access token invalid", func(t *testing.T) {
		ap, err := ic.CheckAccessToken(testTokenWrong1)
		equals(t, gqlmeetup.ErrUnauthorized, err)
		assert(t, ap == nil, "accessTokenPayload should be nil")
	})

	t.Run("check refresh token invalid", func(t *testing.T) {
		rt, err := ic.CheckRefreshToken(testTokenWrong1, testHash)
		equals(t, gqlmeetup.ErrUnauthorized, err)
		assert(t, rt == nil, "accessTokenPayload should be nil")
	})

	t.Run("check refresh token invalidated by changing pwd", func(t *testing.T) {
		rt, err := ic.CheckRefreshToken(testRefreshToken, "wrong")
		equals(t, gqlmeetup.ErrUnauthorized, err)
		assert(t, rt == nil, "accessTokenPayload should be nil")
	})

	t.Run("decode refresh token", func(t *testing.T) {
		id, err := ic.DecodeRefreshToken(testRefreshToken)
		ok(t, err)
		equals(t, testUserID, id)
	})

	t.Run("decode refresh token bad token", func(t *testing.T) {
		id, err := ic.DecodeRefreshToken("bad token")
		equals(t, gqlmeetup.ErrUnauthorized, err)
		assert(t, id == 0, "in case of a bad token returned ID should be 0")
	})
}

func TestStoreAndRetrieve(t *testing.T) {
	t.Parallel()
	ic := jwt.TokenService{Secret: testTokenSecret}
	ctx := context.Background()
	newCtx := ic.Store(ctx, testATPayload)
	ap, err := ic.Retrieve(newCtx)
	ok(t, err)
	equals(t, testATPayload, ap)
}

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
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
