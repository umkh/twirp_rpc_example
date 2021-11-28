package service

import "github.com/umkh/twirp_rpc_example/internal/store"

type Service struct {
	book *book
}

func New(store store.Store) *Service {
	return &Service{
		book: NewBook(store),
	}
}

func (s *Service) Book() Book {
	return s.book
}
