package pageparse

import (
	"io"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Review represents a review DOM element from the Shopify app review page
type Review struct {
	Date   time.Time
	Rating int
	Title  string
	Text   string
}

// Parse will take in an HTML document and return a slice of Reviews parsed
// from it.
func Parse(reader io.Reader) ([]Review, bool, error) {
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatal(err)
	}

	var reviews []Review
	// Find the review items
	doc.Find(".review-listing").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		reviews = append(reviews, buildReview(s))
	})

	// If there were reviews on the page return false
	// for the completed boolean so that we know to check
	// the next page of reviews
	var completed bool
	if len(reviews) > 0 {
		completed = false
	}

	return reviews, completed, nil
}

func buildReview(element *goquery.Selection) Review {
	var ret Review
	rating, date := parseReviewMetaData(element)
	ret.Title = element.Find("h3.review-listing-header__text").First().Text()
	ret.Text = element.Find(".truncate-content-copy").Text()

	ret.Rating = rating
	ret.Date = date
	return ret
}

func parseReviewMetaData(element *goquery.Selection) (int, time.Time) {
	var rating int
	var date time.Time
	element.Find(".review-metadata").Children().Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			strRating := s.Find(".ui-star-rating").First().AttrOr("data-rating", "0")
			intRating, err := strconv.Atoi(strRating)
			if err != nil {
				panic(err)
			}
			rating = intRating
		}
		if i == 1 {
			dateStr := s.Find(".review-metadata__item-value").First().Text()
			date = parseDate(dateStr)
		}
	})
	return rating, date
}

func parseDate(dateStr string) time.Time {
	date, err := time.Parse("January 2, 2006", strings.TrimSpace(dateStr))
	if err != nil {
		panic(err)
	}
	return date
}
