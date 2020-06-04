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
func Scrape(appName string, optionalMaxPage ...int) {
	pageCount := 1

	pages := make(chan *http.Response, 1)
	for i := 1; i <= pageCount; i++ {
		go asyncFetch(appName, i, pages)
	}
	var reviews []pageparse.Review
	// finished := false

	// for finished != true {
	for res := range pages {
		// Request the HTML page.
		fmt.Println("here")
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
}

func fetchPage(appName string, page int) *http.Response {
	url := "https://apps.shopify.com/" + appName + "/reviews?page=" + strconv.Itoa(page)
	fmt.Println(url)
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	return res
}

func asyncFetch(appName string, pageCount int, ch chan<- *http.Response) {
	ch <- fetchPage(appName, pageCount)
}
