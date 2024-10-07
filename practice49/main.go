package main

import "fmt"

func permute(nums []int) [][]int {
	var result [][]int
	var backtrack func(int)
	backtrack = func(first int) {
		if first == len(nums) {
			perm := make([]int, len(nums))
			copy(perm, nums)
			result = append(result, perm)
			return
		}
		for i := first; i < len(nums); i++ {
			nums[first], nums[i] = nums[i], nums[first]
			backtrack(first + 1)
			nums[first], nums[i] = nums[i], nums[first]
		}
	}
	backtrack(0)
	return result
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println("順列:", permute(nums))
}
