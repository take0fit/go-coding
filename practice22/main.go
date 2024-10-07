package main

import (
	"fmt"
)

func commonCharacters(s1, s2 string) []rune {
	charMap := make(map[rune]bool)
	for _, char := range s1 {
		charMap[char] = true
	}
	var common []rune
	for _, char := range s2 {
		if charMap[char] {
			common = append(common, char)
			charMap[char] = false // 重複を避ける
		}
	}
	return common
}

func main() {
	s1 := "apple"
	s2 := "grape"
	fmt.Printf("共通の文字: %c\n", commonCharacters(s1, s2))
}
