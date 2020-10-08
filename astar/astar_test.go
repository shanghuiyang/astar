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
		a := New(s)
		_, err := a.FindPath(org, des)
		assert.NoError(t, err)

		actual := "\n" + a.String()
		assert.Equal(t, expecteds[i], actual)
	}

	for _, str := range noWayScens {
		s := scene.BuildFromStr(str)
		a := New(s)
		path, err := a.FindPath(org, des)
		assert.Error(t, err)
		assert.Equal(t, ErrNoWay, err)
		assert.Nil(t, path)
	}

}
