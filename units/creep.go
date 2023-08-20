package creep

import (
	"CodeColony/utils"
	"fmt"
	"math"
	"time"
)
	

	type creep struct {
		name   string
		position utils.Position
		maxCapacity int
		usedCapacity int
	}

	func NewCreep(position utils.Position) creep{
	t := time.Now()
	c := creep {
		name: fmt.Sprintf( "creep %02d%02d", t.Minute(), t.Second()),
		position: position,
		maxCapacity: 100,
		usedCapacity: 0,
		}
	return c
	
	}

	func (c *creep) Move(direction utils.Direction) {
		 c.position.UpdatePosition(direction)
	}
	
	func (c creep) GetName() string{
		return c.name
	}
	
	func (c creep) GetPosition() utils.Position{
		return c.position
	}

	func (c *creep) MoveTo(point utils.Position){
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
		c.Move(utils.Up)
		
		if math.Abs(float64(dx)) > math.Abs(float64(dy)) { // there is probably some edgecase here.
			// c.moveXCoordinate(move(dx))	
		} else {
			//	c.moveYCoordinate(move(dy))
	}
//	returutils.Positione
}
