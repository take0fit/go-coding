package main

import (
	"fmt"
)

func countOccurrences(s string, target rune) int {
	count := 0
	for _, char := range s {
		if char == target {
			count++
		}
	}
	return count
}

func main() {
	str := "hello world"
	char := 'l'
	fmt.Printf("文字 '%c' の出現回数は %d 回です\n", char, countOccurrences(str, char))
}
