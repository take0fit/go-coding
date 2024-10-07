package main

import (
	"fmt"
)

// 問題: 整数のリストから重複する要素を削除し、一意の要素のみを持つ新しいリストを返す関数を作成してください。
func removeDuplicates(arr []int) []int {
	uniqueSet := make(map[int]struct{})
	var result []int
	for _, num := range arr {
		if _, exists := uniqueSet[num]; !exists {
			uniqueSet[num] = struct{}{}
			result = append(result, num)
		}
	}
	return result
}

func main() {
	arr := []int{1, 2, 2, 3, 4, 4, 5}
	fmt.Println(removeDuplicates(arr)) // [1 2 3 4 5]
}
