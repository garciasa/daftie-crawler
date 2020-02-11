package domain

import "time"

// Stat model
type Stat struct {
	tableName struct{}  `sql:"stat"`
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

// GetStats get all items in stats
func (d *Domain) GetStats() ([]Stat, error) {
	stat, err := d.DB.StatRepo.GetStats()
	if err != nil {
		return nil, err
	}

	return stat, nil
}
