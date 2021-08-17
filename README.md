<img src="a-star.gif" width=20% height=20% />

# A-Star
[![Build Status](https://travis-ci.org/shanghuiyang/astar.svg?branch=main)](https://travis-ci.org/shanghuiyang/astar)

A-Star algorithm implemented with Go.

## Usage
see the [example/main.go](example/main.go) for complete usage.
```go
package main

import (
	"fmt"

	"github.com/shanghuiyang/astar"
	"github.com/shanghuiyang/astar/tilemap"
)

func main() {
	r, c := 20, 20
	m := tilemap.New(r, c)
	for x := 4; x < 13; x++ {
		m.SetWall(9, x)
	}

	org := &astar.Point{X: 3, Y: 3}		// origin
	des := &astar.Point{X: 15, Y: 15}	// destination
	a := astar.New(m)
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
	// #   .              #
	// #    .             #
	// #     .            #
	// #      .           #
	// #       .....      #
	// #   #########.     #
	// #             .    #
	// #              .   #
	// #              .   #
	// #              .   #
	// #              .   #
	// #              B   #
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
#     ##       #        #  .  ##       #
#              #        #   .          #
#            B #        #    ........B #
################        ################

----------------------------------------

################        ################
#              #        #              #
# A            #        # A....        #
####### ########        #######.########
#              #        #       .      #
#            B #        #        ....B #
################        ################

----------------------------------------

################        ################
#              #        #     ..       #
# A    #       #        # A  . #.      #
#      #       #        #  ..  # .     #
#      #       #        #      #  .    #
#      #     B #        #      #   ..B #
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
