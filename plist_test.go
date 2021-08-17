package astar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPList(t *testing.T) {
	l := NewPList(10)
	assert.NotNil(t, l)
	assert.Equal(t, 10, len(l))
}

func TestAppend(t *testing.T) {
	p1 := &Point{X: 1, Y: 1}
	p2 := &Point{X: 2, Y: 2}

	l := PList{p1}
	l.Append(p2)
	assert.Equal(t, 2, len(l))
	assert.Equal(t, PList{p1, p2}, l)
}

func TestFront(t *testing.T) {
	p1 := &Point{X: 1, Y: 1}
	p2 := &Point{X: 2, Y: 2}
	l := PList{p1}

	l.Front(p2)
	assert.Equal(t, 2, len(l))
	assert.Equal(t, PList{p2, p1}, l)
}

func TestRemove(t *testing.T) {
	p1 := &Point{X: 1, Y: 1}
	p2 := &Point{X: 2, Y: 2}
	l := PList{p1, p2}

	l.Remove(p2)
	assert.Equal(t, 1, len(l))
	assert.Equal(t, PList{p1}, l)
}

func TestClear(t *testing.T) {
	p1 := &Point{X: 1, Y: 1}
	p2 := &Point{X: 2, Y: 2}
	l := PList{p1, p2}

	l.Clear()
	assert.Equal(t, 0, len(l))
}

func TestFind(t *testing.T) {
	p1 := &Point{X: 1, Y: 1}
	p2 := &Point{X: 2, Y: 2}
	p3 := &Point{X: 3, Y: 3}
	l := PList{p1, p2}

	idx := l.Find(p1)
	assert.Equal(t, 0, idx)

	idx = l.Find(p2)
	assert.Equal(t, 1, idx)

	idx = l.Find(p3)
	assert.Equal(t, -1, idx)
}
