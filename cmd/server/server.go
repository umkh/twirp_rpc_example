package main

import (
	"context"
	"github.com/twitchtv/twirp"
	"github.com/umkh/twirp_rpc_example/internal/book"
	pb_book "github.com/umkh/twirp_rpc_example/pkg/genproto/services/book"
	"log"
	"net/http"
)

func WithJsonBase(base http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Content-Type", "application/json")
		base.ServeHTTP(w, r)
	})
}

func main() {
	server := &book.Server{}
	handler := pb_book.NewBookAPIServer(
		server,
		twirp.WithServerPathPrefix("/api/v1"),
		twirp.WithServerHooks(NewLoggingServerHooks()),
	)
	h := WithJsonBase(handler)
	mux := http.NewServeMux()
	mux.Handle(handler.PathPrefix(), h)

	log.Fatal(http.ListenAndServe(":8089", mux))
}

// NewLoggingServerHooks logs request and errors to stdout in the service
func NewLoggingServerHooks() *twirp.ServerHooks {
	return &twirp.ServerHooks{
		RequestRouted: func(ctx context.Context) (context.Context, error) {
			sr, _ := twirp.ServiceName(ctx)
			log.Println("Service: " + sr)
			method, _ := twirp.MethodName(ctx)
			log.Println("Method: " + method)
			return ctx, nil
		},
		Error: func(ctx context.Context, twerr twirp.Error) context.Context {
			log.Println("Error: " + string(twerr.Code()))
			return ctx
		},
		ResponseSent: func(ctx context.Context) {
			log.Println("Response Sent (error or success)")
		},
	}
}
