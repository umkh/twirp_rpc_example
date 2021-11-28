package service

import (
	"github.com/umkh/twirp_rpc_example/internal/dto"
)

type Container interface {
	Book() Book
}

type Book interface {
	Create(dto *dto.Book) error
	List() ([]dto.Book, error)
}
