package main

import "fmt"

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if backtrack(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func backtrack(board [][]byte, word string, i, j, index int) bool {
	if index == len(word) {
		return true
	}

	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] != word[index] {
		return false
	}

	temp := board[i][j]
	board[i][j] = ' '
	found := backtrack(board, word, i+1, j, index+1) ||
		backtrack(board, word, i-1, j, index+1) ||
		backtrack(board, word, i, j+1, index+1) ||
		backtrack(board, word, i, j-1, index+1)
	board[i][j] = temp
	return found

}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Println("単語が存在するか:", exist(board, word))
}
