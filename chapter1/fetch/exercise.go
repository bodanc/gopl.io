package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// fetch() checks for the presence of the http:// prefix and appends it if necessary;
// it also uses io.Copy() instead of ioutil.ReadAll() to read the response body;
func fetch() {

	const pf = "http://"

	for _, url := range os.Args[1:] {

		newUrl := ""

		if strings.HasPrefix(url, "http://") {
			continue
		} else {
			newUrl = pf + url
		}

		fmt.Println(newUrl)

		resp, err := http.Get(newUrl)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		// instead of using ioutil.ReadAll(), we can use io.Copy() to copy the response body to os.Stdout without
		// requiring a large buffer to hold the entire stream;
		b, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		resp.Body.Close()

		fmt.Printf("%s", b)

	}

}
