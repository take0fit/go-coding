package main

import (
	"fmt"
)

// 問題: 与えられた整数の階乗を計算する関数を作成してください。
func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	num := 6
	fmt.Println(factorial(num))
}
