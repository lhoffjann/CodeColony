package main

import (
	"fmt"
	"math"
)

// we got a map that is build as a coordinate system.
// we got a main building that can produce creeps
// we also got creeps that can move, collect and carry
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

func (c *Creep) moveTowardsEnergySource(e EnergySource){
	destination := e.getPosition()
	ex,ey := destination[0], destination[1]
	cx,cy := c.position[0], c.position[1]
	dx := ex - cx
	fmt.Printf("dx is %d \n", dx)
	dy := ey - cy
	fmt.Printf("dy is %d \n", dy)
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

	


}



func main()  {
	energySource := EnergySource{
		id: 1,
		position: [2]int{10, -32},
	}

	creep := Creep {
		id : 1,
		position: [2]int{0,0} ,
	} 
	// this is my game engine
	for i := 0; i < 100; i++ {
		creep.moveTowardsEnergySource(energySource)
		fmt.Println(creep.position)
	}
	
}
