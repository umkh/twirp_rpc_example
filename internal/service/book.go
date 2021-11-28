package service

import (
	"github.com/umkh/twirp_rpc_example/internal/dto"
	"github.com/umkh/twirp_rpc_example/internal/entities"
	"github.com/umkh/twirp_rpc_example/internal/store"
)

type book struct {
	store store.Store
}

func NewBook(store store.Store) *book {
	return &book{store: store}
}

func (b *book) Create(dto *dto.Book) error {
	return b.store.Book().Create(&entities.Book{
		Name:   dto.Name,
		Author: dto.Author,
	})
}

func (b *book) List() ([]dto.Book, error) {
	return []dto.Book{}, nil
}
