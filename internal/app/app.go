package app

import (
	"context"
	"github.com/umkh/twirp_rpc_example/internal/web/http_server"
	"log"
	"net/http"
	"time"
)

type App struct {
	httpServer *http.Server
}

func New() *App {
	app := new(App)
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
