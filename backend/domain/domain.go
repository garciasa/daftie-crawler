package domain

// HouseRepo bla bla
type HouseRepo interface {
	GetAllHouses() ([]House, error)
	GetHousesByProvider(string) ([]House, error)
	GetLastHouses() ([]House, error)
}

// StatRepo bla bla
type StatRepo interface {
	GetStats() ([]Stat, error)
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
