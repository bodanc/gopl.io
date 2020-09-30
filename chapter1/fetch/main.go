// like curl, fetch prints the content found at each specified url;
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	for _, url := range os.Args[1:] {

		// the http.Get() function makes an http request and, if there is no error, returns the result as an
		// http.Response struct;
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		// the Body filed of resp contains the server response as a readable stream, which ioutil.ReadAll() reads;
		b, err := ioutil.ReadAll(resp.Body) // returns a []byte

		// we close the Body stream to avoid any leaking resources;
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("%s", b)

	}

}
