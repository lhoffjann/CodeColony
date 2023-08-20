package main

import (
	"CodeColony/internal"
	"fmt"
	"unicode/utf8"
	"os"
	"os/exec"
)

// we got a map that is build as a coordinate system.

// we got a main building that can produce creeps

// we also got creeps that can move, collect and carry
// there should also be natural structures that the creep has to move around
// also there should definitly be more than one creep



func main()  {
	energySource := internal.EnergySource{
		Position: internal.Position{X: 10,Y: 10,},
	}
	obstacle:= internal.Obstacle{
		Position: internal.Position{X:5, Y: 5},
	}

	world := internal.World{
		Dimensions: [2]int{10, 10},
		EnergySources: []internal.EnergySource{energySource},
		Obstacles: [] internal.Obstacle{obstacle},
		
	}
	c := internal.NewCreep(internal.Position{X:0, Y:0,},  world)	
	fmt.Println(c.GetName())
	printBoard(world,c)
	for i := 0; i < 10; i++ {
		fmt.Scanln()
		c.Move(internal.UpRight)
		fmt.Println(c.GetPosition())
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
		printBoard(world, c)
	}
}


func printBoard(w internal.World, c internal.Creep){
cells := make([][]int, w.Dimensions[0]+1)
    for i := range cells {
        cells[i] = make([]int, w.Dimensions[1]+1)
    }
	for _, p := range returnPositionOfObstacles(w) {
		cells[p.X][p.Y] = 1	
	}
	for _, p := range returnPositionOfEnergySources(w) {
		cells[p.X][p.Y] = 2	
	}
	cells[c.GetPosition().X][c.GetPosition().Y] = 3
	for _, row := range cells {
		for _, cell := range row {
		switch cell {
			case 0:
				fmt.Print(" " + getStringForEmoji("\u2b1c"))
			case 1:
				fmt.Print(" "+  getStringForEmoji("\u26f0\ufe0f")+ " ")
			case 2:
				fmt.Print(" " + getStringForEmoji("\U0001f536"))
			case 3:
				fmt.Print(" " + getStringForEmoji("\U0001f47e"))
			}
		}
		fmt.Println()
	}
}
func getStringForEmoji(s string) string{
	r ,_ := utf8.DecodeRuneInString(s)
return string(r)
}

func returnPositionOfObstacles(w internal.World) []internal.Position{
	var ListOfPostion []internal.Position
	for _, o := range w.Obstacles {
		ListOfPostion = append(ListOfPostion, o.Position)
	}
	return ListOfPostion
}
func returnPositionOfEnergySources(w internal.World) []internal.Position{
	var ListOfPostion []internal.Position
	for _, o := range w.EnergySources {
		ListOfPostion = append(ListOfPostion, o.Position)
	}
	return ListOfPostion
}

