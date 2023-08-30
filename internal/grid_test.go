package internal

import (
	"fmt"
	"testing"
)

func TestGetDirection(t *testing.T) {
      // Defining the columns of the table
        var tests = []struct {
		name string
		from Position
		to Position
		want  Direction
        }{
            // the table itself
            {"Should be UpLeft", Position{X:1, Y:1}, Position{X:0, Y:2}, UpLeft},
            {"Should be Up", Position{X:1, Y:1}, Position{X:1, Y:2}, Up},
            {"Should be UpRight", Position{X:1, Y:1}, Position{X:2, Y:2}, UpRight},
            {"Should be Left", Position{X:1, Y:1}, Position{X:0, Y:1}, Left},
            {"Should be Right", Position{X:1, Y:1}, Position{X:2, Y:1}, Right},
            {"Should be DownLeft", Position{X:1, Y:1}, Position{X:0, Y:0}, DownLeft},
            {"Should be Down", Position{X:1, Y:1}, Position{X:1, Y:0}, Down},
            {"Should be DownRight", Position{X:1, Y:1}, Position{X:2, Y:0}, DownRight},
        }
      // The execution loop
        for _, tt := range tests {
            t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.want)
                ans, _ := getDirection(tt.from, tt.to)
                if ans != tt.want {
                    t.Errorf("got %d, want %d", ans, tt.want)
                }
            })
        }
    }
