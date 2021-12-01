package service

import (
	"github.com/umkh/twirp_rpc_example/internal/dto"
	"github.com/umkh/twirp_rpc_example/internal/entities"
)

type Container interface {
	Book() Book
}

type Book interface {
	Create(dto *dto.Book) error
	List() ([]entities.Book, error)
}
