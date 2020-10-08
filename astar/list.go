package astar

import (
	"fmt"
)

// List ...
type List []*Point

// NewList ...
func NewList(size int) List {
	l := make(List, size)
	return l
}

// Append ...
func (l *List) Append(p *Point) {
	*l = append(*l, p)
}

// Front ...
func (l *List) Front(p *Point) {
	*l = append(List{p}, *l...)
}

// Remove ...
func (l *List) Remove(p *Point) {
	new := List{}
	for _, pt := range *l {
		if pt.X == p.X && pt.Y == p.Y {
			continue
		}
		new.Append(pt)
	}
	*l = new
}

// Clear ...
func (l *List) Clear() {
	*l = nil
}

// Find ...
func (l *List) Find(p *Point) int {
	for i, pt := range *l {
		if pt.X == p.X && pt.Y == p.Y {
			return i
		}
	}
	return -1
}

// String ...
func (l List) String() string {
	var s string
	for _, p := range l {
		s += fmt.Sprintf("(%v, %v) ", p.X, p.Y)
	}
	return s
}
