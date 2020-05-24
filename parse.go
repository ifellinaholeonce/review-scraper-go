package review

import (
	"fmt"
	"io"

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
	dfs(doc, "")
	// 1. Find all the .review-listing divs
	// 2. For each review-listing node
	//   2.a get the date
	//   2.b get the rating
	//   2.c get the title
	//   2.d get the text
	// 3. return the Reviews
	return nil, nil
}

func shopifyReviewListing(node *html.Node) []*html.Node {
	if node.Type == html.ElementNode && node.Data == "div" {
		return []*html.Node{node}
	}
	var ret []*html.Node
	return ret
}

func dfs(node *html.Node, padding string) {
	msg := node.Data
	// if node.Type == html.ElementNode {
	// 	msg = "<" + msg + ">"
	// }
	fmt.Println(padding, msg)
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		dfs(child, padding+"  ")
	}
}
