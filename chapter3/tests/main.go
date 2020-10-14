package main

import (
	"fmt"
	"math"
)

const (
	Avogadro = 6.02214129e23
	Planck   = 6.62606957e-34
)

func main() {

	var u uint8 = 255
	fmt.Println(u, u+1, u*u) // 255 0 1

	var i int8 = 127
	fmt.Println(i, i+1, i*i) // 127, -128, 1

	f := 1e100
	fmt.Println(int(f)) // the result is implementation-dependent

	// when printing numbers using the fmt package, we can control the radix and format with the following verbs:
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o) // 438 666 0666
	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x) // 3735928559 deadbeef 0xdeadbeef 0XDEADBEEF

	fmt.Printf("%T %[1]g\n", Avogadro) // float64 6.02214129e+23
	fmt.Printf("%T %[1]g\n", Planck)   // float64 6.62606957e-34

	// for tables of data, the %e (exponent) and %f (no exponent) forms may be the most appropriate;
	for x := 0; x < 8; x++ {
		// 4 digits of precision, aligned in a ten-character field;
		fmt.Printf("x = %d e^x = %10.4e e^x = %10.4f\n", x, math.Exp(float64(x)), math.Exp(float64(x)))
	}

	// mathematically dubious operations;
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // 0 -0 +Inf -Inf NaN

	var dubious1 float64 = math.Sqrt(-1)
	fmt.Println("math.Sqrt(-1) = ", dubious1)


	maxF32 := math.MaxFloat32
	fmt.Println("maxFloat32:", maxF32)

	maxF64 := math.MaxFloat64
	fmt.Println("maxFloat64:", maxF64)

}
