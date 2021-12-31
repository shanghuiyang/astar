<img src="a-star.gif" width=20% height=20% />

# A-Star
[![CI](https://github.com/shanghuiyang/astar/actions/workflows/ci.yml/badge.svg)](https://github.com/shanghuiyang/astar/actions/workflows/ci.yml)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/shanghuiyang/astar/blob/master/LICENSE)

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
```

output,
```
####################
#              .B  #
#             .    #
#   #########.     #
#    ........      #
#   .    #######   #
#  .               #
# A                #
#                  #
####################
path: (7, 2) (6, 3) (5, 4) (4, 5) (4, 6) (4, 7) (4, 8) (4, 9) (4, 10) (4, 11) (4, 12) (3, 13) (2, 14) (1, 15) (1, 16)
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
