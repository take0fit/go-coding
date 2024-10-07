package main

import (
	"fmt"
)

// 問題: n番目のフィボナッチ数を計算する関数を作成してください。
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	n := 10
	fmt.Println(fibonacci(n))
}
