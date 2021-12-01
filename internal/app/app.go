package app

import (
	"context"
	"github.com/umkh/twirp_rpc_example/internal/config"
	"github.com/umkh/twirp_rpc_example/internal/store"
	"log"
	"net/http"
	"time"
)

type App struct {
	cfg        config.Config
	store      store.Store
	httpServer *http.Server
	shutdown   []func() error
}

func New() *App {
	app := new(App)
	app.cfg = config.Set()

	app.initPostgreSQL()
	app.initHTTPServer()

	return app
}

func (a *App) Start(ctx context.Context) {
	go func() {
		if err := a.httpServer.ListenAndServe(); err != http.ErrServerClosed {
			log.Printf("http server error: %v", err)
		}
	}()
	log.Printf("http server started port: %s", ":8089")

	<-ctx.Done()
}

func (a *App) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := a.httpServer.Shutdown(ctx); err != nil {
		log.Printf("REST Server Graceful Shutdown Failed: %s\n", err)
	}

	for i := 0; i < len(a.shutdown); i++ {
		if err := a.shutdown[len(a.shutdown)-1-i](); err != nil {
			log.Printf("Shutdown error: %v", err)
		}
	}

	log.Println("Gracefully shutdown")
}
