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

// NewStatRepo initiazlizing Stat repo
func NewStatRepo(DB *pg.DB) *StatRepo {
	return &StatRepo{DB: DB}
}
