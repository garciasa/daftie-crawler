package main

import (
	"github.com/boltdb/bolt"
	"github.com/gorilla/mux"
)

// App for crawling daftie
type App struct {
	databaseName string
	db           *bolt.DB
	router       *mux.Router
}

// House Info about House advert
type House struct {
	BrandLink      string `json:"brandlink"`
	Price          string `json:"price"`
	Date           string `json:"date"`
	NewDevelopment bool   `json:"newdevelopment"`
	Meters         string `json:"meters"`
	Eircode        string `json:"eircode"`
}
