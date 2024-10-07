package main

import (
	"fmt"
	"strings"
)

func countWords(s string) int {
	words := strings.Fields(s)
	return len(words)
}

func main() {
	str := "Hello world this is a test"
	fmt.Println("単語数は", countWords(str), "です")
}
