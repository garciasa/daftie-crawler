package boltdb

import (
	"backend/domain"
	"encoding/json"

	"github.com/boltdb/bolt"
)

// HouseRepo in charge of managing houses
type HouseRepo struct {
	DB *bolt.DB
}

// NewHouseRepo initialize HouseRepo
func NewHouseRepo(DB *bolt.DB) *HouseRepo {
	return &HouseRepo{DB: DB}
}

// GetAll all houses in database
func (h *HouseRepo) GetAll() ([]domain.House, error) {
	var houses []domain.House
	err := h.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("houses"))
		b.ForEach(func(k, v []byte) error {
			item := domain.House{}
			_ = json.Unmarshal(v, &item)
			houses = append(houses, domain.House{
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
	if err != nil{
		return nil, err
	}

	return houses, nil
}

// Save persits a house if it's new or something has change on
// the existing one
func (h *HouseRepo) Save(house *domain.House) error {
	err := h.DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("houses"))
		if err != nil {
			return err
		}

		// Check if exits
		v := b.Get([]byte(house.BrandLink))
		if len(v) != 0 {
			item := domain.House{}
			_ = json.Unmarshal(v, &item)
			if house.Price == item.Price {
				// If exists but price is the same then do nothing
				return nil
			}
			//Update price
			item.Price = house.Price
			encoded, err := json.Marshal(item)
			if err != nil {
				return err
			}
			return b.Put([]byte(house.BrandLink), encoded)

		}

		encoded, err := json.Marshal(house)
		if err != nil {
			return err
		}

		return b.Put([]byte(house.BrandLink), encoded)
	})

	return err
}
