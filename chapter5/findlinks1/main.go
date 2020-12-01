// findlinks1 prints the links in an html document read from standard input;
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

// visit() traverses an HTML node tree, extracts the link from the href attribute of each anchor element,
// appends the links to a slice of strings, and returns the resulting slice:
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	// to descend the tree for a node n, visit() recursively calls itself for each of n's children,
	// which are held in the FirstChild linked list;
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}

	return links
}

func main() {

	// html.Parse() reads a sequence of bytes, parses them, and returns the root of the HTML document tree,
	// which is an html.Node
	document, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	for _, link := range visit(nil, document) {
		fmt.Println(link)
	}

}
