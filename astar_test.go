package astar

import (
	"testing"

	"github.com/shanghuiyang/astar/tilemap"
	"github.com/stretchr/testify/assert"
)

var (
	tilemaps = []string{
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
#  .           #
#   .          #
#    ........B #
################
`, `
################
#              #
# A            #
#  .  ##       #
#   .          #
#    ........B #
################
`, `
################
#              #
# A....        #
#######.########
#       .      #
#        ....B #
################
`, `
################
#     ..       #
# A  . #.      #
#  ..  # .     #
#      #  .    #
#      #   ..B #
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

	for i, str := range tilemaps {
		m := tilemap.BuildFromStr(str)
		a := New(m)
		_, err := a.FindPath(org, des)
		assert.NoError(t, err)

		actual := "\n" + a.String()
		assert.Equal(t, expecteds[i], actual)
	}

	for _, str := range noWayScens {
		m := tilemap.BuildFromStr(str)
		a := New(m)
		path, err := a.FindPath(org, des)
		assert.Error(t, err)
		assert.Equal(t, ErrNoWay, err)
		assert.Nil(t, path)
	}

}
