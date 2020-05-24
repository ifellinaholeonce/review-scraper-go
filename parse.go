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
	return nil, nil
}

func dfs(node *html.Node, padding string) {
	msg := node.Data
	if node.Type == html.ElementNode {
		msg = "<" + msg + ">"
	}
	fmt.Println(padding, msg)
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		dfs(child, padding+"  ")
	}
}
