// dup1 prints the text of each line that appears more than once in the standard input, preceded by its count;

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	counts := make(map[string]int)

	// input, of type bufio.Scanner, reads from the program's standard input;
	input := bufio.NewScanner(os.Stdin)

	// each call to input.Scan reads the next line and removes the newline character from the end;
	for input.Scan() {
		// input.Text() returns the line we just scanned from os.Stdin
		line := input.Text()
		// counts[line] represents the actual integer stored in the map for a specific string;
		// we add each line coming in from STDIN to the map and increment by 1;
		counts[line] = counts[line] + 1
		// the 2 lines above could also be written as: counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}
