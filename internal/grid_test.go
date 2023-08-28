package internal

import "testing"

func TestGetDirection( t *testing.T){
	from := Position{X:0, Y:0}
	to := Position{X:1, Y: 1}
	result, _ := getDirection(from, to)
	if result != UpRight {
	t.Errorf("Result was incorrect, got: %d, want: %d.", result, UpRight)
	}
}
