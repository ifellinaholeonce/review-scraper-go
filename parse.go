package review

import (
	"fmt"
	"io"
	"strings"
	"time"

	"golang.org/x/net/html"
)

// Review represents a review DOM element from the Shopify app review page
type Review struct {
	Date   string
	Rating int
	Title  string
	Text   string
}

// Parse will take in an HTML document and return a slice of Reviews parsed
// from it.
func Parse(reader io.Reader) ([]Review, error) {
	doc, err := html.Parse(reader)
	if err != nil {
		return nil, err
	}
	// 1. Find all the .review-listing divs
	nodes := shopifyReviewListing(doc)
	for _, node := range nodes {
		fmt.Println(node)
	}
	// 2. For each review-listing node
	// 3. return the Reviews
	return nil, nil
}

func buildReview(node *html.Node) Review {
	//   2.a get the date
	//   2.b get the rating
	//   2.c get the title
	//   2.d get the text
	var ret Review
	return ret
}

func shopifyReviewListing(node *html.Node) []*html.Node {
	if node.Type == html.ElementNode &&
		node.Data == "div" &&
		strings.TrimSpace(node.Attr[0].Val) == "review-listing" {
		return []*html.Node{node}
	}
	var ret []*html.Node
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		ret = append(ret, shopifyReviewListing(child)...)
	}
	return ret
}

func getDate(node *html.Node) time.Time {
	var date time.Time
	for _, attr := range node.Attr {
		date, err := time.Parse("Jan 02, 2006", "May 22, 2020")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(date)
		fmt.Println(attr)
	}
	return date
}
