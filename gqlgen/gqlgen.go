package gqlgen

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
)

// NewQueryHandler returns a new GraphQL query handler.
func NewQueryHandler(resolver *Resolver) http.Handler {
	srv := handler.New(NewExecutableSchema(Config{
		Resolvers: resolver,
		Directives: DirectiveRoot{
			HasRole: HasRole(resolver.Session),
		},
	}))
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.POST{})
	srv.SetQueryCache(lru.New(1000))
	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})
	return srv
}

// NewPlaygroundHandler returns a new handler that serves GraphQL Playground.
func NewPlaygroundHandler() func(string) http.Handler {
	return func(endpoint string) http.Handler {
		return playground.Handler("Playground", endpoint)
	}
}
