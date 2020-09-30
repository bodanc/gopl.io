// dup2 prints the count and the text of lines that appear more than once in the input;
// it can read from stdin or from a list of named files;

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// a map is a REFERENCE to the data structure created by make;
	count := make(map[string]int)

	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, count)
	} else {
		for _, arg := range files {
			// os.Open() returns 2 values: an open *os.File which will later be used by bufio.Scanner and an err;
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				// the continue statement goes to the next iteration of the enclosing for loop;
				continue
			}
			// when a map is passed to a function, the function receives a copy of the REFERENCE;
			countLines(f, count)
			f.Close()
		}
	}

	for line, n := range count {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}

func countLines(f *os.File, count map[string]int) {

	input := bufio.NewScanner(f)

	for input.Scan() {
		// count[input.Text()]++
		line := input.Text()
		count[line] += 1
	}

}
