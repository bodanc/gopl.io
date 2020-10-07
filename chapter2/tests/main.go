package main

import (
	"fmt"
	"github.com/bodanc/gopl.io/chapter2/tempconv"
	"log"
	"os"
)

var A int
var B int
var C int
var cWD string

// any .go file may contain any number of init() functions;
func init() {
	A = 1
	fmt.Println("outer cWD:", cWD)
}

func init() {
	B = 2
}

func init() {
	C = 3
}

func init() {
	cWD, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v\n", err)
	}
	fmt.Println("inner cWD:", cWD)
}

func main() {

	d := A + B + C
	fmt.Println(d)

	// this example illustrates scope rules, not good style!
	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a' // not equivalent to unicode.ToUpper()
			fmt.Printf("%c", x)
		}
	}
	fmt.Println()



	// the expression new(T) creates an unnamed variable of type T, initializes it to the zero value of T, and returns
	// its address, which is a value of type *T;

	// p, of type *int, points to an unnamed int variable;
	p := new(int)
	fmt.Printf("%T %[1]v\n", p) // *int 0xc00001a080
	fmt.Println(*p)

	*p = 2
	fmt.Println(*p)

	bla := new(string)
	fmt.Printf("%T %[1]v\n", bla) // *string 0xc000108040
	fmt.Println(*bla)

	*bla = "watermelon"
	fmt.Println(*bla) // "watermelon"

	myNumber := new(int)
	fmt.Printf("%T\n", myNumber)

	fmt.Println(gcd(20, 15))
	fmt.Println(fib(10))

	myMap := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7}
	value1, ok := myMap["a"]
	value2, ok := myMap["b"]
	value3, ok := myMap["c"]
	fmt.Println(value1, ok) // 1 true
	fmt.Println(value2, ok) // 2 true
	fmt.Println(value3, ok) // 3 true

	fmt.Printf("brrrr! it's %v", tempconv.AbsoluteZeroC)

	fmt.Println(tempconv.CToF(tempconv.BoilingC))

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
