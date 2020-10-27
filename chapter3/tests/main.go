package main

import (
	"fmt"
	"math"
)

const (
	Avogadro = 6.02214129e23
	Planck   = 6.62606957e-34
)

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func itob(i int) bool {
	return i != 0
}

// test if the input string contains another string as a prefix without decoding the UTF-8 code points
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

// test if the input string contains another string as a suffix, without decoding the UTF-8 code points
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

// test if the input string contains a substring, without decoding the UTF-8 code points
func Contains(s, substring string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substring) {
			return true
		}
	}
	return false
}

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

	// the built-in function complex() creates a complex number from its real and imaginary components
	var xCmpl complex128 = complex(1, 2) // (1+2i)
	var yCmpl complex128 = complex(3, 4) // (3+4i)
	fmt.Println(xCmpl * yCmpl)           // (-5+10i)

	// the built-in functions real() and imag() extract the real and imaginary components of complex numbers
	fmt.Println(real(xCmpl)) // 1
	fmt.Println(imag(xCmpl)) // 2

	fmt.Println(real(yCmpl)) // 3
	fmt.Println(imag(yCmpl)) // 4

	fmt.Println(real(xCmpl * yCmpl)) // -5
	fmt.Println(imag(xCmpl * yCmpl)) // 10

	strings()
	stringsAndByteSlices()

}
