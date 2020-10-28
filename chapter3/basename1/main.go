package main

import "fmt"

func main() {

	str1 := basename("/var/log/file.a")
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

// the basename() function removes any prefix of 's' that looks like a file system path with components separated
// by slashes, as well as any suffix that looks like a file type
func basename(s string) string {

	// discard the last '/' and everything before it
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// preserve everything before the last '.'
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}
