package astar

import (
	"fmt"
)

// Point ...
type Point struct {
	X      int
	Y      int
	H      int
	G      int
	F      int
	Parent *Point
}

func (p *Point) String() string {
	return fmt.Sprintf("%v, %v", p.X, p.Y)
}
