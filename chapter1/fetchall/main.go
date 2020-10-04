package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {

	start := time.Now()

	// the main() function, which runs in a goroutine, creates a channel of strings;
	ch := make(chan string)

	// for each command line argument, the go statement starts a new goroutine that calls fetch() asynchronously;
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // starts a goroutine
	}

	// the loop reads the lines written to the 'ch' channel and prints them;
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from the 'ch' channel
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}

func fetch(url string, ch chan<- string) {

	start := time.Now()

	response, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send the error to channel 'ch'
		return
	}

	// the io.Copy() function reads the body of the response and discards it by writing to the ioutil.Discard output stream;
	// io.Copy() also returns the number of bytes written (byte count);
	nbytes, err := io.Copy(ioutil.Discard, response.Body)
	response.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("error while reading %s: %v\n", url, err)
		return
	}

	secs := time.Since(start).Seconds()

	// as each result arrives, fetch() sends a summary line on the 'ch' channel;
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)

}
