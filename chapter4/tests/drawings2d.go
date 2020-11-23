// a small program to demonstrate struct embedding and anonymous fields

package main

import "fmt"

type Point struct {
	X, Y int
}

// a Circle has fields for the X and Y coordinates of its center, and a Radius
type Circle struct {
	Point  // anonymous field (named type or pointer to a named type); a Point is embedded within Circle
	Radius int
}

// a Wheel has all the features of a Circle, plus Spokes, the number of inscribed radial spokes
type Wheel struct {
	// an anonymous field need not be a struct type
	Circle // a Circle is embedded within Wheel
	Spokes int
}

func drawWheel() {

	// thanks to embedding, we can refer to the names at the leaves of the implicit tree
	var w Wheel
	w.X = 8      // equivalent to w.Circle.Point.X = 8
	w.Y = 8      // equivalent to w.Circle.Point.Y = 8
	w.Radius = 5 // equivalent to w.Circle.Radius = 5
	w.Spokes = 16

	w2 := Wheel{
		Circle: Circle{Point{
			X: 8,
			Y: 8,
		}, 5},
		Spokes: 16,
	}

	w3 := Wheel{}
	w3.Circle.Point.X = 1

	w = Wheel{Circle{Point{9, 9}, 6}, 17}

	w = Wheel{
		Circle: Circle{
			Point:  Point{10, 2},
			Radius: 4, // a comma is necessary
		},
		Spokes: 19, // a comma is necessary
	}

	fmt.Println(w)
	fmt.Println(w2)
	fmt.Println(w3)

	// the # adverb causes the %v verb to include the name of each struct field
	fmt.Printf("%#v\n", w)
	w.X = 42
	fmt.Printf("%#v\n", w)

}
