package main

import (
	"fmt"
)

func findMax(arr []int) (int, error) {
	if len(arr) == 0 {
		return 0, fmt.Errorf("配列が空です")
	}
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	return max, nil
}

func main() {
	arr := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}
	max, err := findMax(arr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("最大値は", max, "です")
}
