package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	/*	result1 := 42.6666

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

		math.Sin(1)*/

	response1, err := errorHandlingStrategy1()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(response1)

	_, err = errorHandlingStrategy2()
	if err != nil {
		log.Println(err)
	}

	file, err := os.Open("file.txt")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(file)

/*	err = wait.WaitForServer("https://nosuchthing.io")
	fmt.Println(err)*/

/*	// 3rd error handling strategy
	if err := wait.WaitForServer("https://nosuchthing.io"); err != nil {
		fmt.Fprintf(os.Stderr, "the url is down: %v\n", err)
		os.Exit(1)
	}*/

/*	if err := wait.WaitForServer("https://nosuchthing.io"); err != nil {
		// as with all the log functions, log.Fatalf prefixes the time and date to the error message
		log.Fatalf("the url is down: %v\n", err)
	}*/

	dir, err := ioutil.TempDir("", "temp")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create temp dir: %v", err)
	}
	fmt.Println(dir)
	// the call to os.RemoveAll may fail, but we can safely ignore it because the o/s periodically cleans out
	// the temp dir;
	os.RemoveAll(dir)

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
