package gqlgen

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/fwojciec/gqlmeetup"
)

// HasRole verifies the user authorization for a resource.
func HasRole(session gqlmeetup.SessionService) func(context.Context, interface{}, graphql.Resolver, Role) (interface{}, error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver, role Role) (interface{}, error) {
		user := session.GetUser(ctx)
		if user == nil {
			return nil, gqlmeetup.ErrUnauthorized
		}
		if role == RoleAdmin && !user.Admin {
			return nil, gqlmeetup.ErrUnauthorized
		}
		return next(ctx)
	}
}
