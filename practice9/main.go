package main

import (
	"fmt"
)

func isSumEven(arr []int) bool {
	total := 0
	for _, num := range arr {
		total += num
	}
	return total%2 == 0
}

func main() {
	arr := []int{1, 2, 3, 4}
	if isSumEven(arr) {
		fmt.Println("合計は偶数です")
	} else {
		fmt.Println("合計は奇数です")
	}
}
