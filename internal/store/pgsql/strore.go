package pgsql

import (
	"github.com/jmoiron/sqlx"
	"github.com/umkh/twirp_rpc_example/internal/store"
)

type pgsql struct {
	db   *sqlx.DB
	book *book
}

func New(db *sqlx.DB) *pgsql {
	return &pgsql{db: db}
}

func (p *pgsql) Book() store.Book {
	if p.book == nil {
		p.book = NewBook(p.db)
	}
	return p.book
}
