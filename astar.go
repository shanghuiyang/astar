package astar

import (
	"errors"
	"math"

	"github.com/shanghuiyang/astar/tilemap"
)

var (
	// ErrNoWay ...
	ErrNoWay = errors.New("no way")
)

// AStar ...
type AStar struct {
	origin      *Point
	dest        *Point
	openlist    PList
	closedlist  PList
	path        PList
	tilemap     *tilemap.Tilemap
	baseTilemap *tilemap.Tilemap
}

// New ...
func New(m *tilemap.Tilemap) *AStar {
	return &AStar{
		baseTilemap: m,
	}
}

// FindPath ...
func (a *AStar) FindPath(org, des *Point) (PList, error) {
	a.tilemap = a.baseTilemap.Clone()
	a.origin = org
	a.dest = des
	a.tilemap.Set(org.X, org.Y, 'A')
	a.tilemap.Set(des.X, des.Y, 'B')

	a.openlist.Clear()
	a.closedlist.Clear()
	a.path.Clear()
	a.openlist.Append(org)
	return a.find()
}

// String ...
func (a *AStar) String() string {
	return a.tilemap.String()
}

// Draw ...
func (a *AStar) Draw() {
	a.tilemap.Draw()
}

func (a *AStar) find() (PList, error) {
	cur, err := a.minf()
	if err != nil {
		return nil, err
	}
	a.openlist.Remove(cur)

	if cur.X == a.dest.X && cur.Y == a.dest.Y {
		a.genPath(cur)
		return a.path, nil
	}

	a.closedlist.Append(cur)
	walkable := a.getWalkable(cur)
	for _, p := range walkable {
		a.updateWeight(p)
		if a.closedlist.Find(p) >= 0 {
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

func (a *AStar) getWalkable(p *Point) PList {
	var around PList
	x, y := p.X, p.Y
	left := a.tilemap.Get(x, y-1)
	up := a.tilemap.Get(x-1, y)
	right := a.tilemap.Get(x, y+1)
	down := a.tilemap.Get(x+1, y)
	leftup := a.tilemap.Get(x-1, y-1)
	rightup := a.tilemap.Get(x-1, y+1)
	leftdown := a.tilemap.Get(x+1, y-1)
	rightdown := a.tilemap.Get(x+1, y+1)
	if (left == ' ') || (left == 'B') {
		around.Append(&Point{x, y - 1, 0, 0, 0, p})
	}
	if (leftup == ' ') || (leftup == 'B') {
		around.Append(&Point{x - 1, y - 1, 0, 0, 0, p})
	}
	if (up == ' ') || (up == 'B') {
		around.Append(&Point{x - 1, y, 0, 0, 0, p})
	}
	if (rightup == ' ') || (rightup == 'B') {
		around.Append(&Point{x - 1, y + 1, 0, 0, 0, p})
	}
	if (right == ' ') || (right == 'B') {
		around.Append(&Point{x, y + 1, 0, 0, 0, p})
	}
	if (rightdown == ' ') || (rightdown == 'B') {
		around.Append(&Point{x + 1, y + 1, 0, 0, 0, p})
	}
	if (down == ' ') || (down == 'B') {
		around.Append(&Point{x + 1, y, 0, 0, 0, p})
	}
	if (leftdown == ' ') || (leftdown == 'B') {
		around.Append(&Point{x + 1, y - 1, 0, 0, 0, p})
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
	if a.tilemap.Get(p.X, p.Y) != 'A' && a.tilemap.Get(p.X, p.Y) != 'B' {
		a.tilemap.Set(p.X, p.Y, '.')
	}
	a.path.Front(p)
	if p.Parent != nil {
		a.genPath(p.Parent)
	}
}
