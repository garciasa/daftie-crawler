package main

import (
	"backend/boltdb"
	"backend/domain"
	"backend/handlers"
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"net/http"
	"os"
)

func main() {
	DB := boltdb.Connect("houses.db")

	defer DB.Close()

	domainDB := domain.DB{
		HouseRepo: boltdb.NewHouseRepo(DB),
	}

	d := &domain.Domain{DB: domainDB}

	err := d.Parse()
	if err != nil {
		log.Fatalf("cannot parse site %v", err)
	}

	r := handlers.SetupRouter(d)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Parsing every hour
	c := cron.New()
	c.AddFunc("@every 1h", func() { _ = d.Parse() })
	c.Start()

	err = http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	if err != nil {
		log.Fatalf("cannot start server %v", err)
	}

}
