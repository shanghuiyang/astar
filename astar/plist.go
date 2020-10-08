package astar

import (
	"fmt"
)

// PList ...
type PList []*Point

// NewPList ...
func NewPList(size int) PList {
	l := make(PList, size)
	return l
}

// Append ...
func (l *PList) Append(p *Point) {
	*l = append(*l, p)
}

// Front ...
func (l *PList) Front(p *Point) {
	*l = append(PList{p}, *l...)
}

// Remove ...
func (l *PList) Remove(p *Point) {
	new := PList{}
	for _, pt := range *l {
		if pt.X == p.X && pt.Y == p.Y {
			continue
		}
		new.Append(pt)
	}
	*l = new
}

// Clear ...
func (l *PList) Clear() {
	*l = nil
}

// Find ...
func (l *PList) Find(p *Point) int {
	for i, pt := range *l {
		if pt.X == p.X && pt.Y == p.Y {
			return i
		}
	}
	return -1
}

// String ...
func (l PList) String() string {
	var s string
	for _, p := range l {
		s += fmt.Sprintf("(%v, %v) ", p.X, p.Y)
	}
	return s
}
