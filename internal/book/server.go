package book

import (
	"context"
	pb_book "github.com/umkh/twirp_rpc_example/pkg/genproto/services/book"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
}

func (s *Server) Create(ctx context.Context, req *pb_book.CreateReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *Server) Get(ctx context.Context, req *pb_book.GetReq) (*pb_book.Book, error) {
	return &pb_book.Book{
		Name:   "Golang advanced",
		Author: "Umidjon Mukhtorov",
	}, nil
}
