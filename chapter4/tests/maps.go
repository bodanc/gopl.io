package main

import (
	"fmt"
	"sort"
)

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for key, xValue := range x { // alpha 1
		if yValue, ok := y[key]; !ok || yValue != xValue {
			return false
		}
	}
	return true
}

func equalWrong(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for key, xValue := range x { // A, 0
		if xValue != y[key] { // 0 != y[A]
			return false
		}
	}
	return true
}

func maps() {

	fmt.Println("maps")

	map1 := make(map[string]int)
	map2 := map[string]int{}

	fmt.Printf("map1: %v map2: %v\n", map1, map2)
	fmt.Printf("map1: %T map2: %T\n", map1, map2)
	fmt.Printf("len(map1) == len(map2): %t\n", len(map1) == len(map2))

	fmt.Println("a non-existent key:", map1["aNonexistentKey"])

	ages := map[string]int{
		"a": 10,
		"b": 11,
		"c": 12,
	}

	ages["d"] = 13
	ages["e"] = 14
	ages["f"] = 15
	ages["g"] = 16
	ages["h"] = 17
	ages["i"] = 18
	ages["j"] = 19
	ages["k"] = 20

	fmt.Println("ages[\"a\"]:", ages["a"])

	ages["a"] = 100
	fmt.Println("ages[\"a\"]:", ages["a"])

	ages["a"]++
	fmt.Println("ages[\"a\"]:", ages["a"])

	fmt.Println(ages)
	delete(ages, "k")
	fmt.Println(ages)
	delete(ages, "x") // attempt at deleting a non-existent key

	// the order of map iteration is unspecified (random), an intentional design choice
	for key, value := range ages {
		value += 100
		fmt.Println(key, value)
	}

	names := make([]string, 0, len(ages))
	for name := range ages { // implicitly ignore the value
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names { // ignore the index, assign the slice element to the name variable
		fmt.Printf("%s %d ", name, ages[name])
	}
	fmt.Println()

	// assigning values to a nil map will cause a panic
	/*	var stars map[string]int
		stars["pistol star"] = 1 // panic: assignment to entry in nil map
	*/

	// a way to determine if a certain element is present in a map or not
	age, ok := ages["x"]
	fmt.Println(age, ok) // 0 false

	if age, ok := ages["x"]; !ok {
		fmt.Println(age)
		fmt.Println("the element is not preset")
	}

	x := make(map[string]int)
	y := make(map[string]int)

	x["alpha"] = 1
	y["alpha"] = 1
	x["beta"] = 2
	y["beta"] = 2
	x["gamma"] = 3
	y["gamma"] = 3
	x["delta"] = 4
	y["delta"] = 4
	x["epsilon"] = 5
	y["epsilon"] = 5
	x["zeta"] = 6
	y["zeta"] = 60

	fmt.Println(equal(x, y))

	// xValue != y[key]
	fmt.Println(equalWrong(map[string]int{"A": 0}, map[string]int{"B": 1}))

}
