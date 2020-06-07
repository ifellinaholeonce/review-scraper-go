package main

import (
	"database/sql"
	"fmt"
	"log"
	"review-scraper-go/appscrape"
	"review-scraper-go/db"
	"review-scraper-go/pageparse"
)

func main() {
	db.Init()
	defer db.Conn.Close()

	var appName string
	var pageCount int

	fmt.Println("What is the name of the app?")
	_, err := fmt.Scan(&appName)
	if err != nil {
		panic(err)
	}
	appID := findOrCreateApp(appName)

	fmt.Println("How many pages? (0 for all pages)")
	_, err = fmt.Scan(&pageCount)
	if err != nil {
		panic(err)
	}

	reviews := appscrape.Scrape(appName, pageCount)
	if len(reviews) > 0 {
		success := saveReviewRecords(reviews, appID)
		if !success {
			log.Fatal("Failed to persist review records")
		}
	}
}

func saveReviewRecords(records []pageparse.Review, appID int) bool {
	for _, record := range records {
		sqlStatement := `
		INSERT INTO reviews (rating, title, text, app_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id`

		_, err := db.Conn.Exec(
			sqlStatement,
			record.Rating,
			record.Title,
			record.Text,
			appID)
		if err != nil {
			panic(err)
		}
	}
	return true
}

func findOrCreateApp(appName string) int {
	var id int

	sqlStatement := `SELECT id FROM apps WHERE name ILIKE $1`

	row := db.Conn.QueryRow(sqlStatement, appName)

	switch err := row.Scan(&id); err {
	case sql.ErrNoRows:
		fmt.Println("App not found, creating app record.")
		id = createAppRecord(appName)
	case nil:
		fmt.Println(id)
	default:
		panic(err)
	}
	return id
}

func createAppRecord(appName string) int {
	var id int

	sqlStatement := `
		INSERT INTO apps (name)
		VALUES($1)
		RETURNING id`
	err := db.Conn.QueryRow(sqlStatement, appName).Scan(&id)
	if err != nil {
		panic(err)
	}
	return id
}
