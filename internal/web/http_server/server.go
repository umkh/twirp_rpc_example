package http_server

import (
	"github.com/umkh/twirp_rpc_example/internal/config"
	"net/http"
)

type Params struct {
	cfg        config.Config
	httpServer *http.Server
}

func New(cfg config.Config) *Params {
	return &Params{
		cfg: cfg,
	}
}

func (p *Params) register(mux *http.ServeMux) {
	// Book service
	book := NewBook()
	mux.HandleFunc(book.PathPrefix(), middleware(book).ServeHTTP)

}

func (p *Params) Server() *http.Server {
	mux := http.NewServeMux()
	p.register(mux)

	return &http.Server{
		Addr:    p.cfg.HTTPPort,
		Handler: mux,
	}
}
