package domain

import "time"

import "github.com/google/uuid"

// House Info about House advert
type House struct {
	tableName   struct{}  `sql:"house"`
	ID          uuid.UUID `json:"id"`
	URL         string    `json:"url"`
	Price       string    `json:"price"`
	Title       string    `json:"title"`
	Beds        int       `json:"beds"`
	Baths       int       `json:"baths"`
	Provider    string    `json:"provider"`
	Eircode     string    `json:"eircode"`
	DateRenewed time.Time `json:"date_renewed"`
	FirstListed time.Time `json:"first_listed"`
	Propertyid  string    `json:"property_id"`
}

// GetAllHouses bla bla
func (d *Domain) GetAllHouses() ([]House, error) {
	houses, err := d.DB.HouseRepo.GetAllHouses()
	if err != nil {
		return nil, err
	}
	return houses, nil
}

// GetHousesByProvider bla bla
func (d *Domain) GetHousesByProvider(provider string) ([]House, error) {
	houses, err := d.DB.HouseRepo.GetHousesByProvider(provider)
	if err != nil {
		return nil, err
	}
	return houses, nil
}

// GetLastHouses bla bla
func (d *Domain) GetLastHouses() ([]House, error) {
	houses, err := d.DB.HouseRepo.GetLastHouses()
	if err != nil {
		return nil, err
	}
	return houses, nil
}

// Save bla bla
/*
func (d *Domain) Save(house *House) error {
	err := d.DB.HouseRepo.Save(house)
	if err != nil {
		return err
	}
	return nil
}
*/
