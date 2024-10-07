package main

import (
	"fmt"
)

func reverseList(arr []int) []int {
	n := len(arr)
	reversed := make([]int, n)
	for i, val := range arr {
		reversed[n-i-1] = val
	}
	return reversed
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("逆順リスト:", reverseList(arr))
}
