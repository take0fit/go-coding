package main

import "fmt"

func longestValidParentheses(s string) int {
	maxLen := 0
	stack := []int{}
	stack = append(stack, -1)

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				currentLen := i - stack[len(stack)-1]
				if currentLen > maxLen {
					maxLen = currentLen
				}
			}
		}
	}

	return maxLen
}

func main() {
	s := "(()"
	fmt.Println("最長の有効な括弧の長さ:", longestValidParentheses(s))
}
