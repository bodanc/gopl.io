package main

import (
	"flag"
	"fmt"
	"strings"
)

// the optional -n flag causes echo to omit the trailing newline that would normally be printed;
var n = flag.Bool("n", false, "omit the trailing newline")

// the optional -s flag causes echo to separate the output arguments by the contents of the string sep
// instead of the default single space;
var sep = flag.String("s", " ", "separator")

func main() {

	flag.Parse()

	fmt.Printf("%T\n", n)
	fmt.Printf("%T\n", sep)

	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}

}
