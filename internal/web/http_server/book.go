package http_server

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/twitchtv/twirp"
	"github.com/umkh/twirp_rpc_example/internal/service"
	pb_book "github.com/umkh/twirp_rpc_example/pkg/genproto/services/book"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
)

type Book struct {
	service service.Service
}

func NewBook() (string, http.HandlerFunc) {
	srv := pb_book.NewBookAPIServer(
		&Book{},
		twirp.WithServerPathPrefix(urlPrefix),
	)

	return srv.PathPrefix(), middleware(srv).ServeHTTP
}

func (s *Book) Create(ctx context.Context, req *pb_book.CreateReq) (*empty.Empty, error) {
	if err := req.Validate(); err != nil {
		return nil, twirp.NewError(twirp.InvalidArgument, err.Error())
	}
	return &emptypb.Empty{}, nil
}

func (s *Book) List(ctx context.Context, req *empty.Empty) (*pb_book.ListRes, error) {
	return &pb_book.ListRes{}, nil
}
