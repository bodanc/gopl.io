package main

import (
	"fmt"
)

func main() {

	// the expression new(T) creates an unnamed variable of type T, initializes it to the zero value of T, and returns
	// its address, which is a value of type *T;

	// p, of type *int, points to an unnamed int variable;
	p := new(int)
	fmt.Printf("%T %[1]v\n", p) // *int 0xc00001a080
	fmt.Println(*p)             // 0

	*p = 2
	fmt.Println(*p) // 2

	bla := new(string)
	fmt.Printf("%T %[1]v\n", bla) // *string 0xc000108040
	fmt.Println(*bla)             // ""

	*bla = "bogdan"
	fmt.Println(*bla) // "bogdan"

	myNumber := new(int)
	fmt.Printf("%T\n", myNumber)

	fmt.Println("gcd of 20 and 15 is:", gcd(20, 15))
	fmt.Println("Fibonacci for 20:", fib(20))

}

// gcd() returns the greatest common divisor of x and y;
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y // tuple assignment
	}
	return x
}

// fib() computes the n-th Fibonacci number iteratively;
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y // tuple assignment
	}
	return x
}
