package main

import (
	"fmt"

	"github.com/shanghuiyang/a-star/astar"
	"github.com/shanghuiyang/a-star/tilemap"
)

func main() {

	// build a tilemap with walls
	r, c := 20, 20
	m := tilemap.New(r, c)
	for x := 4; x < 13; x++ {
		m.SetWall(9, x)
	}

	// define the origin and destination
	org := &astar.Point{X: 3, Y: 3}
	des := &astar.Point{X: 15, Y: 15}

	// find the path using a-star algorithm
	a := astar.New(m)
	path, err := a.FindPath(org, des)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	// draw the tilemap with the path
	a.Draw()
	fmt.Printf("path: %v\n\n", path)
	return
}
