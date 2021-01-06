package main

import (
	"fmt"
	"github.com/bodanc/gopl.io/chapter5/findlinks1"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
)

func main() {

/*	for _, url := range os.Args[1:] {
		links, err := findLinks2(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "the function findlinks2 encountered an error: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}*/

	for _, url := range os.Args[1:] {
		links, err := findLinks2Log(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "the function findLinks2 encountered an error: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}

}

// findlinks2 performs an HTTP GET request for the url, parses the response as HTML,
// and extracts and returns the links;
func findLinks2(url string) ([]string, error) { // returns the list of discovered links
	resp, err := http.Get(url) // resp is an http.Response type
	if err != nil {
		return nil, err // a nil []string
	}

	if resp.StatusCode != http.StatusOK {
		// we must ensure that resp.Body is closed so that network resources are properly released,
		// even in case of error;
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		// constructs a new error message containing relevant information: where did the error occur and the url of
		// the document being parsed;
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	return findlinks1.Visit(nil, doc), nil
}

// the multi-valued result returned by calling findLinks2() can be returned from a multi-valued calling function;
func findLinks2Log(url string) ([]string, error) {
	log.Printf("findLinks2 %s", url)
	return findLinks2(url)
}
