package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println(strings.Join(os.Args[:], " "))

	for i, v := range os.Args[:] {
		fmt.Println(i, v)
	}

}
