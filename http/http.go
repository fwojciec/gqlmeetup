package http

import (
	"net"
	"net/http"
	"time"

	"github.com/fwojciec/gqlmeetup"
	"github.com/rs/cors"
)

// Server is the HTTP server serving the GraphQL endpoints.
type Server struct {
	QueryHandler      http.Handler
	PlaygroundHandler func(string) http.HandlerFunc
	TokenService      gqlmeetup.TokenService
	DataLoaderService gqlmeetup.DataLoaderService
}

// Run runs the server.
func (s *Server) Run(ln net.Listener) error {
	mux := http.NewServeMux()
	tm := TokenMiddleware(s.TokenService)
	dm := DataloaderMiddleware(s.DataLoaderService)
	mux.Handle("/query", dm(tm(s.QueryHandler)))
	mux.Handle("/", s.PlaygroundHandler("/query"))
	handler := cors.Default().Handler(mux)
	server := &http.Server{
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
		IdleTimeout:  time.Second,
		Handler:      handler,
	}
	return server.Serve(ln)
}
