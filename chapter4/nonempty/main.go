package main

import "fmt"

// nonEmpty() returns a slice which holds only the non-empty strings;
// note: the underlying array is modified during the function call;
func nonEmpty1(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonEmpty2(strings []string) []string {
	out := strings[:0] // zero-length slice of the original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func main() {

	data := []string{"a", "b", "c", "d", "", "", "g", "h"}

	// the subtle part is that the input slice and the output slice share the same underlying array;
	// this does avoid the need to allocate a new array, at the expense of overwriting part of the slice content;
	fmt.Printf("%q\n", nonEmpty1(data))
	fmt.Printf("%q\n", data)

	// to avoid the situation above, we can rewrite the values to the initial slice;
	data1New := []string{"a", "b", "c", "d", "", "", "g", "h"}
	data1New = nonEmpty1(data1New)
	fmt.Printf("%q\n", data1New)

	data2 := []string{"a", "b", "c", "d", "", "", "g", "h"}
	fmt.Printf("%q\n", nonEmpty2(data2))
	fmt.Printf("%q\n", data)

	data2New := []string{"a", "b", "c", "d", "", "", "g", "h"}
	data2New = nonEmpty2(data2New)
	fmt.Printf("%q\n", data2New)

}
