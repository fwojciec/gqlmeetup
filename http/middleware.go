package http

import (
	"net/http"

	"github.com/fwojciec/gqlmeetup"
)

// DataloaderMiddleware stores Loaders as a request-scoped context value.
func DataloaderMiddleware(dls gqlmeetup.DataLoaderService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			nextCtx := dls.Initialize(ctx)
			r = r.WithContext(nextCtx)
			next.ServeHTTP(w, r)
		})
	}
}
