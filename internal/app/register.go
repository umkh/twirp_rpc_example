package app

import (
	"github.com/umkh/twirp_rpc_example/internal/gateways/http_server"
	"github.com/umkh/twirp_rpc_example/internal/gateways/websocket"
)

func (a *App) register() {
	a.mux.HandleFunc("/ws", websocket.WsEndpoint)
	a.mux.HandleFunc(http_server.NewBook())
	a.mux.HandleFunc(http_server.NewAuth())
}
