package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {

	value := shaBitDiff([]byte("a"), []byte("A"))
	fmt.Printf("%T %[1]v\n", value)
}

func popCount(b byte) int {
	count := 0
	for ; b != 0; count++ {
		b &= b - 1
	}
	return count
}

func bitDiff(a, b []byte) int {
	count := 0
	for i := 0; i < len(a) || i < len(b); i++ {
		switch {
		case i >= len(a):
			count += popCount(b[i])
		case i >= len(b):
			count += popCount(a[i])
		default:
			count += popCount(a[i] ^ b[i])
		}
	}
	return count
}

// shaBitDiff() returns the number of bits that are different between the SHA256 hashes of 2 byte slices;
func shaBitDiff(a, b []byte) int {

	shaA := sha256.Sum256(a)
	shaB := sha256.Sum256(b)

	fmt.Printf("%x\n", shaA)
	fmt.Printf("%x\n", shaB)

	return bitDiff(shaA[:], shaB[:])

}
