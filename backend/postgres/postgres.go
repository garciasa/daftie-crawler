package postgres

import (
	"log"

	"github.com/go-pg/pg/v9"
	// pq driver is loaded in this way
	_ "github.com/lib/pq"
)

// New Postgres connection
func New(opts *pg.Options) *pg.DB {
	db := pg.Connect(opts)

	if db == nil {
		log.Fatal("Error connecting to db...")
	}

	return db
}
