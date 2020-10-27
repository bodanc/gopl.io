package main

import (
	"fmt"
	"unicode/utf8"
)

func strings() {

	// a string is an immutable sequence of bytes;
	s := "hello, world"

	// the builtin len() function returns the number of bytes, not runes, in a string;
	fmt.Println(len(s))            // 12 bytes
	fmt.Println(s[0], s[7], s[11]) // 104 119 100 (h, w and d) <- unicode code points (runes)

	fmt.Println(s[0:5]) // hello
	fmt.Println(s[7:])  // world

	fmt.Println("goodbye" + s[5:]) // goodbye, world

	string1 := "\xe4\xb8\x96\xe7\x95\x8c"
	string2 := "\u4e16\u754c"         // numeric code point value, in 16-bit form
	string3 := "\U00004e16\U0000754c" // numeric code point value, in 32-bit form

	fmt.Println(string1)
	fmt.Println(string2)
	fmt.Println(string3)
	fmt.Println(string1 == string2)
	fmt.Println(string2 == string3)
	fmt.Println(string1 == string3)

	fmt.Println(HasPrefix("bogdan", "b"))

	fmt.Println(Contains("page", "a"))

	s = "hello, 世界"
	fmt.Println(len(s))                    // 13 bytes
	fmt.Println(utf8.RuneCountInString(s)) // 9 runes

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("index: %d\tcharacter: %c\tsize: %d\n", i, r, size)
		i += size
	}

	// the range loop, when applied to a string, performs UTF-8 decoding implicitly;
	for i, r := range "hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	str := "プローグラム"
	fmt.Printf("% x\n", str) // e3 83 97 e3 83 ad e3 83 bc e3 82 b0 e3 83 a9 e3 83 a0
	// a rune conversion applied to an UTF-8 encoded string returns the sequence of Unicode code points
	// that the string encodes;
	r := []rune(str)
	fmt.Printf("%x\n", r) // [30d7 30ed 30fc 30b0 30e9 30e0]
	// if a slice of runes is converted to a string, it produces the concatenation of the UTF-8 encodings
	// of each rune;
	fmt.Println(string(r))

	strNew := "bogdan"
	fmt.Printf("% x\n", strNew) // 62 6f 67 64 61 6e (hex digits)
	r = []rune(strNew)
	fmt.Printf("%x\n", r)  // [62 6f 67 64 61 6e] (unicode code points)
	fmt.Println(string(r)) // bogdan

	// converting an integer value to a string interprets the integer as a rune value and yields the
	// UTF-8 representation of that rune;
	fmt.Println(string(65)) // A
	// if the rune is invalid, the replacement character is substituted;
	fmt.Println(string(1234567)) // �

}

func stringsAndByteSlices() {



}