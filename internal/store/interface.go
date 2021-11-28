package store

import "github.com/umkh/twirp_rpc_example/internal/entities"

type Store interface {
	Book() Book
}

type Book interface {
	Create(e *entities.Book) error
	List() ([]entities.Book, error)
}
