package main

import (
	"fmt"
	"math"
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
type Creep struct {
	id   int
	position [2]int
}

func (c *Creep) moveXCoordinate(move int){
	c.position[0] = c.position[0] + move
}

func (c *Creep) moveYCoordinate(move int){
	c.position[1] = c.position[1] + move 
}

func (c *Creep) moveTowardsEnergySource(e EnergySource) bool{
// make it so that the creep will not move once it is directly infront of the source
	destination := e.getPosition()
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



func main()  {
	energySource := EnergySource{
		id: 1,
		position: [2]int{10, -32},
	}

	creep := Creep {
		id : 1,
		position: [2]int{0,0},
	} 
	// this is my game engine
	for i := 0; i < 100; i++ {
		fmt.Println(creep.moveTowardsEnergySource(energySource))
		fmt.Println(creep.position)
	}
	
}
