// an alternative approach to reading input: read the entire input into memory in one big chunk, split it into lines
// all at once, and only then process the input line by line;

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	count := make(map[string]int)

	for _, filename := range os.Args[1:] {

		// ioutil.ReadFile() reads the entire content of a named file and returns a []byte;
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		// strings.Split() is the opposite of strings.Join() and splits a string into a slice of substrings;
		for _, line := range strings.Split(string(data), "\n") {
			count[line]++
		}

	}

	for line, number := range count {
		if number > 1 {
			fmt.Printf("%d\t%s\n", number, line)
		}
	}

}
