package main

import (
	"fmt"
)

// 問題: 整数のリストから最小値を見つける関数を作成してください。
func findMin(arr []int) (int, error) {
	if len(arr) == 0 {
		return 0, fmt.Errorf("配列が空です")
	}
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min, nil
}

func main() {
	arr := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}
	min, err := findMin(arr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("最小値は", min, "です")
}
