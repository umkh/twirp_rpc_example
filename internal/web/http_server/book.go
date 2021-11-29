package http_server

import (
	"context"
	"net/http"

	"github.com/twitchtv/twirp"
	pb_book "github.com/umkh/twirp_rpc_example/pkg/genproto/services/book"
	"google.golang.org/protobuf/types/known/emptypb"
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
	if err := req.Validate(); err != nil {
		return nil, twirp.NewError(twirp.InvalidArgument, err.Error())
	}
	return &emptypb.Empty{}, nil
}

func (s *Book) List(ctx context.Context, req *emptypb.Empty) (*pb_book.ListRes, error) {
	return &pb_book.ListRes{}, nil
}
