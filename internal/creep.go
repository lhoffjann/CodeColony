package internal 

import (
	"fmt"
	"math"
	"time"
)
	

	type Creep struct {
		name   string
		position Position
		maxCapacity int
		usedCapacity int
	}

	func NewCreep(position Position) Creep{
	t := time.Now()
	c := Creep {
		name: fmt.Sprintf( "creep %02d%02d", t.Minute(), t.Second()),
		position: position,
		maxCapacity: 100,
		usedCapacity: 0,
		}
	return c
	
	}

	func (c *Creep) Move(direction Direction) {
		 c.position.UpdatePosition(direction)
	}
	
	func (c Creep) GetName() string{
		return c.name
	}
	
	func (c Creep) GetPosition() Position{
		return c.position
	}

	func (c *Creep) MoveTo(point Position){
	// make it so that the creep will noutils.Position once it is directly infront of the source
		dx := point.X - c.position.X
		dy := point.Y - c.position.Y
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
//	returutils.Positione
}
