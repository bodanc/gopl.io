package main

import (
	"fmt"
)

func main() {

	a := [...]int{0, 1, 2, 3, 4}
	fmt.Println(a)

	// since a slice contains a pointer to an element of an array, passing a slice to a function permits the function
	// to modify the underlying array elements;
	reverse(a[:])
	fmt.Println(a) // [4 3 2 1 0]

	s := []int{0, 1, 2, 3, 4, 5}
	// rotate s left by two positions;
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s)

/*	input := bufio.NewScanner(os.Stdin)
outer:
	for input.Scan() {
		var numbers []int
		for _, s := range strings.Fields(input.Text()) {
			x, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				fmt.Println(err)
				continue outer
			}
			numbers = append(numbers, int(x))
		}
		reverse(numbers)
		fmt.Printf("%v\n", numbers)
	}*/

	s1 := []string{"a", "b", "c"}
	s2 := []string{"a", "b", "c", "d"}

	fmt.Println(equal(s1, s2)) // false

}

// reverse() reverses the elements of an []int slice in place;
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// equal() compares two slices of strings;
func equal(x, y []string) bool {

	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}

	return true

}
