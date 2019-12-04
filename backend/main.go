package main

import "backend/boltdb"
import "backend/handlers"
import "os"
import "net/http"
import "fmt"
import "log"

func main() {
	DB := boltdb.Connect("houses.db")

	defer DB.Close()

	r := handlers.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	if err != nil {
		log.Fatalf("cannot start server %v", err)
	}

}
