package astar

import (
	"errors"
	"math"

	"github.com/shanghuiyang/a-star/scene"
)

var (
	// ErrNoWay ...
	ErrNoWay = errors.New("no way")
)

// AStar ...
type AStar struct {
	origin    *Point
	dest      *Point
	openlist  []*Point
	closelist []*Point
	path      []*Point
	scene     *scene.Scene
}

// New ...
func New(org, des *Point, s *scene.Scene) *AStar {
	a := &AStar{
		origin: org,
		dest:   des,
		scene:  s,
	}
	s.Set(org.X, org.Y, 'A')
	s.Set(des.X, des.Y, 'B')
	a.openlist = append(a.openlist, org)
	return a
}

// Run ...
func (a *AStar) Run() ([]*Point, error) {
	cur, err := a.minf()
	if err != nil {
		return nil, err
	}
	if err := a.removeFromOpenList(cur); err != nil {
		return nil, err
	}
	if cur.X == a.dest.X && cur.Y == a.dest.Y {
		a.genPath(cur)
		return a.path, nil
	}

	a.addToCloseList(cur)
	walkable := a.getWalkable(cur)
	for _, p := range walkable {
		a.addToOpenList(p)
	}
	return a.Run()
}

func (a *AStar) minf() (*Point, error) {
	if len(a.openlist) == 0 {
		return nil, ErrNoWay
	}
	min := a.openlist[0]
	for i := 1; i < len(a.openlist); i++ {
		if a.openlist[i].F < min.F {
			min = a.openlist[i]
		}
	}
	return min, nil
}

func (a *AStar) getWalkable(p *Point) []*Point {
	var around []*Point
	row, col := p.X, p.Y
	left := a.scene.Get(row, col-1) //s.scene[row][col-1]
	up := a.scene.Get(row-1, col)
	right := a.scene.Get(row, col+1)
	down := a.scene.Get(row+1, col)
	leftup := a.scene.Get(row-1, col-1)
	rightup := a.scene.Get(row-1, col+1)
	leftdown := a.scene.Get(row+1, col-1)
	rightdown := a.scene.Get(row+1, col+1)
	if (left == ' ') || (left == 'B') {
		around = append(around, &Point{row, col - 1, 0, 0, 0, p})
	}
	if (leftup == ' ') || (leftup == 'B') {
		around = append(around, &Point{row - 1, col - 1, 0, 0, 0, p})
	}
	if (up == ' ') || (up == 'B') {
		around = append(around, &Point{row - 1, col, 0, 0, 0, p})
	}
	if (rightup == ' ') || (rightup == 'B') {
		around = append(around, &Point{row - 1, col + 1, 0, 0, 0, p})
	}
	if (right == ' ') || (right == 'B') {
		around = append(around, &Point{row, col + 1, 0, 0, 0, p})
	}
	if (rightdown == ' ') || (rightdown == 'B') {
		around = append(around, &Point{row + 1, col + 1, 0, 0, 0, p})
	}
	if (down == ' ') || (down == 'B') {
		around = append(around, &Point{row + 1, col, 0, 0, 0, p})
	}
	if (leftdown == ' ') || (leftdown == 'B') {
		around = append(around, &Point{row + 1, col - 1, 0, 0, 0, p})
	}
	return around
}

func (a *AStar) addToOpenList(p *Point) {
	a.updateWeight(p)
	if a.checkExist(p, a.closelist) {
		return
	}
	if !a.checkExist(p, a.openlist) {
		a.openlist = append(a.openlist, p)
	} else {
		if a.openlist[a.findPoint(p, a.openlist)].F > p.F {
			a.openlist[a.findPoint(p, a.openlist)].Parent = p.Parent
		}
	}
}

// Update G, H, F of the point
func (a *AStar) updateWeight(p *Point) {
	if a.checkRelativePos(p) == 1 {
		p.G = p.Parent.G + 10
	} else {
		p.G = p.Parent.G + 14
	}
	absx := (int)(math.Abs((float64)(a.dest.X - p.X)))
	absy := (int)(math.Abs((float64)(a.dest.Y - p.Y)))
	p.H = (absx + absy) * 10
	p.F = p.G + p.H
}

func (a *AStar) removeFromOpenList(p *Point) error {
	index := a.findPoint(p, a.openlist)
	if index == -1 {
		return errors.New("not fount the point in the open_list")
	}
	a.openlist = append(a.openlist[:index], a.openlist[index+1:]...)
	return nil
}

func (a *AStar) addToCloseList(p *Point) {
	// removeFromOpenList(p)
	if a.scene.Get(p.X, p.Y) != 'A' {
		a.scene.Set(p.X, p.Y, '.')
	}
	a.closelist = append(a.closelist, p)
}

func (a *AStar) checkExist(p *Point, list []*Point) bool {
	for _, pt := range list {
		if p.X == pt.X && p.Y == pt.Y {
			return true
		}
	}
	return false
}

func (a *AStar) findPoint(p *Point, list []*Point) int {
	for idx, pt := range list {
		if p.X == pt.X && p.Y == pt.Y {
			return idx
		}
	}
	return -1
}

func (a *AStar) checkRelativePos(p *Point) int {
	parent := p.Parent
	hor := (int)(math.Abs((float64)(p.X - parent.X)))
	ver := (int)(math.Abs((float64)(p.Y - parent.Y)))
	return hor + ver
}

func (a *AStar) genPath(p *Point) {
	if a.scene.Get(p.X, p.Y) != 'A' && a.scene.Get(p.X, p.Y) != 'B' {
		a.scene.Set(p.X, p.Y, '*')
	}
	a.path = append([]*Point{p}, a.path...)
	if p.Parent != nil {
		a.genPath(p.Parent)
	}
}
