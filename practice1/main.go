package main

import (
	"fmt"
)

// 問題: 1から100までの数字を順に出力してください。
// ただし、数字が3の倍数のときは「Fizz」、5の倍数のときは「Buzz」、15の倍数のときは「FizzBuzz」と出力してください。
func main() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
