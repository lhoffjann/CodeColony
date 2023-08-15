package main

import "fmt"

// we got a map that is build as a coordinate system.
// we got a main building that can produce creeps
// we also got creeps that can move, collect and carry

type Creep struct {
	id   int
	position [2]int
}

func (c *Creep) moveUpwards(){
	fmt.Println(c.position[1])
	c.position[1] = c.position[1] + 1 
}


func main()  {


	creep := Creep {
		id : 1,
		position: [2]int{0,0} ,
	} 
	// this is my game engine
	for i := 0; i < 10; i++ {
		creep.moveUpwards()
		fmt.Println(creep.position)

	}
	
}
