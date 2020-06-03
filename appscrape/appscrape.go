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

	avg := calcAvg(scores)
	median := calcMedian(scores)
	fmt.Println("The average is", avg)
	fmt.Println("median", median)
}

func calcAvg(scores []int) float32 {
	var avg float32
	sum := sum(scores...)
	avg = float32(sum) / float32(len(scores))
	fmt.Println("scores sum", sum)
	fmt.Println("count", len(scores))

	return avg
}

func sum(input ...int) int {
	sum := 0

	for i := range input {
		sum += input[i]
	}

	return sum
}

func calcMedian(scores []int) int {
	var median int
	middle := len(scores) / 2
	if len(scores)%2 == 0 {
		median = (scores[middle-1] + scores[middle]) / 2
	} else {
		median = scores[middle]
	}
	return median
}
