package main

import (
	"CodeColony/units"
	"fmt"
)
// we got a map that is build as a coordinate system.

// we got a main building that can produce creeps

// we also got creeps that can move, collect and carry
// there should also be natural structures that the creep has to move around
// also there should definitly be more than one creep

type Obstacles struct {
	listOfObstacles [][]int
}

type EnergySource struct{
	id int
	position [2]int
}

func (e EnergySource) getPosition() [2]int {
	return e.position
}
// should the creep also have a orientation? thatwould mean it can move only forward and has to turn.
func main()  {
	energySource := EnergySource{
		id: 1,
		position: [2]int{10, -32},
	}
	c := creep.NewCreep([2]int{0,0})	
	// this is my game engine
	fmt.Println(c.GetName())
	for i := 0; i < 10; i++ {
		c.MoveTo(energySource.getPosition())
		fmt.Println(c.GetPosition())
	}
}
