package gqlgen

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/fwojciec/gqlmeetup"
)

// HasRole verifies the user authorization for a particular resource.
func HasRole(auth gqlmeetup.TokenService) func(context.Context, interface{}, graphql.Resolver, Role) (interface{}, error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver, role Role) (interface{}, error) {
		au, err := auth.Retrieve(ctx)
		if err != nil {
			return nil, gqlmeetup.ErrUnauthorized
		}
		if role == RoleAdmin && !au.IsAdmin {
			return nil, gqlmeetup.ErrUnauthorized
		}
		return next(ctx)
	}
}