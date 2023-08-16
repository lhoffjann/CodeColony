package main

import "fmt"

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




func main()  {


	creep := Creep {
		id : 1,
		position: [2]int{0,0} ,
	} 
	// this is my game engine
	for i := 0; i < 10; i++ {
		creep.moveXCoordinate(1)
		creep.moveYCoordinate(-1)
		fmt.Println(creep.position)
	}
	
}
