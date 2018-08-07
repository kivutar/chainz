package handler

import (
	"net/http"

	"golang.org/x/net/context"
)

func AddContext(ctx context.Context, h http.Handler) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
