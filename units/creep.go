package creep

import (
	"fmt"
	"math"
	"time"
)
	type creep struct {
		name   string
		position [2]int
	}

	func NewCreep(position [2]int) creep{
	t := time.Now()
	c := creep {
		name: fmt.Sprintf( "creep %02d%02d", t.Minute(), t.Second()),
		position: position,
		}
	return c
	
	}
	func (c *creep) moveXCoordinate(move int){
		c.position[0] = c.position[0] + move
	}

	func (c *creep) moveYCoordinate(move int){
		c.position[1] = c.position[1] + move 
	}
	
	func (c creep) GetName() string{
		return c.name
	}
	
	func (c creep) GetPosition() [2]int{
		return c.position
	}
	//func (c Creep) checkIfNextToEnergySource(e []EnergySource) bool {
	//for _, v := range e {
	//		
	//	}	
	//}

	func (c *creep) MoveTowardsPoint(destination [2]int) bool{
	// make it so that the creep will not move once it is directly infront of the source
		ex,ey := destination[0], destination[1]
		cx,cy := c.position[0], c.position[1]
		dx := ex - cx
		dy := ey - cy
		if math.Abs(float64(dx)) == 1 && math.Abs(float64(dy)) == 1 {
			return true
		}
		move := func (dist int) int {
			if dist >= 0 {
				return 1
			}
			return -1
		}

		if math.Abs(float64(dx)) > math.Abs(float64(dy)) { // there is probably some edgecase here.
			c.moveXCoordinate(move(dx))	
		} else {
			c.moveYCoordinate(move(dy))
	}
	return false
}
