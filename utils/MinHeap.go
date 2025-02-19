package utils

import "container/heap"

type MinHeap[T any] []*Item[T]

type Item[T any] struct {
	Value    T
	Priority int
	index    int
}

func NewMinHeap[T any]() *MinHeap[T] {
	mh := &MinHeap[T]{}
	heap.Init(mh)
	return mh
}

func (pq MinHeap[T]) Len() int { return len(pq) }

func (pq MinHeap[T]) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq MinHeap[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *MinHeap[T]) Push(x any) {
	n := len(*pq)
	item := x.(*Item[T])
	item.index = n
	*pq = append(*pq, item)
}

func (pq *MinHeap[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
