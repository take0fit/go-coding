package main

import (
	"fmt"
)

func difference(arr1, arr2 []int) []int {
	m := make(map[int]bool)
	for _, num := range arr2 {
		m[num] = true
	}

	var result []int
	for _, num := range arr1 {
		if !m[num] {
			result = append(result, num)
		}
	}
	return result
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{4, 5, 6, 7, 8}
	fmt.Println("差分リスト:", difference(arr1, arr2))
}
