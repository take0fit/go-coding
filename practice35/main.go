package main

import "fmt"

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 5
	fmt.Println("インデックス:", binarySearch(nums, target))
}
