package main

import (
	"fmt"
	"net/http"
)

func errorHandlingStrategy1() (*http.Response, error) {

	url := "https://mirror.centos.org/"
	// propagate the error so that a failure in a subroutine becomes a failure of the calling routine;
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return response, nil

}

func errorHandlingStrategy2() (*http.Response, error) {

	url := "http://mirror.centos.org/"
	response, err := http.Get(url)
	if err != nil {
		return nil, err // a nil []string
	}

	if response.StatusCode != http.StatusOK {
		response.Body.Close()
		// fmt.Errorf formats an error message using fmt.Sprintf and returns a new error value;
		return nil, fmt.Errorf("error retrieving %s: %s\n", url, response.Status)
	}

	return response, nil

}
