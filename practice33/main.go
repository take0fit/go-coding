package main

import (
	"fmt"
)

func gcdRecursive(a, b int) int {
	if b == 0 {
		return a
	}
	return gcdRecursive(b, a%b)
}

func main() {
	fmt.Println("最大公約数は", gcdRecursive(56, 98), "です")
}
