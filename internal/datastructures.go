package internal


type Position struct {
	X int
	Y int
}
func (p Position) AreEqual(position Position) bool{
	return p.X == position.X && p.Y == position.Y	
}

func (p *Position) UpdatePosition (d Direction){
	x := p.X + d.Coordinates()[0]
	y := p.Y + d.Coordinates()[1]
	p.X = x 	
	p.Y = y
}

type Obstacle struct {
	Position Position
}

type EnergySource struct{
	Position Position
}

type World struct {
	Dimensions [2]int
	Obstacles []Obstacle
	EnergySources []EnergySource
}

func (w World) PositionExists (position Position) bool {
	return  !(position.X < 0 || position.X > w.Dimensions[0] || position.Y < 0 || position.Y > w.Dimensions[1])
}

func (w World) PositionOccupied (position Position) bool{

	for _, o := range w.Obstacles {
		if o.Position.AreEqual(position) {
			return o.Position.AreEqual(position)
		}  
	}
	for _, es := range w.EnergySources {
		if es.Position.AreEqual(position) {
			return es.Position.AreEqual(position)
		}  
	}
	return false
}

func (w World) ReturnNeighbors(position Position) []Position{
	var neighbors []Position   
	for i := UpLeft; i <= DownRight; i++ {
		neighbor := Position{X: position.X + i.Coordinates()[0], Y:position.X + i.Coordinates()[1],}
		if w.PositionExists(neighbor) {
			neighbors = append(neighbors, neighbor)
		}
	}
	return neighbors
}

func (w World) ReturnFreeNeighbors(position Position) []Position {
	var freeNeighbors []Position
	neighbors := w.ReturnNeighbors(position)
	for _ , p := range neighbors {
		if !w.PositionOccupied(p) {
			freeNeighbors = append(freeNeighbors, p)
		}
	}
	return freeNeighbors
}

type Direction int
const (
	UpLeft Direction = iota
	Up 
	UpRight 
	Right 
	Left 
	DownLeft
	Down 
	DownRight
	)

func (d Direction) Coordinates() [2]int {
	switch d {
		case UpLeft:
			return [2]int{-1,1} 
		case Up: 
			return [2]int{1,0} 
		case UpRight:
			return [2]int{1,1}
		case Left:
			return [2]int{-1,0}
		case Right:
			return [2]int{1,0}
		case DownLeft:
			return [2]int{-1,-1}
		case Down:
			return [2]int{0,-1}
		case DownRight:
			return [2]int{1,-1}
		}
		return [2]int{0,0}
	}
