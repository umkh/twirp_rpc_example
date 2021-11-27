package http_server

import (
	"context"
	"github.com/twitchtv/twirp"
	pb_book "github.com/umkh/twirp_rpc_example/pkg/genproto/services/book"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
)

type Book struct{}

func NewBook() (string, http.HandlerFunc) {
	srv := pb_book.NewBookAPIServer(
		&Book{},
		twirp.WithServerPathPrefix(urlPrefix),
	)

	return srv.PathPrefix(), middleware(srv).ServeHTTP
}

func (s *Book) Create(ctx context.Context, req *pb_book.CreateReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *Book) Get(ctx context.Context, req *pb_book.GetReq) (*pb_book.Book, error) {
	return &pb_book.Book{
		Name:   "Golang advanced",
		Author: "Umidjon Mukhtorov",
	}, nil
}
