package http_server

import (
	"net/http"
)

func New(port string) *http.Server {
	mux := http.NewServeMux()
	register(mux)

	return &http.Server{
		Addr:    port,
		Handler: mux,
	}
}

func register(mux *http.ServeMux) {
	// Book service
	mux.HandleFunc(NewBook())
}
