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
type Task func()
	
type Creep struct {
		name   string
		state State
		position *Position
		maxCapacity int
		usedCapacity int
		memory []Task
		world World	
}

func NewCreep(position *Position, world World) *Creep{
	t := time.Now()
	c := &Creep {
		name: fmt.Sprintf( "creep %02d%02d", t.Minute(), t.Second()),
		state: idle,	
		position: position,
		maxCapacity: 100,
		usedCapacity: 0,
		memory: make([]Task,0),
		world: world,
		}
	return c
	
	}
func(c *Creep) SetPosition(position Position){
	c.position = &position
}
func (c Creep) GetName() string{
	return c.name
}
	
func (c Creep) GetPosition() Position{
	return *c.position
}

func (c *Creep) Move(direction Direction) {
	p := c.position
	d := direction.Coordinates()
	newPosition:= Position{X: p.X + d[0], Y:p.Y + d[1]}
	fmt.Println(newPosition)
	if c.world.PositionExists(newPosition) && !c.world.PositionOccupied(newPosition) {
		p.UpdatePosition(direction)
	}
}

func(c *Creep) executeTaskInMemory(){
	if len(c.memory) > 0 {
		c.memory[0]()
		fmt.Println(c.GetPosition())
		c.memory = c.memory[1:]
	}
}

func (c *Creep) MoveToStructure(position Position){
	path := findPath(*c.position, position, c.world)
	var actionList []Task
	for i := 0; i < len(path)-1; i++ {
		if i == 0 {
			direction, err := getDirection(*c.position, path[i])
			if err != nil {
				fmt.Printf("Could not append Direction")
				return
			}
		actionList = append(actionList,func() {c.Move(direction)})
		}else {
			direction, err := getDirection(path[i], path[i+1])
			if err != nil {
				fmt.Printf("Could not append Direction")
				return
		}
		actionList = append(actionList,func(){
		c.Move(direction)
		fmt.Println(direction.Coordinates())})
		} 
	c.memory = actionList	
	}
}

