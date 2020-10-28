package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := basename("/log/file.a")
	fmt.Println(str1)

	str2 := basename("/home/bogdan/file.a.b")
	fmt.Println(str2)

	str3 := basename("//a//b//c.txt")
	fmt.Println(str3)

	str4 := basename("a.b.c.d.e.f.g")
	fmt.Println(str4)

	str5 := basename("/a/b/c/file.x.y.z")
	fmt.Println(str5)

}

func basename(s string) string {

	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]

	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}

	return s

}
