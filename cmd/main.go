package main

import (
	"context"
	"github.com/umkh/twirp_rpc_example/internal/app"
	"os"
	"os/signal"
)

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	application := app.New()
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-quit
		cancel()
	}()

	application.Run(ctx)
}
