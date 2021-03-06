package postgres

import (
	"backend/domain"

	"github.com/go-pg/pg/v9"
)

// HouseRepo type
type HouseRepo struct {
	DB *pg.DB
}

// GetAllHouses get all scraped houses from db
func (h *HouseRepo) GetAllHouses() ([]domain.House, error) {
	houses := make([]domain.House, 0)
	err := h.DB.Model(&houses).
		Where("date_renewed is not null").
		Where("date_renewed <= current_date - interval '7 days'").
		Order("date_renewed DESC").
		Select()
	if err != nil {
		return nil, err
	}
	housesNotDate := make([]domain.House, 0)
	err = h.DB.Model(&housesNotDate).
		Where("date_renewed is null"). 
		Where("provider='myhome.ie'").
		Order("propertyid DESC").
		Select()
	if err != nil {
		return nil, err
	}

	houses = append(houses, housesNotDate...)

	return houses, nil
}

// GetHousesByProvider get the houses based on website provider
func (h *HouseRepo) GetHousesByProvider(provider string) ([]domain.House, error) {
	// var houses []domain.House
	houses := make([]domain.House, 0)
	err := h.DB.Model(&houses).
		Where("provider = ?", provider).
		Order("date_renewed DESC").
		Select()
	if err != nil {
		return nil, err
	}

	return houses, nil
}

//GetLastHouses Added in the last 7 days
func (h *HouseRepo) GetLastHouses() ([]domain.House, error) {
	houses := make([]domain.House, 0)
	err := h.DB.Model(&houses).
		Where("date_renewed > current_date - interval '7 days'").
		Order("date_renewed DESC").
		Select()
	if err != nil {
		return nil, err
	}

	return houses, nil
}

// GetTotalHouses Total parsed
func (h *HouseRepo) GetTotalHouses() (int, error){
	total, err := h.DB.Model(&domain.House{}).Count()
	if err != nil {
		return 0, err
	}

	return total, nil
}

// GetHousesPerPage Houses per page
func (h *HouseRepo) GetHousesPerPage(page int)([]domain.House, error){
	filter := new(domain.HouseFilter)

	houses := make([]domain.House, 0)
	err := h.DB.Model(&houses).
		WhereStruct(filter).
		Limit(10).
		Offset(page).
		Order("date_renewed DESC").
		Select()
		 
	if err != nil {
		return nil, err
	}

	return houses, nil
}

// NewHouseRepo return intialized repo
func NewHouseRepo(DB *pg.DB) *HouseRepo {
	return &HouseRepo{DB: DB}
}
