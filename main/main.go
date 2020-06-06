package main

import (
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
	fmt.Println("How many pages? (0 for all pages)")
	_, err = fmt.Scan(&pageCount)
	if err != nil {
		panic(err)
	}

	reviews := appscrape.Scrape(appName, pageCount)
	if len(reviews) > 0 {
		success := saveReviewRecords(reviews)
		if !success {
			log.Fatal("Failed to persist review records")
		}
	}
}

func saveReviewRecords(records []pageparse.Review) bool {
	err := db.Conn.Ping()
	if err != nil {
		panic(err)
	}

	for _, record := range records {
		sqlStatement := `
		INSERT INTO reviews (rating, title, text)
		VALUES ($1, $2, $3)`

		_, err := db.Conn.Exec(sqlStatement, record.Rating, record.Title, record.Text)
		if err != nil {
			panic(err)
		}
	}
	return true
}
