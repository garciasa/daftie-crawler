package domain

// HouseRepo bla bla
type HouseRepo interface {
	GetAllHouses() ([]House, error)
	GetHousesByProvider(string) ([]House, error)
	GetHousesPerPage(int)([]House, error)
	GetLastHouses() ([]House, error)
	GetTotalHouses() (int, error)
}

// StatRepo bla bla
type StatRepo interface {
	GetStats() ([]Stat, error)
	GetStatsForCharts(string, int)([]Chart, error)
}

// DB database struct
type DB struct {
	HouseRepo HouseRepo
	StatRepo  StatRepo
}

//Domain main business logic
type Domain struct {
	DB DB
}
