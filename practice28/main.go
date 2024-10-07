package main

import (
	"fmt"
)

func sumMapValues(m map[int]int) int {
	sum := 0
	for _, value := range m {
		sum += value
	}
	return sum
}

func main() {
	m := map[int]int{1: 10, 2: 20, 3: 30}
	fmt.Println("値の合計は", sumMapValues(m), "です")
}
