package main

import "fmt"

func appendInt(x []int, y int) []int {

	var z []int
	zlen := len(x) + 1

	// each call to appendInt() must check if the slice has sufficient capacity to hold the new elements in the
	// existing array;
	if zlen <= cap(x) {
		// if there is room to grow, we can extend the slice by defining a larger slice;
		z = x[:zlen]
	} else {
		// there is insufficient space; allocate a new array and grow by doubling for amortized linear complexity;
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}

	z[len(x)] = y
	return z

}

// with the modifications below, we can more closely match the behavior of the built-in append() function;
func appendSlice(x []int, y ...int) []int {

	var z []int
	zlen := len(x) + len(y)

	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}

	copy(z[len(x):], y)
	return z

}

func main() {

	/*	var x, y []int

		for i := 0; i < 10; i++ {
			y = appendInt(x, i)
			fmt.Printf("%d len=%d cap=%d %v\n", i, len(y), cap(y), y)
			x = y
		}*/

	var x = []int{0, 1, 2, 3}
	var y = []int{4, 5, 6}

	result := appendSlice(x, y...)
	fmt.Printf("len: %d cap: %d result: %v\n", len(result), cap(result), result)

}
