package main

import (
	"fmt"
	"sort"
)

// 問題: 整数のリストの中央値を求める関数を作成してください。リストはソートされていない場合があります。
func median(arr []int) (float64, error) {
	if len(arr) == 0 {
		return 0, fmt.Errorf("配列が空です")
	}
	sort.Ints(arr)
	n := len(arr)
	if n%2 == 0 {
		return float64(arr[n/2-1]+arr[n/2]) / 2.0, nil
	}
	return float64(arr[n/2]), nil
}

func main() {
	arr := []int{3, 5, 1, 4, 2}
	med, err := median(arr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("中央値は", med, "です")

	arr = []int{3, 5, 1, 4, 2, 6}
	med, err = median(arr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("中央値は", med, "です")
}
