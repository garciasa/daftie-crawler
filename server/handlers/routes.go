package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/boltdb/bolt"
)

func (app *App) routes() {
	app.router.HandleFunc("/api/", app.handleAPI())
}

func (app *App) handleAPI() http.HandlerFunc {
	houses := []House{}

	err := app.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("houses"))
		b.ForEach(func(k, v []byte) error {
			item := House{}
			_ = json.Unmarshal(v, &item)
			houses = append(houses, House{
				BrandLink:      string(k),
				Price:          item.Price,
				Date:           item.Date,
				NewDevelopment: item.NewDevelopment,
				Meters:         item.Meters,
				Eircode:        item.Eircode,
			})
			return nil
		})

		return nil
	})
	return func(w http.ResponseWriter, r *http.Request) {
		if err != nil {
			fmt.Fprintf(w, "%s", err)
		}
		json.NewEncoder(w).Encode(houses)
	}
}
