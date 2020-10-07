package main

import (
	"fmt"

	"github.com/shanghuiyang/a-star/astar"
	"github.com/shanghuiyang/a-star/scene"
)

func main() {

	r, c := 20, 20
	s := scene.New(r, c)
	for i := 4; i < 13; i++ {
		s.SetWall(9, i)
	}

	org := &astar.Point{X: 3, Y: 3}
	des := &astar.Point{X: 15, Y: 15}
	a := astar.New(org, des, s)
	path, err := a.Run()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	s.Draw()
	fmt.Println("path: ", path)
}
