package main

import (
	"container/heap"
	"fmt"
)

// MinHeap implements the heap.Interface and holds the items.
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	// Initialize a min-heap.
	h := &MinHeap{28, 17, 5}
	heap.Init(h)

	// Push and Pop operations.
	heap.Push(h, 3)
	fmt.Printf("Minimum: %d\n", (*h)[0])

	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
