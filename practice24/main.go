package main

import (
	"fmt"
)

func transpose(matrix [][]int) [][]int {
	rows := len(matrix)
	cols := len(matrix[0])
	transposed := make([][]int, cols)
	for i := range transposed {
		transposed[i] = make([]int, rows)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposed[j][i] = matrix[i][j]
		}
	}
	return transposed
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	trans := transpose(matrix)
	fmt.Println("転置行列:")
	for _, row := range trans {
		fmt.Println(row)
	}
}
