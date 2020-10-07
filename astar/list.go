package astar

// List ...
type List []*Point

// NewList ...
func NewList(size int) List {
	l := make(List, size)
	return l
}

// Append ...
func (l List) Append(p *Point) {
	l = append(l, p)
}

// Front ...
func (l List) Front(p *Point) {
	l = append(List{p}, l...)
}

// Remove ...
func (l List) Remove(p *Point) {
	new := List{}
	for _, pt := range l {
		if pt.X == p.X && pt.Y == p.Y {
			continue
		}
		new.Append(pt)
	}
	l = new
}
