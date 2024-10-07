package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	maxLen, left := 0, 0

	for right := 0; right < len(s); right++ {
		if idx, ok := m[s[right]]; ok && idx >= left {
			left = idx + 1
		}
		m[s[right]] = right
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}

func main() {
	s := "abcabcbb"
	fmt.Println("最長部分文字列の長さ:", lengthOfLongestSubstring(s))
}
