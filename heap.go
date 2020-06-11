package main

import (
	"container/heap"
)

type Leaf struct {
	Value interface{}
	Count int
	Index int
}

type LeafHeap []*Leaf

func (h LeafHeap) Len() int { return len(h) }

func (h LeafHeap) Less(i, j int) bool { return h[i].Count < h[j].Count }

func (h LeafHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].Index = i
	h[j].Index = j
}

func (h *LeafHeap) Push(v interface{}) {
	n := len(*h)
	elem := v.(*Leaf)
	elem.Index = n
	*h = append(*h, elem)
}

func (h *LeafHeap) Pop() interface{} {
	old := *h
	n := len(old)
	elem := old[n-1]
	old[n-1] = nil
	elem.Index = -1
	*h = old[0 : n-1]
	return elem
}

func (h *LeafHeap) Update(l *Leaf) {
	heap.Fix(h, l.Index)
}
