package app

import (
	"log"

	"github.com/jmoiron/sqlx"
)

const sqlDriver = "postgres"

func (a *App) initPostreSQL() {
	conn, err := sqlx.Connect(sqlDriver, a.cfg.PostgreSQLURL())
	if err != nil {
		log.Panicf("PostgreSQL connect error: %v", err)
	}

	a.db = conn
}
