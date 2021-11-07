package main

import (
	"fmt"

	"github.com/shanghuiyang/astar"
	"github.com/shanghuiyang/astar/tilemap"
)

// a map with 10 rows and 20 cols
const strmap = `
####################
#                  #
#                  #
#   #########      #
#                  #
#        #######   #
#                  #
#                  #
#                  #
####################
`

func main() {
	// build a map from string
	m := tilemap.BuildFromStr(strmap)

	// define the origin and destination
	org := &astar.Point{X: 7, Y: 2}
	des := &astar.Point{X: 1, Y: 16}

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
}
