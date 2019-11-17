package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

func (app *App) connect() {
	db, err := bolt.Open(app.databaseName, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	app.db = db

}

func (h *House) save(db *bolt.DB) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("houses"))
		if err != nil {
			return err
		}

		// Check if exits
		v := b.Get([]byte(h.BrandLink))
		if len(v) != 0 {
			item := House{}
			_ = json.Unmarshal(v, &item)
			if h.Price == item.Price {
				// If exists but price is the same then do nothing
				return nil
			}
			//Update price
			item.Price = h.Price
			encoded, err := json.Marshal(item)
			if err != nil {
				return err
			}
			return b.Put([]byte(h.BrandLink), encoded)

		}

		encoded, err := json.Marshal(h)
		if err != nil {
			return err
		}

		return b.Put([]byte(h.BrandLink), encoded)

	})
	return err
}
