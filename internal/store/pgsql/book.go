package pgsql

import "github.com/umkh/twirp_rpc_example/internal/entities"

type book struct{}

func NewBook() *book {
	return &book{}
}

func (b *book) Create(e *entities.Book) error {
	return nil
}

func (b *book) List() ([]entities.Book, error) {
	return []entities.Book{}, nil
}
