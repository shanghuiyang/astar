package astar

import (
	"testing"
	// "fmt"

	"github.com/shanghuiyang/a-star/scene"
	"github.com/stretchr/testify/assert"
)

var (
	scenes = []string{
		`
################
#              #
# A            #
#              #
#              #
#            B #
################
`, `
################
#              #
# A            #
#     ##       #
#              #
#            B #
################
`, `
################
#              #
# A            #
####### ########
#              #
#            B #
################
`, `
################
#              #
# A    #       #
#      #       #
#      #       #
#      #     B #
################
`,
	}

	expecteds = []string{
		`
################
#              #
# A            #
#  *           #
#   *          #
#    ********B #
################
`, `
################
#              #
# A            #
#  *  ##       #
#   *          #
#    ********B #
################
`, `
################
#              #
# A****        #
#######*########
#       *      #
#        ****B #
################
`, `
################
#     **       #
# A  * #*      #
#  **  # *     #
#      #  *    #
#      #   **B #
################
`,
	}
)

var (
	noWayScens = []string{
		`
################
#              #
# A            #
#              #
################
#            B #
################
`, `
################
#       #      #
# A     #      #
#       #      #
#       #      #
#       #    B #
################
`, `
################
#         ##   #
# A      ##    #
#       ##     #
#      ##      #
#     ##     B #
################
`, `
################
#   #          #
# A #          #
#####          #
#              #
#            B #
################
`, `
################
#              #
# A            #
#              #
#          #####
#          # B #
################
`,
	}
)

func TestAStar(t *testing.T) {
	org := &Point{X: 2, Y: 2}
	des := &Point{X: 5, Y: 13}

	for i, str := range scenes {
		s := scene.BuildFromStr(str)
		a := New(org, des, s)
		_, err := a.Run()
		assert.NoError(t, err)

		actual := "\n" + s.String()
		assert.Equal(t, expecteds[i], actual)
	}

	for _, str := range noWayScens {
		s := scene.BuildFromStr(str)
		a := New(org, des, s)
		path, err := a.Run()
		assert.Error(t, err)
		assert.Equal(t, ErrNoWay, err)
		assert.Nil(t, path)
	}

}
