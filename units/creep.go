package creep

import (
	"fmt"
	"math"
	"time"
)
	
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

	type creep struct {
		name   string
		position [2]int
		maxCapacity int
		usedCapacity int
	}

	func NewCreep(position [2]int) creep{
	t := time.Now()
	c := creep {
		name: fmt.Sprintf( "creep %02d%02d", t.Minute(), t.Second()),
		position: position,
		maxCapacity: 100,
		usedCapacity: 0,
		}
	return c
	
	}

	func (c *creep) Move(direction Direction) {
		c.position[0] = c.position[0] + direction.Coordinates()[0]
		c.position[1] = c.position[1] + direction.Coordinates()[1]
	}
	
	func (c creep) GetName() string{
		return c.name
	}
	
	func (c creep) GetPosition() [2]int{
		return c.position
	}

	func (c *creep) MoveTo(point [2]int){
	// make it so that the creep will not move once it is directly infront of the source
		ex,ey := point[0], point[1]
		cx,cy := c.position[0], c.position[1]
		dx := ex - cx
		dy := ey - cy
		if math.Abs(float64(dx)) == 1 && math.Abs(float64(dy)) == 1 {
			//return true
		}
		move := func (dist int) int {
			if dist >= 0 {
				return 1
			}
			return -1
		}
		move(1.0)
		c.Move(Up)
		
		if math.Abs(float64(dx)) > math.Abs(float64(dy)) { // there is probably some edgecase here.
			// c.moveXCoordinate(move(dx))	
		} else {
			//	c.moveYCoordinate(move(dy))
	}
//	return false
}
