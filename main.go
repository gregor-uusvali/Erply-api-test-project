package main

import (
	"Erply-api-test-project/driver"
	"Erply-api-test-project/handlers"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

const (
	portNumber = ":8080"
	dsn        = "./database/database.db?_foreign_keys=on"
)

func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()
	log.Println("Starting application on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", routes()))

}

func run() (*driver.DB, error) {
	log.Println("Connecting to database...")
	db, err := driver.ConnectSQL(dsn)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	log.Println("Connected to database!")

	db.InitDB(dsn)

	repo := handlers.NewRepo(db)
	handlers.NewHandlers(repo)

	return db, nil
}
