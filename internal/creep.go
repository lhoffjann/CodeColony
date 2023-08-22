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
	newPosition:= Position{X: c.position.X + direction.Coordinates()[0], Y:c.position.Y + direction.Coordinates()[1]}
			
		fmt.Println(!c.world.PositionOccupied(newPosition))
		fmt.Println(c.world.PositionExists(newPosition))
		fmt.Println(c.world.PositionExists(newPosition) && !c.world.PositionOccupied(newPosition))
		if c.world.PositionExists(newPosition) && !c.world.PositionOccupied(newPosition) {
			c.position.UpdatePosition(direction)
		}
	}
	
	func (c Creep) GetName() string{
		return c.name
	}
	
	func (c Creep) GetPosition() Position{
		return c.position
	}
	func(c Creep) GetPathTo(position Position) []Position{
			path:= findPath(c.position, position, c.world)
		for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
			path[i], path[j] = path[j], path[i]
		}
		return path
	}

	func (c *Creep) MoveTo(point Position){
		path:= findPath(c.position, point, c.world)
		for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
			path[i], path[j] = path[j], path[i]
		}
		fmt.Println(path)
}
