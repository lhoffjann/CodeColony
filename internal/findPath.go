package internal

import (
	"container/heap"
	"fmt"
	"math"
)

type Item struct {
	Value    Position
	Priority float64
	Index    int
	cost     float64
	parent   *Item
	open     bool
	closed   bool
}
type nodeMap map[Position]*Item

// get gets the Pather object wrapped in a node, instantiating if required.
func (nm nodeMap) get(p Position) *Item {
	n, ok := nm[p]
	if !ok {
		n = &Item{
			Value: p,
		}
		nm[p] = n
	}
	return n
}
func heuristic(a Position, b Position) float64 {
	return (math.Abs(float64(a.X)) - math.Abs(float64(b.X))) + (math.Abs(float64(a.Y)) - math.Abs(float64(b.Y)))
}

// need to be refactored

func findPath(start, goal Position, w World) (path []Position, found bool) {
	frontier := &PriorityQueue{}
	nm := nodeMap{}
	heap.Init(frontier)
	fromNode := nm.get(start)
	fromNode.open = true
	heap.Push(frontier, fromNode)
	for {
		if frontier.Len() == 0 {
			return
		}
		current := heap.Pop(frontier).(*Item)
		current.open = false
		current.closed = true
		if current == nm.get(goal) {
			curr := current
			p := []Position{}
			for curr != nil {
				p = append(p, curr.Value)
				curr = curr.parent
			}
			for i, j := 0, len(p)-1; i < j; i, j = i+1, j-1 {
				p[i], p[j] = p[j], p[i]
			}
			fmt.Println(p)
			return p, true
		}
		for _, neighbor := range w.ReturnFreeNeighbors(current.Value) {
			cost := current.cost + 1.0
			neighborNode := nm.get(neighbor)
			if cost < neighborNode.cost {
				if neighborNode.open {
					heap.Remove(frontier, neighborNode.Index)
				}
				neighborNode.open = false
				neighborNode.closed = false
			}
			if !neighborNode.open && !neighborNode.closed {
				neighborNode.cost = cost
				neighborNode.open = true
				neighborNode.Priority = cost + heuristic(goal, neighborNode.Value)
				neighborNode.parent = current
				heap.Push(frontier, neighborNode)
			}
		}
	}
}
