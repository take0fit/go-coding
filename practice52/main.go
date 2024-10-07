package main

import (
	"fmt"
)

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}
	if !wordSet[endWord] {
		return 0
	}

	queue := []string{beginWord}
	level := 1

	for len(queue) > 0 {
		nextQueue := []string{}
		for _, word := range queue {
			wordBytes := []byte(word)
			for i := 0; i < len(wordBytes); i++ {
				originalChar := wordBytes[i]
				for c := 'a'; c <= 'z'; c++ {
					wordBytes[i] = byte(c)
					newWord := string(wordBytes)
					if newWord == endWord {
						return level + 1
					}
					if wordSet[newWord] {
						nextQueue = append(nextQueue, newWord)
						delete(wordSet, newWord)
					}
				}
				wordBytes[i] = originalChar
			}
		}
		queue = nextQueue
		level++
	}

	return 0
}

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	fmt.Println("最小ステップ数:", ladderLength(beginWord, endWord, wordList))
}
