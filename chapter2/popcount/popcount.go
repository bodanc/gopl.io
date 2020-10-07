// package initialization;
// pre-computing a table of values is a useful programming technique;

package main

import "fmt"

// pc[i] is the population count of i;
var pc [256]byte

// init() precomputes a table of results (pc) for each possible 8-bit value;
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount() returns the number of set bits, that is bits with a value of 1, in a unit64 value (the population count);
func PopCount(x uint64) int {

	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])

}

func main() {

	fmt.Println(PopCount(255))

}
