package internal

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
	"unicode/utf8"
)

type Game struct {
	world        World
	creeps       *Creep
	obstacle     []Obstacle
	EnergySource []EnergySource
	homeBase     HomeBase
}

func generateRandomPositions(x, y, n int) []Position {
	rand.Seed(time.Now().UnixNano())
	positions := make([]Position, n)

	for i := 0; i < n; i++ {
		var newPos Position

		for {
			newPos = Position{X: rand.Intn(x), Y: rand.Intn(y)}
			valid := true
			for j := 0; j < i; j++ {
				dx := newPos.X - positions[j].X
				dy := newPos.Y - positions[j].Y
				if dx*dx+dy*dy < (x/2)*(y/2) {
					valid = false
					break
				}
			}

			if valid {
				break
			}
		}

		positions[i] = newPos
	}

	return positions
}

func generateUniquePositions(n, maxX, maxY int, existingPositions []Position) []Position {
	rand.Seed(time.Now().UnixNano())
	positions := make([]Position, 0, n)

	usedPositions := make(map[Position]bool)
	for _, v := range existingPositions {
		usedPositions[v] = true
	}

	for len(positions) < n {
		x := rand.Intn(maxX) // Replace maxX with the maximum X coordinate you want
		y := rand.Intn(maxY) // Replace maxY with the maximum Y coordinate you want
		newPosition := Position{x, y}

		// Check if the generated position is unique
		if !usedPositions[newPosition] {
			positions = append(positions, newPosition)
			usedPositions[newPosition] = true
		}
	}

	return positions
}

func generateMap(x, y int) (obstacles []Obstacle, energySources []EnergySource) {
	rand.Seed(time.Now().UnixNano())
	countObstacles := rand.Intn(x*y/2-x*y/5) + x*y/5
	structures := generateRandomPositions(x, y, 2)
	homebase := EnergySource{Position: structures[0]}
	energySource := EnergySource{Position: structures[1]}
	positions := generateUniquePositions(countObstacles, x, y, structures)
	for _, v := range positions {
		obstacles = append(obstacles, Obstacle{Position: v})
	}

	energySources = []EnergySource{homebase, energySource}
	return obstacles, energySources

}

func NewGame(x, y int) *Game {
	obstacles, energySources := generateMap(x, y)
	world := World{
		Dimensions:    [2]int{x, y},
		EnergySources: energySources,
		Obstacles:     obstacles,
	}
	c := NewCreep(&Position{X: 0, Y: 0}, world)
	g := &Game{
		world:    world,
		obstacle: obstacles,
		creeps:   c,
	}
	return g
}
func (g *Game) Tick() {
	creep := g.creeps
	creep.MoveToStructure(Position{9, 9})
	g.display()
	for i := 0; i < 100; i++ {
		fmt.Scanln()
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
		creep.executeTaskInMemory()
		fmt.Printf("This is my current Postion %d\n", creep.position)
		g.display()
	}
}

func (g *Game) display() {
	world := g.world
	c := g.creeps
	cells := make([][]int, world.Dimensions[1]+1)
	for i := range cells {
		cells[i] = make([]int, world.Dimensions[0]+1)
	}
	for _, p := range returnPositionOfObstacles(world) {
		cells[p.Y][p.X] = 1
	}
	for _, p := range returnPositionOfEnergySources(world) {
		cells[p.Y][p.X] = 2
	}
	cells[c.GetPosition().Y][c.GetPosition().X] = 3
	for _, row := range cells {
		for _, cell := range row {
			switch cell {
			case 0:
				fmt.Print(" " + getStringForEmoji("\u2b1c"))
			case 1:
				fmt.Print(" " + getStringForEmoji("\u26f0\ufe0f") + " ")
			case 2:
				fmt.Print(" " + getStringForEmoji("\U0001f536"))
			case 3:
				fmt.Print(" " + getStringForEmoji("\U0001f47e"))
			}
		}
		fmt.Println()
	}
}
func getStringForEmoji(s string) string {
	r, _ := utf8.DecodeRuneInString(s)
	return string(r)
}

func returnPositionOfObstacles(w World) []Position {
	var ListOfPostion []Position
	for _, o := range w.Obstacles {
		ListOfPostion = append(ListOfPostion, o.Position)
	}
	return ListOfPostion
}
func returnPositionOfEnergySources(w World) []Position {
	var ListOfPostion []Position
	for _, o := range w.EnergySources {
		ListOfPostion = append(ListOfPostion, o.Position)
	}
	return ListOfPostion
}
