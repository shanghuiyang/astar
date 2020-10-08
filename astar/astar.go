package astar

import (
	"errors"
	"fmt"
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
	openlist  List
	closelist List
	path      List
	scene     *scene.Scene
	baseScene *scene.Scene
}

// New ...
func New(s *scene.Scene) *AStar {
	return &AStar{
		baseScene: s,
	}
}

// FindPath ...
func (a *AStar) FindPath(org, des *Point) (List, error) {
	a.scene = a.baseScene.Copy()
	a.origin = org
	a.dest = des
	a.scene.Set(org.X, org.Y, 'A')
	a.scene.Set(des.X, des.Y, 'B')

	a.openlist.Clear()
	a.closelist.Clear()
	a.path.Clear()
	a.openlist.Append(org)
	return a.find()
}

// String ...
func (a *AStar) String() string {
	return a.scene.String()
}

// Draw ...
func (a *AStar) Draw() {
	fmt.Print(a)
}

func (a *AStar) find() (List, error) {
	cur, err := a.minf()
	if err != nil {
		return nil, err
	}
	a.openlist.Remove(cur)

	if cur.X == a.dest.X && cur.Y == a.dest.Y {
		a.genPath(cur)
		return a.path, nil
	}

	a.closelist.Append(cur)
	walkable := a.getWalkable(cur)
	for _, p := range walkable {
		a.updateWeight(p)
		if a.closelist.Find(p) >= 0 {
			continue
		}

		if idx := a.openlist.Find(p); idx >= 0 {
			if a.openlist[idx].F > p.F {
				a.openlist[idx].Parent = p.Parent
			}
			continue
		}
		a.openlist.Append(p)
	}
	return a.find()
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

func (a *AStar) getWalkable(p *Point) List {
	var around List
	row, col := p.X, p.Y
	left := a.scene.Get(row, col-1)
	up := a.scene.Get(row-1, col)
	right := a.scene.Get(row, col+1)
	down := a.scene.Get(row+1, col)
	leftup := a.scene.Get(row-1, col-1)
	rightup := a.scene.Get(row-1, col+1)
	leftdown := a.scene.Get(row+1, col-1)
	rightdown := a.scene.Get(row+1, col+1)
	if (left == ' ') || (left == 'B') {
		around.Append(&Point{row, col - 1, 0, 0, 0, p})
	}
	if (leftup == ' ') || (leftup == 'B') {
		around.Append(&Point{row - 1, col - 1, 0, 0, 0, p})
	}
	if (up == ' ') || (up == 'B') {
		around.Append(&Point{row - 1, col, 0, 0, 0, p})
	}
	if (rightup == ' ') || (rightup == 'B') {
		around.Append(&Point{row - 1, col + 1, 0, 0, 0, p})
	}
	if (right == ' ') || (right == 'B') {
		around.Append(&Point{row, col + 1, 0, 0, 0, p})
	}
	if (rightdown == ' ') || (rightdown == 'B') {
		around.Append(&Point{row + 1, col + 1, 0, 0, 0, p})
	}
	if (down == ' ') || (down == 'B') {
		around.Append(&Point{row + 1, col, 0, 0, 0, p})
	}
	if (leftdown == ' ') || (leftdown == 'B') {
		around.Append(&Point{row + 1, col - 1, 0, 0, 0, p})
	}
	return around
}

func (a *AStar) updateWeight(p *Point) {
	if a.checkPos(p) == 1 {
		p.G = p.Parent.G + 10
	} else {
		p.G = p.Parent.G + 14
	}
	x := (int)(math.Abs((float64)(a.dest.X - p.X)))
	y := (int)(math.Abs((float64)(a.dest.Y - p.Y)))
	p.H = (x + y) * 10
	p.F = p.G + p.H
}

func (a *AStar) checkPos(p *Point) int {
	parent := p.Parent
	h := (int)(math.Abs((float64)(p.X - parent.X)))
	v := (int)(math.Abs((float64)(p.Y - parent.Y)))
	return h + v
}

func (a *AStar) genPath(p *Point) {
	if a.scene.Get(p.X, p.Y) != 'A' && a.scene.Get(p.X, p.Y) != 'B' {
		a.scene.Set(p.X, p.Y, '*')
	}
	a.path.Front(p)
	if p.Parent != nil {
		a.genPath(p.Parent)
	}
}
