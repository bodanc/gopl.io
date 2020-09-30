package main

import (
	"fmt"
	"os"
)

func main() {

	s, separator := "", ""

	// we choose to not declare an initialization statement or a post statement for this loop;
	// since range returns both an index and the value of an element at that index, we choose to ignore it by using the
	// blank identifier;
	for _, arg := range os.Args[1:] {
		s += separator + arg
		separator = " "
	}

	fmt.Println(s)

	sPrime, separatorPrime := "", ""

	for _, argPrime := range os.Args[0:] {
		sPrime += separatorPrime + argPrime
		separatorPrime = " "
	}

	fmt.Println(sPrime)

}
