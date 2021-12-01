package http_server

import (
	"context"
	"github.com/twitchtv/twirp"
	"github.com/umkh/twirp_rpc_example/internal/dto"
	"github.com/umkh/twirp_rpc_example/internal/service"
	pb_book "github.com/umkh/twirp_rpc_example/pkg/genproto/services/book"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Book struct{
	service service.Container
}

func NewBook(service service.Container) pb_book.TwirpServer {
	srv := pb_book.NewBookAPIServer(
		&Book{
			service: service,
		},
		twirp.WithServerPathPrefix(urlPrefix),
	)

	return srv
}

func (s *Book) Create(ctx context.Context, req *pb_book.CreateReq) (*emptypb.Empty, error) {
	if err := req.Validate(); err != nil {
		return nil, twirp.NewError(twirp.InvalidArgument, err.Error())
	}

	if err := s.service.Book().Create(&dto.Book{
		Name: req.Name,
		Author: req.Author,
	}); err != nil {
		return nil, twirp.NewError(twirp.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *Book) List(ctx context.Context, req *emptypb.Empty) (*pb_book.ListRes, error) {
	list, err := s.service.Book().List()
	if err != nil {
		return nil, twirp.NewError(twirp.Internal, err.Error())
	}

	var books []*pb_book.Book

	for _, l := range list {
		books = append(books, &pb_book.Book{
			Name: l.Name,
			Author: l.Author,
		})
	}

	return &pb_book.ListRes{Books: books}, nil
}
