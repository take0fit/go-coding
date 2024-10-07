package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type ListNodeHeap []*ListNode

func (h ListNodeHeap) Len() int           { return len(h) }
func (h ListNodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h ListNodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ListNodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *ListNodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := &ListNodeHeap{}
	heap.Init(h)

	for _, list := range lists {
		if list != nil {
			heap.Push(h, list)
		}
	}

	dummy := &ListNode{}
	current := dummy

	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		current.Next = node
		current = current.Next
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}

	return dummy.Next
}

func main() {
	// リストの例
	l1 := &ListNode{1, &ListNode{4, &ListNode{5, nil}}}
	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	l3 := &ListNode{2, &ListNode{6, nil}}

	lists := []*ListNode{l1, l2, l3}

	merged := mergeKLists(lists)
	fmt.Print("マージ結果: ")
	for merged != nil {
		fmt.Print(merged.Val, " ")
		merged = merged.Next
	}
}
