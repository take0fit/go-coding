package main

import (
	"fmt"
)

func removeAdjacentDuplicates(s string) string {
	if len(s) == 0 {
		return s
	}

	result := []rune{rune(s[0])}
	for _, char := range s[1:] {
		if char != result[len(result)-1] {
			result = append(result, char)
		}
	}
	return string(result)
}

func main() {
	str := "aabbccddeeff"
	fmt.Println("重複削除後の文字列:", removeAdjacentDuplicates(str))
}
