// boiling prints the boiling point of water;

package main

import "fmt"

// the constant boilingF is a package-level declaration;
const boilingF = 212.0

func main() {

	// variables f and c are local to main();
	var f = boilingF
	var c = (f - 32) * 5 / 9

	fmt.Printf("boiling point of water: %g F or %g C\n", f, c)
	// output:
	// boiling point of water: 212 F or 100 C

}
