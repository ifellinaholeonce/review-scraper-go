package main

import (
	"fmt"
	"review-scraper-go/appscrape"
)

func main() {
	var appName string
	fmt.Println("What is the name of the app?")
	_, err := fmt.Scan(&appName)
	if err != nil {
		panic(err)
	}
	appscrape.Scrape(appName)
}
