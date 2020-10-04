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

	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)

	log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil))

}

// handler() echoes the http request;
func handler(w http.ResponseWriter, r *http.Request) {

	m.Lock()
	count++
	m.Unlock()

	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

	for k, v := range r.Header {
		fmt.Fprintf(w, "header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "host = %q\n", r.Host)
	fmt.Fprintf(w, "remote address = %q\n", r.RemoteAddr)

	// combining the statements is shorter and reduces the scope of the err variable;
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "form[%q] = %q\n", k, v)
	}

}

func counter(w http.ResponseWriter, r *http.Request) {

	m.Lock()
	fmt.Fprintf(w, "count: %d\n", count)
	m.Unlock()

}
