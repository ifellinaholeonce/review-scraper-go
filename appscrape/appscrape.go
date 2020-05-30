package appscrape

import (
	"fmt"
	"log"
	"net/http"
	"review-scraper-go/pageparse"
	"strconv"
)

// Scrape scrapes Shopify reviews for an app. Pass in the app name
// to fit https://apps.shopify.com/{{AppName}}/reviews"
func Scrape(appName string) {
	pageCount := 1
	var reviews []pageparse.Review
	finished := false

	for finished != true {
		// Request the HTML page.
		res, err := http.Get(
			"https://apps.shopify.com/" +
				appName +
				"/reviews?page=" +
				strconv.Itoa(pageCount))
		if err != nil {
			log.Fatal(err)
		}

		defer res.Body.Close()

		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}

		var pageReviews []pageparse.Review
		pageReviews, finished, err = pageparse.Parse(res.Body)
		reviews = append(reviews, pageReviews...)
		if err != nil {
			log.Fatal(err)
		}
		if pageCount >= 10 {
			fmt.Println("Breaking")
			finished = true
			break
		}
		if finished {
			fmt.Println("No reviews on page")
		} else {
			fmt.Println("Fetching another page")
			pageCount++
		}
	}

	fmt.Println(reviews)
}
