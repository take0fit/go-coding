package main

import (
	"fmt"
)

func keys(m map[int]int) []int {
	var keys []int
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

func main() {
	m := map[int]int{1: 10, 2: 20, 3: 30}
	fmt.Println("キーのリスト:", keys(m))
}
