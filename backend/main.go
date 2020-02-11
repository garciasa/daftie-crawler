package main

import (
	"backend/domain"
	"backend/handlers"
	"backend/postgres"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-pg/pg/v9"
	"github.com/joho/godotenv"
)

func main() {

	env := godotenv.Load()
	if env != nil {
		fmt.Print("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv(("DB_HOST"))

	DB := postgres.New(&pg.Options{
		Addr:      dbHost,
		User:      dbUser,
		Password:  dbPass,
		Database:  dbName,
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	})
	defer DB.Close()

	domainDB := domain.DB{
		HouseRepo: postgres.NewHouseRepo(DB),
		StatRepo:  postgres.NewStatRepo(DB),
	}

	d := &domain.Domain{DB: domainDB}

	/*
		err := d.Parse()
		if err != nil {
			log.Fatalf("cannot parse site %v", err)
		}
	*/
	r := handlers.SetupRouter(d)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	// Parsing every hour
	/*
		c := cron.New()
		c.AddFunc("@every 1h", func() { _ = d.Parse() })
		c.Start()
	*/

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	if err != nil {
		log.Fatalf("cannot start server %v", err)
	}

}
