package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "tymm"
	dbname = "review-scraper-go"
)

// TODO: Using global connection variable for now. Refactor this into some sort
// of dependecy injection into the packages.
var Conn *sql.DB

// Init initializes a connection to the postgres DB and saves that connection in
// a global variable.
func Init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	Conn = db

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database!")
}
