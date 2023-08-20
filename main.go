package main

import (
	"CodeColony/units"
	"CodeColony/utils"
	"fmt"
)

// we got a map that is build as a coordinate system.

// we got a main building that can produce creeps

// we also got creeps that can move, collect and carry
// there should also be natural structures that the creep has to move around
// also there should definitly be more than one creep



func main()  {
	energySource := utils.EnergySource{
		Position: utils.Position{X: 10,Y: 32,},
	}
	c := creep.NewCreep(utils.Position{X:1, Y:1,})	
	// this is my game engine
	fmt.Println(c.GetName())
	for i := 0; i < 10; i++ {
		c.Move(utils.UpRight)
		c.MoveTo(energySource.Position)
		fmt.Println(c.GetPosition())
	}
}
