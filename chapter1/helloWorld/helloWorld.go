// package main is special; it defines a standalone executable program, not a library;
package main

// we must tell the compiler which packages are needed by this source file;
// this program uses one function from one other package, called "fmt";
import "fmt"

// within the main package, the main() function is also special; the execution of the program starts here;
// whatever main() does, is what the program does;
func main() {

	fmt.Println("hello")

}
