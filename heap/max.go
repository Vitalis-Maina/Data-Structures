package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	// Initialize a max-heap.
	h := &MaxHeap{2, 1, 5}
	heap.Init(h)

	// Push and Pop operations.
	heap.Push(h, 3)
	fmt.Printf("Maximum: %d\n", (*h)[0])

	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
