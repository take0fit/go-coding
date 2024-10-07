package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	// 文字列を正規化（小文字化と空白除去）
	var filtered []rune
	for _, r := range strings.ToLower(s) {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			filtered = append(filtered, r)
		}
	}
	// 左右から文字を比較
	for i, j := 0, len(filtered)-1; i < j; i, j = i+1, j-1 {
		if filtered[i] != filtered[j] {
			return false
		}
	}
	return true
}

func main() {
	str := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(str)) // true
}
