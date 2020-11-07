package postgres

import (
	"backend/domain"

	"github.com/go-pg/pg/v9"
)

// StatRepo type
type StatRepo struct {
	DB *pg.DB
}

// GetStats get all stats for spiders
func (s *StatRepo) GetStats() ([]domain.Stat, error) {
	// avoid return null if empty array
	stats := make([]domain.Stat, 0)
	err := s.DB.Model(&stats).Select()
	if err != nil {
		return nil, err
	}

	return stats, nil
}

// GetStatsForCharts stats for charts
func(s *StatRepo) GetStatsForCharts(provider string, beds int)([]domain.Chart, error){
	results := make([]domain.Chart, 0)
	_, err := s.DB.Query(&results, ` 
	select extract(YEAR from date_renewed) as year, to_char(date_renewed,'Mon') as month_str, extract(MONTH from date_renewed) as month, count(*) as houses from house
	where provider = 'myhome.ie' and beds = ?
	group by year, month_str, month
	order by year,month
	`, beds)

	if err != nil {
		return nil, err
	}

	return results, nil
}

// NewStatRepo initiazlizing Stat repo
func NewStatRepo(DB *pg.DB) *StatRepo {
	return &StatRepo{DB: DB}
}
