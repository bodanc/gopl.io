package main

import (
	"crypto/sha256"
	sha5122 "crypto/sha512"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	sha384 string
	sha512 string
)

func main() {

	if len(os.Args[1:]) == 0 {
		log.Fatal("too few arguments")
	}

	resultA := [32]byte{}
	resultB := [48]byte{}
	resultC := [64]byte{}

	flag.StringVar(&sha384, "sha384", "", "calculate the sha384 hash")
	flag.StringVar(&sha512, "sha512", "", "calculate the sha512 hash")

	flag.Parse()

	if sha384 != "" {
		resultB = sha5122.Sum384([]byte(sha384))
		fmt.Printf("%x\n", resultB)
	} else if sha512 != "" {
		resultC = sha5122.Sum512([]byte(sha512))
		fmt.Printf("%x\n", resultC)
	} else {
		resultA = sha256.Sum256([]byte(os.Args[1]))
		fmt.Printf("%x\n", resultA)
	}

}
