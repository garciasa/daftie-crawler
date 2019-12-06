package domain

// HouseRepo bla bla
type HouseRepo interface {
	GetAll() ([]House, error)
	Save(house *House) error
}

// DB database struct
type DB struct {
	HouseRepo HouseRepo
}

//Domain main business logic
type Domain struct {
	DB DB
}
