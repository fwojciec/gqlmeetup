package gqlgen_test

import (
	"context"
	"errors"
	"testing"

	"github.com/fwojciec/gqlmeetup"
	"github.com/fwojciec/gqlmeetup/gqlgen"
	"github.com/fwojciec/gqlmeetup/mocks"
)

func TestHasRoleAdmin(t *testing.T) {
	t.Parallel()
	var testError = errors.New("test error")
	tests := []struct {
		name        string
		role        gqlgen.Role
		retrieveErr error
		isAdmin     bool
		err         error
	}{
		{"is not user", gqlgen.RoleUser, testError, false, gqlmeetup.ErrUnauthorized},
		{"is user", gqlgen.RoleUser, nil, false, nil},
		{"is not admin", gqlgen.RoleAdmin, nil, false, gqlmeetup.ErrUnauthorized},
		{"is admin", gqlgen.RoleAdmin, nil, true, nil},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			mockAuth := &mocks.TokenServiceMock{
				RetrieveFunc: func(ctx context.Context) (*gqlmeetup.AccessTokenPayload, error) {
					return &gqlmeetup.AccessTokenPayload{UserEmail: "test@email.com", IsAdmin: tc.isAdmin}, tc.retrieveErr
				},
			}
			dir := gqlgen.HasRole(mockAuth)
			resolver := func(ctx context.Context) (res interface{}, err error) { return nil, nil }
			_, err := dir(context.Background(), nil, resolver, tc.role)
			equals(t, tc.err, err)
		})
	}
}
