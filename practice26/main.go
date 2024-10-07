package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		words[i] = reverseString(word)
	}
	return strings.Join(words, " ")
}

func reverseString(s string) string {
	runes := []rune(s)
	length := len(runes)
	for i := 0; i < length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}
	return string(runes)
}

func main() {
	str := "Hello world"
	fmt.Println("各単語を逆転:", reverseWords(str))
}
