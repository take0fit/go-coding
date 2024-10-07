package main

import (
	"container/heap"
	"fmt"
)

type Element struct {
	Value int
	Count int
}

type MinHeap []Element

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Count < h[j].Count }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Element))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	h := &MinHeap{}
	heap.Init(h)

	for value, count := range freqMap {
		heap.Push(h, Element{Value: value, Count: count})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	var result []int
	for h.Len() > 0 {
		result = append(result, heap.Pop(h).(Element).Value)
	}

	return result
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println("上位", k, "個の要素:", topKFrequent(nums, k))
}
