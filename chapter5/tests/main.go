package main

import (
	"fmt"
	"golang.org/x/net/html"
	"math"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	result1 := 42.6666

	result1 = add(10, 999)
	fmt.Printf("%f\n", result1)

	result2 := subtract(16, 4)
	fmt.Printf("%d\n", result2)

	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(addSlice(s))

	bogdan := Person{
		Name: "franz",
		Age:  99,
	}

	changeName(bogdan)
	fmt.Println(bogdan)

	math.Sin(1)

}

func add(a, b float64) (c float64) {
	c = a + b
	return
}

func subtract(a, b int) (c int) {
	c = a - b
	return
}

func addSlice(a []int) int {
	var result int
	for _, v := range a {
		result += v
	}
	return result
}

func changeName(p Person) {
	p.Name += "kafka"
}

// in a function with named results, the operands of a return statement may be omitted;
// this is called a bare return;
func CountWordsAndImages(url string) (words, images int, err error) { // named results
	resp, err := http.Get(url)
	if err != nil {
		return // bare return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return // bare return
	}

	words, images = someOtherFunc(doc)

	// a bare return is shorthand for returning each of the named result variables in order;
	return // bare return

}

func someOtherFunc(n *html.Node) (int, int) {
	return 1, 2
}
