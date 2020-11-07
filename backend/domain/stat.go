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

// Statistics including total
type Statistics struct{
	Crawler []Stat `json:"crawler"`
	General struct {
		Total int `json:"total"`
	} `json:"general"`

}

// Chart figures for charting
type Chart struct {
	Year int `json:"year"`
	MonthStr string `json:"month_str"`
	Month int `json:"month"`
	Houses int `json:"houses"`
}

// GetStats get all items in stats
func (d *Domain) GetStats() (*Statistics, error) {
	result:= &Statistics{}
	stat, err := d.DB.StatRepo.GetStats()
	if err != nil {
		return nil, err
	}
	
	total, _ := d.DB.HouseRepo.GetTotalHouses()

	result.Crawler = stat;
	result.General.Total = total
	
	return result, nil
}

// GetStatsForCharts figures for charting
func (d *Domain) GetStatsForCharts(provider string, beds int) ([]Chart, error){
	result := make([]Chart, 0)
	result, err := d.DB.StatRepo.GetStatsForCharts(provider, beds)
	if (err != nil){
		return nil, err
	}
	return result, nil
}
