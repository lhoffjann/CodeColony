package internal

import (
	"container/heap"
	"fmt"
	"math"
)

func heuristic(a Position, b Position) float64 {
	return (math.Abs(float64(a.X)) - math.Abs(float64(b.X))) + (math.Abs(float64(a.Y)) - math.Abs(float64(b.Y)))
}

// need to be refactored

func findPath(start Position, goal Position, w World) []Position {
	frontier := make(PriorityQueue, 1)
	frontier[0] = &Item{
		Value:    start,
		Priority: 1.0,
		Index:    0,
	}
	heap.Init(&frontier)
	emptyPosition := Position{}
	cameFrom := map[Position]Position{start: emptyPosition}
	costSoFar := map[Position]float64{start: 0}
	for frontier.Len() != 0 {
		current := heap.Pop(&frontier).(*Item)
		if current.Value == goal {
			break

		}
		for _, next := range w.ReturnFreeNeighbors(current.Value) {

			newCost := costSoFar[current.Value] + 1.0

			_, ok := costSoFar[next]
			if !ok || newCost < costSoFar[next] {
				fmt.Printf("%d: newCost %f, oldCost %f \n", next, newCost, costSoFar[next])
				costSoFar[next] = newCost

				priority := newCost + heuristic(goal, next)
				fmt.Printf("The new Priority is %f", priority)
				item := &Item{
					Value:    next,
					Priority: priority,
				}
				heap.Push(&frontier, item)
				cameFrom[next] = current.Value
			}
		}
	}
	current := goal
	var path []Position
	for current != start {
		path = append(path, current)
		current = cameFrom[current]
	}
	path = append(path, start)
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}

/*
frontier = PriorityQueue()
frontier.put(start, 0)
came_from = dict()
cost_so_far = dict()
came_from[start] = None
cost_so_far[start] = 0

while not frontier.empty():
   current = frontier.get()

   if current == goal:
      break

   for next in graph.neighbors(current):
      new_cost = cost_so_far[current] + graph.cost(current, next)
      if next not in cost_so_far or new_cost < cost_so_far[next]:
         cost_so_far[next] = new_cost
         priority = new_cost + heuristic(goal, next)
         frontier.put(next, priority)
         came_from[next] = current
*/
