package internal 

import (
	"fmt"
	"time"
)
	
type State int8
const (
	idle = 1
	moving = 2
	collecting = 3
	storing = 4
)
	type Creep struct {
		name   string
		state State
		position Position
		maxCapacity int
		usedCapacity int
		memory []func()
		world World	
}

	func NewCreep(position Position, world World) Creep{
	t := time.Now()
	c := Creep {
		name: fmt.Sprintf( "creep %02d%02d", t.Minute(), t.Second()),
		state: idle,	
		position: position,
		maxCapacity: 100,
		usedCapacity: 0,
		world: world,
		
		}
	return c
	
	}
	func(c *Creep) SetPosition(position Position){
		c.position = position
	}

	func (c *Creep) Move(direction Direction) {
	p := c.position
	d := direction.Coordinates()
	newPosition:= Position{X: p.X + d[0], Y:p.Y + d[1]}
		if c.world.PositionExists(newPosition) && !c.world.PositionOccupied(newPosition) {
			p.UpdatePosition(direction)
		}
	}
	
	func (c Creep) GetName() string{
		return c.name
	}
	
	func (c Creep) GetPosition() Position{
		return c.position
	}
	func(c Creep) GetPathTo(position Position) []Position{
		return findPath(c.position, position, c.world)
	}
	func (c *Creep) MoveToStructure(structure Structure){
		path := c.GetPathTo(structure.returnPostion())
		var actionList []func()
		for i, v := range path {
			if i == 0 {
				fmt.Println(c.position)
				fmt.Println(v)
				x,y :=c.position.X - v.X, c.position.Y - v.Y
				
			} 
			
		}
	} 
	
	func (c *Creep) MoveToPoint(point Position){
		//path:= findPath(c.position, point, c.world)
	}
