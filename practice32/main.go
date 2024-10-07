package main

import (
	"fmt"
	"strings"
)

func isSubstring(s1, s2 string) bool {
	return strings.Contains(s2, s1)
}

func main() {
	fmt.Println(isSubstring("test", "this is a test string")) // true
}
