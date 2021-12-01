package app

import (
	"github.com/umkh/twirp_rpc_example/internal/service"
	"github.com/umkh/twirp_rpc_example/internal/store/pgsql"
	"github.com/umkh/twirp_rpc_example/internal/web/http_server"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const sqlDriver = "postgres"

func (a *App) initPostgreSQL() {
	conn, err := sqlx.Connect(sqlDriver, a.cfg.PostgreSQLURL())
	if err != nil {
		log.Panicf("PostgreSQL connect error: %v", err)
	}
	a.shutdown = append(a.shutdown, conn.Close)

	a.store = pgsql.New(conn)
}

func (a *App) initService() {
	a.service = service.New(a.store)
}

func (a *App) initHTTPServer() {
	a.httpServer = http_server.New(a.cfg, a.service).Server()
}
