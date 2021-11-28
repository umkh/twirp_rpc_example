package pgsql

import "github.com/umkh/twirp_rpc_example/internal/store"

type pgsql struct {
	book *book
}

func New() *pgsql {
	return &pgsql{}
}

func (p *pgsql) Book() store.Book {
	if p.book == nil {
		p.book = NewBook()
	}
	return p.book
}
