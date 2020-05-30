package main

import (
	"fmt"
	"review-scraper-go/appscrape"
)

func main() {
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
	appscrape.Scrape(appName, pageCount)
}
