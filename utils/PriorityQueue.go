package utils

import "container/heap"

type PriorityQueue[T any] []*Item[T]

type Item[T any] struct {
	Value    T
	Priority int
	index    int
}

func (pq PriorityQueue[T]) Len() int { return len(pq) }

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue[T]) Push(x any) {
	n := len(*pq)
	item := x.(*Item[T])
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

type MinHeap[T any] struct {
	PriorityQueue[T]
}

func NewMinHeap[T any]() *MinHeap[T] {
	mh := &MinHeap[T]{}
	heap.Init(mh)
	return mh
}

func (mh *MinHeap[T]) Less(i, j int) bool {
	return mh.PriorityQueue[i].Priority < mh.PriorityQueue[j].Priority
}

type MaxHeap[T any] struct {
	PriorityQueue[T]
}

func NewMaxHeap[T any]() *MaxHeap[T] {
	mh := &MaxHeap[T]{}
	heap.Init(mh)
	return mh
}

func (mh *MaxHeap[T]) Less(i, j int) bool {
	return mh.PriorityQueue[i].Priority > mh.PriorityQueue[j].Priority
}
