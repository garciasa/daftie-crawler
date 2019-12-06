package domain

// House Info about House advert
type House struct {
	BrandLink      string `json:"brandlink"`
	Price          string `json:"price"`
	Date           string `json:"date"`
	NewDevelopment bool   `json:"newdevelopment"`
	Meters         string `json:"meters"`
	Eircode        string `json:"eircode"`
}

// GetAll bla bla
func (d *Domain) GetAll()([]House, error){
	houses, err := d.DB.HouseRepo.GetAll()
	if err != nil{
		return nil, err
	}
	return houses, nil
}


// Save bla bla
func (d *Domain) Save(house *House) error {
	err := d.DB.HouseRepo.Save(house)
	if err != nil {
		return err
	}
	return nil
}