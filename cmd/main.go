package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/umkh/twirp_rpc_example/internal/app"
)

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	application := app.New()
	defer application.Shutdown()

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-quit
		cancel()
	}()

	application.Start(ctx)
}
