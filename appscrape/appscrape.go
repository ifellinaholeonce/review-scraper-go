package appscrape

import (
	"fmt"
	"log"
	"net/http"
	"review-scraper-go/pageparse"
)

// URLScrape scrapes Shopify reviews for an app. Pass in the app name
// to fit https://apps.shopify.com/{{AppName}}/reviews"
func URLScrape(AppName string) {
	// Request the HTML page.
	res, err := http.Get("https://apps.shopify.com/" + AppName + "/reviews")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	reviews, err := pageparse.Parse(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reviews)
}
