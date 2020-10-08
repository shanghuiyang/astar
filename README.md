# A-Star
A-Star algorithm implemented with Go.

## Usage
build a scene map first.
```go
    // build a scene map directly
    r, c := 20, 20
	s := scene.New(r, c)
	for i := 4; i < 13; i++ {
		s.SetWall(9, i)
    }
    
    // or, build from string
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
```

find the path using a-star
```go
	org := &astar.Point{X: 3, Y: 3}     // origin
	des := &astar.Point{X: 15, Y: 15}   // destination

	a := astar.New(org, des, s)
	path, err := a.Run()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

    s.Draw()
    fmt.Println("path: ", path)
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
```

see the [main.go](/main.go) for complete usage.

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

################        ################
#       #      #        #       #      #
# A     #      #        # A     #      #
#       #      #        #       #      # no way!
#       #      #        #       #      #
#       #    B #        #       #    B #
################        ################

----------------------------------------

################        ################
#              #        #              #
# A            #        # A            #
#              #        #              # no way!
#          #####        #          #####
#          # B #        #          # B #
################        ################
```
