package main

import (
	"fmt"

	"github.com/shanghuiyang/a-star/astar"
	"github.com/shanghuiyang/a-star/scene"
)

func main() {

	// build a scene map with walls
	r, c := 20, 20
	s := scene.New(r, c)
	for i := 4; i < 13; i++ {
		s.SetWall(9, i)
	}

	// define the origin and destination
	org := &astar.Point{X: 3, Y: 3}
	des := &astar.Point{X: 15, Y: 15}

	// find the path using a-star algorithm
	a := astar.New(s)
	path, err := a.FindPath(org, des)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	// draw the scene with the path
	a.Draw()
	fmt.Printf("path: %v\n\n", path)

	// find another path
	org, des = des, org
	path, err = a.FindPath(org, des)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	a.Draw()
	fmt.Printf("path: %v\n\n", path)

	// or, build a scene map from a string
	str := `
################
#              #
#      #       #
#      #       #
#      #       #
#      #       #
################
`
	s = scene.BuildFromStr(str)

	org = &astar.Point{X: 2, Y: 2}
	des = &astar.Point{X: 5, Y: 13}
	a = astar.New(s)
	path, err = a.FindPath(org, des)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	a.Draw()
	fmt.Printf("path: %v\n\n", path)
	return
}
