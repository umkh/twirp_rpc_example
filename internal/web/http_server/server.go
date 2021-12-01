package http_server

import (
	"github.com/umkh/twirp_rpc_example/internal/config"
	"github.com/umkh/twirp_rpc_example/internal/service"
	"net/http"
)

type Params struct {
	cfg        config.Config
	httpServer *http.Server
	service    service.Container
}

func New(cfg config.Config, service service.Container) *Params {
	return &Params{
		cfg:     cfg,
		service: service,
	}
}

func (p *Params) register(mux *http.ServeMux) {
	// Book service
	book := NewBook(p.service)
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
