package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var backtrack func(target, start int, comb []int)
	backtrack = func(target, start int, comb []int) {
		if target == 0 {
			combCopy := make([]int, len(comb))
			copy(combCopy, comb)
			result = append(result, combCopy)
			return
		}
		for i := start; i < len(candidates); i++ {
			if candidates[i] <= target {
				comb = append(comb, candidates[i])
				backtrack(target-candidates[i], i, comb)
				comb = comb[:len(comb)-1]
			}
		}
	}
	backtrack(target, 0, []int{})
	return result
}

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	fmt.Println("組み合わせの和:", combinationSum(candidates, target))
}
