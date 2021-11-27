package app

import (
	"context"
	"log"
	"net/http"
	"time"
)

type App struct {
	mux        *http.ServeMux
	httpServer *http.Server
}

func New() *App {
	app := new(App)
	app.initHttp()
	return app
}

func (a *App) initHttp() {
	a.mux = http.NewServeMux()
	a.register()
	a.httpServer = &http.Server{
		Addr:    ":8089",
		Handler: a.mux,
	}
}

func (a *App) Run(ctx context.Context) {
	go func() {
		if err := http.ListenAndServe(":8089", a.mux); err != nil {
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
