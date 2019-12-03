package boltdb

import (
	"log"
	"time"

	"github.com/boltdb/bolt"
)

// Connect open database
func Connect(databaseName string) *bolt.DB {
	db, err := bolt.Open(databaseName, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
