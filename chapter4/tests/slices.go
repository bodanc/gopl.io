package main

import "fmt"

func slices() {

	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December"}

	fmt.Println(months[0], "<< empty string because of indexing starting at 1") // empty string

	q2 := months[4:7]
	northernSummer := months[6:9]

	fmt.Println(q2)
	fmt.Println(northernSummer)

	for _, s := range northernSummer {
		for _, q := range q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}

	// thee length of a slice represents the number of elements contained within;
	// the capacity is the number of elements between the start of the slice and the end of the underlying array;
	fmt.Println(len(q2))             // 3
	fmt.Println(len(northernSummer)) // 3
	fmt.Println(cap(q2))             // 9
	fmt.Println(cap(northernSummer)) // 7

	// slicing beyond the capacity of a slice causes a panic;
	//fmt.Println(northernSummer[:16]) // panic: runtime error: slice bounds out of range [:16] with capacity 7

	// slicing beyond the length of a slice is possible, as long as it falls within capacity;
	endlessSummer := northernSummer[:5]
	fmt.Println(endlessSummer) // [June July August September October]

	// a nil slice has a length and capacity of 0
	var s []int
	fmt.Println(len(s), cap(s), s == nil) // 0 0 true
	s = nil
	fmt.Println(len(s), cap(s), s == nil) // 0 0 true

	// as with any type that can have nil values, the nil value of a particular slice type can be written
	// using a conversion expression such as []int(nil);
	s = []int(nil)
	fmt.Println(len(s), cap(s), s == nil) // 0 0 true

	// there are non-nil slices of length and capacity zero, such as []int{}
	s = []int{}
	fmt.Println(len(s), cap(s), s == nil) // 0 0 false
	// make([]T, cap)[:len]
	sNew := make([]int, 3)[3:]
	fmt.Println(len(sNew), cap(sNew), sNew == nil) // 0 0 false

	// the built-in append function appends items to slices;
	var runes []rune
	// the loop uses append() to build the slice of 16 runes encoded by the string literal;
	for _, r := range "bogdan, eleonora" {
		runes = append(runes, r)
	}
	// the %q verb prints a safely escaped single-quoted character literal;
	fmt.Printf("%q\n", runes)

}
