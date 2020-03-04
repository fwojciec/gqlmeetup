package bcrypt_test

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/fwojciec/gqlmeetup"
	"github.com/fwojciec/gqlmeetup/bcrypt"
)

func TestCheckerHasher(t *testing.T) {
	t.Parallel()

	var (
		testPwd  = "testPwd"
		testHash string
	)

	ch := &bcrypt.PasswordService{}

	t.Run("hash", func(t *testing.T) {
		h, err := ch.Hash(testPwd)
		ok(t, err)
		testHash = h
	})

	t.Run("compare ok", func(t *testing.T) {
		err := ch.Check(string(testHash), testPwd)
		ok(t, err)
	})

	t.Run("compare not ok", func(t *testing.T) {
		err := ch.Check("123", testPwd)
		equals(t, gqlmeetup.ErrPwdCheck, err)
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
