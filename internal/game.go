package internal

import (
	"fmt"
	"unicode/utf8"
)

type Game struct {
	world World
	creeps []Creep
	obstacle []Obstacle
	EnergySource []EnergySource
	homeBase HomeBase
}

func NewGame() Game{
	g:= Game{}
	return g
}

func(g Game) display() {
	world:= g.world
	c:= g.creeps[0]
	cells := make([][]int, world.Dimensions[0]+1)
    for i := range cells {
        cells[i] = make([]int, world.Dimensions[1]+1)
    }
	for _, p := range returnPositionOfObstacles(world) {
		cells[p.X][p.Y] = 1	
	}
	for _, p := range returnPositionOfEnergySources(world) {
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

func returnPositionOfObstacles(w World) []Position{
	var ListOfPostion []Position
	for _, o := range w.Obstacles {
		ListOfPostion = append(ListOfPostion, o.Position)
	}
	return ListOfPostion
}
func returnPositionOfEnergySources(w World) []Position{
	var ListOfPostion []Position
	for _, o := range w.EnergySources {
		ListOfPostion = append(ListOfPostion, o.Position)
	}
	return ListOfPostion
}




