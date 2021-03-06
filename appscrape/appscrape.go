package appscrape

import (
	"fmt"
	"log"
	"net/http"
	"review-scraper-go/calculate"
	"review-scraper-go/pageparse"
	"strconv"
)

// Scrape scrapes Shopify reviews for an app. Pass in the app name
// to fit https://apps.shopify.com/{{AppName}}/reviews"
func Scrape(appName string, maxPage int) []pageparse.Review {
	MAX_CONCURRENT := 10
	var pageCount int
	if maxPage == 0 {
		// Detect max page count
		pageCount = calcTotalPages(appName)
	} else {
		pageCount = maxPage
	}

	var results []*http.Response
	pages := make(chan int, MAX_CONCURRENT)
	done := make(chan bool)

	go func() {
		for {
			page, more := <-pages
			if more {
				results = append(results, asyncFetch(appName, page))
				fmt.Println("received page number", page)
			} else {
				done <- true
				return
			}
		}
	}()

	for i := 1; i <= pageCount; i++ {
		fmt.Println("queueing page number", i)
		pages <- i
	}

	close(pages)
	<-done

	var reviews []pageparse.Review

	for _, res := range results {
		// Request the HTML page.
		var pageReviews []pageparse.Review
		pageReviews, err := pageparse.Parse(res.Body)
		defer res.Body.Close()
		reviews = append(reviews, pageReviews...)
		if err != nil {
			log.Fatal(err)
		}
	}

	var scores []int
	for _, review := range reviews {
		scores = append(scores, review.Rating)
	}

	fmt.Println("The average is", calculate.CalcAvg(scores))
	fmt.Println("median", calculate.CalcMedian(scores))
	return reviews
}

func fetchPage(appName string, page int) *http.Response {
	url := "https://apps.shopify.com/" + appName + "/reviews?page=" + strconv.Itoa(page)
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	return res
}

func asyncFetch(appName string, pageCount int) *http.Response {
	return fetchPage(appName, pageCount)
}

func calcTotalPages(appName string) int {
	// Grab first page
	res := fetchPage(appName, 1)
	// Parse the total review count
	total := (pageparse.ParseReviewsTotalCount(res.Body) / 10) + 1
	fmt.Println("total pages is", total)
	return total
}
