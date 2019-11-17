package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robfig/cron/v3"
)

// DOMAIN main domain to use
const DOMAIN = "https://www.daft.ie"

func main() {

	app := &App{
		databaseName: "houses.db",
		router:       mux.NewRouter(),
	}

	app.connect()
	defer app.db.Close()

	app.routes()
	_, _ = app.parse()

	// Parsing every hour
	c := cron.New()
	c.AddFunc("@every 1h", func() { _, _ = app.parse() })
	c.Start()
	// Starting API server
	http.ListenAndServe(":8000", app.router)

}
