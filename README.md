# A-Star
[![Build Status](https://travis-ci.org/shanghuiyang/a-star.svg?branch=main)](https://travis-ci.org/shanghuiyang/a-star)

A-Star algorithm implemented with Go.

## Usage
see the [main.go](/main.go) for complete usage.

build a scene map.
```go
package main

import (
	"fmt"

	"github.com/shanghuiyang/a-star/scene"
)

func main() {

	// build a scene map with walls
	r, c := 20, 20
	s := scene.New(r, c)
	for i := 4; i < 13; i++ {
		s.SetWall(9, i)
	}
	s.Draw()
	//
	// ####################
	// #                  #
	// #                  #
	// #                  #
	// #                  #
	// #                  #
	// #                  #
	// #                  #
	// #                  #
	// #   #########      #
	// #                  #
	// #                  #
	// #                  #
	// #                  #
	// #                  #
	// #                  #
	// #                  #
	// #                  #
	// #                  #
	// ####################
	//


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
	s.Draw()
	//
 	// ################
	// #              #
	// #      #       #
	// #      #       #
	// #      #       #
	// #      #       #
	// ################
	//
	return
}
```


find a path using the scene map.
```go
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

	org := &astar.Point{X: 3, Y: 3}		// origin
	des := &astar.Point{X: 15, Y: 15}	// destination
	a := astar.New(s)
	path, err := a.FindPath(org, des)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	a.Draw()
	//
	// ####################
	// #                  #
	// #                  #
	// #  A               #
	// #   *              #
	// #    *             #
	// #     *            #
	// #      *           #
	// #       *****      #
	// #   #########*     #
	// #             *    #
	// #              *   #
	// #              *   #
	// #              *   #
	// #              *   #
	// #              B   #
	// #                  #
	// #                  #
	// #                  #
	// ####################
	//

	// swap origin and destination, find another path
	org, des = des, org
	path, err = a.FindPath(org, des)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	a.Draw()
	//
	// ####################
	// #                  #
	// #                  #
	// #  B               #
	// #  *               #
	// #  *               #
	// #  *               #
	// #  *               #
	// #  *               #
	// #  *#########      #
	// #   *******        #
	// #          *       #
	// #           *      #
	// #            *     #
	// #             *    #
	// #              A   #
	// #                  #
	// #                  #
	// #                  #
	// ####################
	//
	return
}
```

## More Cases
```
################        ################
#              #        #              #
# A            #        # A            #
#     ##       #        #  *  ##       #
#              #        #   *          #
#            B #        #    ********B #
################        ################

----------------------------------------

################        ################
#              #        #              #
# A            #        # A****        #
####### ########        #######*########
#              #        #       *      #
#            B #        #        ****B #
################        ################

----------------------------------------

################        ################
#              #        #     **       #
# A    #       #        # A  * #*      #
#      #       #        #  **  # *     #
#      #       #        #      #  *    #
#      #     B #        #      #   **B #
################        ################

----------------------------------------

################
#       #      #
# A     #      #
#       #      #            no way!
#       #      #
#       #    B #
################

----------------------------------------

################
#              #
# A            #
#              #            no way!
#          #####
#          # B #
################ 
```
