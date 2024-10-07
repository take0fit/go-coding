package main

import (
	"fmt"
)

func power(base, exponent int) int {
	result := 1
	for exponent > 0 {
		if exponent%2 == 1 {
			result *= base
		}
		base *= base
		exponent /= 2
	}
	return result
}

func main() {
	fmt.Println("2の10乗は", power(2, 10), "です")
}
