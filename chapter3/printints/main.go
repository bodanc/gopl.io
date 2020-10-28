package main

import (
	"bytes"
	"fmt"
)

func intsToString(values []int) string {

	// a bytes.Buffer variable requires no initialization because its zero value is usable
	var buf bytes.Buffer

	// bytes.Buffer.WriteByte() is fine for ASCII characters, but not for arbitrary runes
	buf.WriteByte('[')

	// when appending the UTF-8 encoding of an arbitrary rune to a bytes.Buffer, it's best to use the WriteRune
	// method;
	buf.WriteRune(97)

	for i, v := range values {
		if i == 0 {
			buf.WriteString(" ")
		}
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}

	buf.WriteByte(']')

	return buf.String()

}

func main() {

	fmt.Println(intsToString([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))

}
