package gqlgen_test

import (
	"context"
	"testing"

	"github.com/fwojciec/gqlmeetup"
	"github.com/fwojciec/gqlmeetup/gqlgen"
	"github.com/fwojciec/gqlmeetup/mocks"
)

func TestHasRoleDirective(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		role    gqlgen.Role
		isUser  bool
		isAdmin bool
		err     error
	}{
		{"is not user", gqlgen.RoleUser, false, false, gqlmeetup.ErrUnauthorized},
		{"is user", gqlgen.RoleUser, true, false, nil},
		{"is not admin", gqlgen.RoleAdmin, true, false, gqlmeetup.ErrUnauthorized},
		{"is admin", gqlgen.RoleAdmin, true, true, nil},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sessionMock := &mocks.SessionServiceMock{
				GetUserFunc: func(ctx context.Context) *gqlmeetup.User {
					if !tc.isUser {
						return nil
					}
					return &gqlmeetup.User{
						Email: "test@email.com",
						Admin: tc.isAdmin,
					}
				},
			}
			dir := gqlgen.HasRole(sessionMock)
			resolver := func(ctx context.Context) (res interface{}, err error) { return nil, nil }
			_, err := dir(context.Background(), nil, resolver, tc.role)
			equals(t, tc.err, err)
		})
	}
}
