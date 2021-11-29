package app

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/umkh/twirp_rpc_example/internal/config"
	"github.com/umkh/twirp_rpc_example/internal/web/http_server"
)

type App struct {
	cfg        config.Config
	db         *sqlx.DB
	httpServer *http.Server
}

func New() *App {
	app := new(App)
	app.cfg = config.Set()

	app.initPostreSQL()

	app.httpServer = http_server.New(":8089")
	return app
}

func (a *App) Start(ctx context.Context) {
	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
	log.Printf("http server started port: %s", ":8089")

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	if err := a.httpServer.Shutdown(ctx); err != nil {
		log.Printf("REST Server Graceful Shutdown Failed: %s\n", err)
	}

	log.Println("Gracefully shutdown")
}
