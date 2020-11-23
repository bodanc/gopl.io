package main

import "fmt"

// a named struct 'S' can declare a field of type '*S' which lets us create recursive data structures
type tree struct {
	value       int
	left, right *tree
}

// Sort
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues() appends the elements of t to values in order and returns the resulting slice
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}

	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}

	return t
}

func main() {

	v := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	Sort(v)
	fmt.Println(v) // [0 1 2 3 4 5 6 7 8 9]

}
