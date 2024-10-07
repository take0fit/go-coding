package main

import (
	"fmt"
)

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	// tの文字カウントを作成
	tCount := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tCount[t[i]]++
	}

	required := len(tCount)
	left, right := 0, 0
	windowCount := make(map[byte]int)
	formed := 0
	ans := []int{-1, 0, 0}

	for right < len(s) {
		char := s[right]
		windowCount[char]++

		if tCount[char] > 0 && windowCount[char] == tCount[char] {
			formed++
		}

		for left <= right && formed == required {
			char = s[left]
			if ans[0] == -1 || right-left+1 < ans[0] {
				ans[0] = right - left + 1
				ans[1] = left
				ans[2] = right
			}

			windowCount[char]--
			if tCount[char] > 0 && windowCount[char] < tCount[char] {
				formed--
			}
			left++
		}
		right++
	}

	if ans[0] == -1 {
		return ""
	}
	return s[ans[1] : ans[2]+1]
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println("最小ウィンドウ:", minWindow(s, t))
}
