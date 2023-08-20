package utils

import (
	"container/heap"
)
// shamelessly copied from the golang docs
type Item struct {
	Value  [2]int   
	Priority int   
	Index int 
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  
	item.Index = -1 
         *pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value [2]int, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}
