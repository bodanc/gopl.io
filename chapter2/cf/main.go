// cf converts its numeric argument to Celsius and Fahrenheit;

package main

import (
	"fmt"
	"github.com/bodanc/gopl.io/chapter2/tempconv"
	"os"
	"strconv"
)

func main() {

	for _, arg := range os.Args[1:] {

		temp, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		c := tempconv.Celsius(temp)
		f := tempconv.Fahrenheit(temp)

		fmt.Printf("%s = %s\n", c, tempconv.CToF(c))
		fmt.Printf("%s = %s\n", f, tempconv.FToC(f))

	}

}
