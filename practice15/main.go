package main

import (
	"fmt"
	"math"
)

// 問題: 与えられた整数が素数であるかどうかを判定する関数を作成してください。
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	// 2から平方根までの数字で割り切れるかをチェック
	sqrt := int(math.Sqrt(float64(n)))
	fmt.Println(sqrt)
	for i := 2; i <= sqrt; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	number := 37
	fmt.Println(isPrime(number))
}
