package main

import (
	"crypto/sha256"
	"fmt"
)

func arrays() {

	var a [3]int

	// by default, the elements of a new array are initially set to the zero value of the element type;
	fmt.Printf("%T %[1]v\n", a)
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	// print the index and element of the array;
	for index, value := range a {
		fmt.Printf("index: %d value: %d\n", index, value)
	}

	// we can use an array literal to initialize an array with a list of values
	var r = [3]int{1, 2}
	fmt.Println("last element of array 'r' of length 3:", r[len(r)-1])

	// in an array literal, an ellipsis can be used to determine the length of the array based on the number of
	// elements it is initialized with;
	q := [...]int{1, 2, 3}
	fmt.Printf("%d\n", len(q)) // 3
	fmt.Printf("%T\n", q)      // [3]int

	// it is possible to create an array by specifying a list of index and value pairs;
	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
	fmt.Println(USD, symbol[USD]) // 0 $
	fmt.Println(EUR, symbol[EUR]) // 1 €
	fmt.Println(GBP, symbol[GBP]) // 2 £
	fmt.Println(RMB, symbol[RMB]) // 3 ¥

	for i, v := range symbol {
		fmt.Println(i, v)
	}

	myArray := [10]string{0: "a", 1: "b", 2: "c", 9: "j"}
	for i, v := range myArray {
		fmt.Printf("i: %d v: %v", i, v)
	}
	fmt.Println()

	// we can define an array rNew with 10 elements (0 - 9) = to zero except for the last, which has a value of -1;
	rNew := [...]int{9: -1}
	fmt.Println(rNew)

	// if an array's element type is comparable, then the array type is comparable too;
	e := [2]int{1, 2}
	f := [2]int{1, 3}
	g := [...]int{1, 2}
	fmt.Println(e == f, e == g, f == g) // false true false

	// Sum256() produces the SHA256 cryptographic hash or digest of a message stored in an arbitrary byte slice;
	c1 := sha256.Sum256([]byte("hello"))
	c2 := sha256.Sum256([]byte("Hello"))
	c3 := sha256.Sum256([]byte("x"))
	c4 := sha256.Sum256([]byte("y"))

	// %x prints all the elements of an array or byte slice in hexadecimal form;
	// %t prints the boolean value of an expression;
	fmt.Printf("%x\n%x\n%t\n%T\n%T\n", c1, c2, c1 == c2, c1, c2)

	// unlike other prog. lang. go does not implicitly pass arrays by reference as arguments in the function call;
	zeroOut(&c1)
	fmt.Printf("%x\n", c1)
	zeroOut(&c2)
	fmt.Printf("%x\n", c2)

	zeroOutAlternative(&c3)
	fmt.Printf("%x\n", c3)
	zeroOutAlternative(&c4)
	fmt.Printf("%x\n", c4)

}

// zeroOut() zeroes out the content of a [32]byte array;
func zeroOut(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

// zeroOutAlternative() provides an alternative to zeroing out the content of the [32]byte array passed as an arg.
func zeroOutAlternative(ptr *[32]byte) {
	*ptr = [32]byte{}
}
