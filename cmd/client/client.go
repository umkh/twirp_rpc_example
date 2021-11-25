package main

import (
	"context"
	"fmt"
	"github.com/twitchtv/twirp"
	pb_book "github.com/umkh/twirp_rpc_example/pkg/genproto/services/book"
	"log"
	"net/http"
)

func main() {
	client := pb_book.NewBookAPIJSONClient(
		"http://localhost:8089",
		&http.Client{},
		twirp.WithClientPathPrefix("/api/v1"),
	)

	book, err := client.Get(context.Background(), &pb_book.GetReq{
		Id: "1234",
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(book)
}
