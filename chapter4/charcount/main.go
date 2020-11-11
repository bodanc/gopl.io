// charcount counts the occurrences of each distinct Unicode code point in its input;
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {

	counts := make(map[rune]int)       // the count of Unicode characters
	var utfLength [utf8.UTFMax + 1]int // the count of UTF8 encodings lengths
	invalid := 0                       // the count of invalid UTF8 characters

	in := bufio.NewReader(os.Stdin)

	for {
		// the Reader.ReadRune() method performs UTF8 decoding and returns a rune, the number of bytes and an error
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		// if the rune represents an invalid Unicode character
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		counts[r]++
		utfLength[n]++

	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n) // %q: a single-quoted character literal
	}

	fmt.Print("\nlen\tcount\n")
	for i, n := range utfLength {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF8 characters\n", invalid)
	}

}
