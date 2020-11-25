package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {

	//arrays()
	//slices()

	//maps()
	//readRune()

	//structs()
	//drawWheel()
	booksJSON()

}

func readRune() {

	// a holds the os.Stdin input data inside a []byte
	a := bufio.NewReader(os.Stdin)

	// Reader.ReadRune() returns the rune representation of the data inside the []byte
	for {
		r, s, err := a.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "bla(): %v\n", err)
			os.Exit(1)
		}

		if r == 10 {
			continue
		}

		fmt.Printf("rune: %v character: %[1]q size: %d\n", r, s)
	}

}
