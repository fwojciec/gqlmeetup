package gqlgen

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

// NewQueryHandler returns a new GraphQL query handler.
func NewQueryHandler(resolver *Resolver) http.Handler {
	return handler.NewDefaultServer(NewExecutableSchema(Config{
		Resolvers: resolver,
		Directives: DirectiveRoot{
			HasRole: HasRole(resolver.Session),
		},
	}))
}

// NewPlaygroundHandler returns a new handler that serves GraphQL Playground.
func NewPlaygroundHandler() func(string) http.Handler {
	return func(endpoint string) http.Handler {
		return playground.Handler("Playground", endpoint)
	}
}
