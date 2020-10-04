// a minimal server that echoes url paths and provides a counter for a specific url;

package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var m sync.Mutex
var count int


func main() {

	// a handler pattern that ends with a slash matches any url that has the pattern as a prefix;
	http.HandleFunc("/count", counter)
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

// handler() echoes the path component of the requested url;
func handler(w http.ResponseWriter, r *http.Request) {

	m.Lock()
	count++
	m.Unlock()

	fmt.Fprintf(w, "url.path = %q\n", r.URL.Path)

}

// counter() echoes the number of calls so far;
func counter(w http.ResponseWriter, r *http.Request) {

	m.Lock()
	fmt.Fprintf(w, "count: %d\n", count)
	m.Unlock()

}
