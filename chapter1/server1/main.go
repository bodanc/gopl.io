package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// connect a HANDLER function to incoming URLs whose path begins with '/' which means all URLs;
	http.HandleFunc("/", handler) // each http request will call handler()

	// start an http server listening for incoming requests on port 8000;
	log.Fatal(http.ListenAndServe("localhost:8001", nil))

}

// handler() echoes the Path component of the requested URL;
func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "url.path = %q\n", r.URL.Path)

}
