package http_server

import (
	"net/http"
)

func middleware(base http.Handler) http.Handler {
	return func(base http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-type", "application/json")
			base.ServeHTTP(w, r)
		})
	}(base)
}
