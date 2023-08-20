package internal

import (
	"container/heap"
	"math"
)
func heuristic (a [2]float64, b [2]float64) float64 {
	return math.Abs(a[0]-b[0])+math.Abs(a[1]-b[1])
}



func findPath(start [2]int, goal [2]int) [][2]int {
	frontier := make(PriorityQueue,1)
	frontier.Push(
		Item{
			Value: start,
			Priority: 1,
			Index: 0,
	})
	heap.Init(&frontier)
	// cameFrom := map[[2]int][2]int{start: [2]int{}}
	// costSoFar := map[[2]int]int{start:0}
	for frontier.Len()!= 0 {
		current := heap.Pop(&frontier).(*Item)
		if current.Value == goal{
			break
		}

	}
	return [][2]int{}
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
