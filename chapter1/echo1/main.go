package main

import (
	"fmt"
	"os"
)

func main() {

	// these 2 variables are implicitly initialized to the zero value for their type;
	// this declaration implicitly initializes s and separator to empty strings;
	var s, separator string

	// i is a loop index variable, which we declare in the first part of the for loop with a short variable declaration;
	for i := 1; i < len(os.Args); i++ {
		// when applied to strings, the + operator concatenates the values;
		// this is an assignment statement using the assignment operator '+=' as in s = s + sep + os.Args[i];
		s += separator + os.Args[i]
		separator = " "
	}

	fmt.Println(s)

}
