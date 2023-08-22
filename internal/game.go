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

func NewGame(x int, y int) Game{
	energySource := EnergySource{
		Position: Position{X: 10,Y: 10,},
	}
	obstacle:= Obstacle{
		Position: Position{X:5, Y: 5},
	}
	obstacle1 := Obstacle{
		Position: Position{X:0, Y: 10},
	}
	obstacle2:= Obstacle{
		Position: Position{X:0, Y: 1},
	}
	obstacle3:= Obstacle{
		Position: Position{X:1, Y: 1},
	}

	energySources:=[]EnergySource{energySource}
	obstacles:= []Obstacle{obstacle, obstacle1, obstacle2,obstacle3}
	world := World{
		Dimensions: [2]int{x, y},
		EnergySources: energySources,
		Obstacles: obstacles,
	}
	c := NewCreep(Position{X:0, Y:0,},  world)
	g:= Game{
		world: world,
		obstacle: obstacles,
		creeps: []Creep{c},
	}
	return g
}
func(g Game) tick(){

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




