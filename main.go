package main

import (
	"CodeColony/internal"
)

// we got a map that is build as a coordinate system.

// we got a main building that can produce creeps

// we also got creeps that can move, collect and carry
// there should also be natural structures that the creep has to move around
// also there should definitly be more than one creep

func main() {
	g := internal.NewGame(10, 10)
	g.Tick()
}
