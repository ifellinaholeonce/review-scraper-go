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
func Scrape(appName string, optionalMaxPage ...int) {
	var maxPageCount int
	if len(optionalMaxPage) > 0 {
		maxPageCount = optionalMaxPage[0]
	} else {
		maxPageCount = 1
	}

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
		if pageCount == maxPageCount {
			finished = true
		}
		if !finished {
			fmt.Println("Fetching another page")
			pageCount++
		}
	}

	var scores []int
	for _, review := range reviews {
		scores = append(scores, review.Rating)
	}

	avg, _, _ := avgMedMode(scores)
	fmt.Println("The average is", avg)
}

func avgMedMode(scores []int) (int, int, int) {
	sum := sum(scores...)
	avg := sum / len(scores)

	return avg, 0, 0
}

func sum(input ...int) int {
	sum := 0

	for i := range input {
		sum += input[i]
	}

	return sum
}
