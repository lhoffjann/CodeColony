package utils

type Position struct {
	X int
	Y int
}

func (p *Position) UpdatePosition (d Direction){
	p.X = p.X + d.Coordinates()[0]
	p.Y = p.Y + d.Coordinates()[1]
}

type Obstacle struct {
	Position Position
}

type EnergySource struct{
	Position Position
}

type World struct {
	Dimensions [2]int
	Obstacles [][]int
	EnergySources []EnergySource
}
func (w World) checkIfPositionExists(position Position){

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
