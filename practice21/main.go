package main

import (
	"fmt"
)

// 問題: 2つの整数リストの共通部分を返す関数を作成してください。
func intersection(arr1, arr2 []int) []int {
	m := make(map[int]bool)
	for _, num := range arr1 {
		m[num] = true
	}

	var result []int
	for _, num := range arr2 {
		if m[num] {
			result = append(result, num)
			m[num] = false // 重複を避ける
		}
	}
	return result
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{4, 5, 6, 7, 8}
	fmt.Println("共通部分:", intersection(arr1, arr2))
}
