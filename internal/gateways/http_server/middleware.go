package http_server

import (
	"log"
	"net/http"
)

func middleware(base http.Handler) http.Handler {
	return func(base http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			log.Println(r.Header.Get("Authorization"))

			w.Header().Set("Content-type", "application/json")
			w.Header().Set("Authorization", "Bearer 123445")
			base.ServeHTTP(w, r)
		})
	}(base)
}