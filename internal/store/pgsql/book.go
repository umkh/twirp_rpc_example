package pgsql

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/umkh/twirp_rpc_example/internal/entities"
	"log"
)

const bookTable = "books"

type book struct {
	db *sqlx.DB
}

func NewBook(db *sqlx.DB) *book {
	return &book{db: db}
}

func (b *book) Create(e *entities.Book) error {
	tx := b.Tx()

	query := fmt.Sprintf("INSERT INTO %s (name, author) VALUES($1, $2)", bookTable)

	if _, err := tx.Exec(query, e.Name, e.Author); err != nil {
		log.Println(err)
		if err := tx.Rollback(); err != nil {
			log.Println(err)
			return err
		}
	}

	return tx.Commit()
}

func (b *book) List() ([]entities.Book, error) {
	var result []entities.Book

	query := fmt.Sprintf("SELECT * FROM %s", bookTable)

	if err := b.db.Select(&result, query); err != nil {
		log.Println(err)
		return result, err
	}

	return result, nil
}

func (b *book) Tx() *sql.Tx {
	tx, err := b.db.Begin()
	if err != nil {
		log.Println(err)
	}

	return tx
}
