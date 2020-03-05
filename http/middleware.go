package http

import (
	"net/http"
	"strings"

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

// TokenMiddleware stores the user email and admin as a request-scoped context
// value if the authorization was successful.
func TokenMiddleware(ts gqlmeetup.TokenService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			h := r.Header.Get("Authorization")
			if h == "" {
				next.ServeHTTP(w, r)
				return
			}
			t := strings.TrimPrefix(h, "Bearer ")
			ap, err := ts.CheckAccessToken(t)
			if err != nil || ap == nil {
				next.ServeHTTP(w, r)
				return
			}
			prevCtx := r.Context()
			nextCtx := ts.Store(prevCtx, ap)
			r = r.WithContext(nextCtx)
			next.ServeHTTP(w, r)
		})
	}
}
