// the coordinate system of a tilemap
//
//  0 +-----------> y
//    |
//    |
//    |
//    v
//    x
//
package tilemap

import (
	"fmt"
)

// Tilemap ...
type Tilemap struct {
	r    int
	c    int
	data [][]byte
}

// New ...
func New(r, c int) *Tilemap {
	data := make([][]byte, r)
	for x := 0; x < r; x++ {
		data[x] = make([]byte, c)
		for y := 0; y < c; y++ {
			d := byte(' ')
			if x == 0 || x == r-1 || y == 0 || y == c-1 {
				d = '#'
			}
			data[x][y] = d
		}
	}
	return &Tilemap{
		r:    r,
		c:    c,
		data: data,
	}
}

// BuildFromStr ...
func BuildFromStr(s string) *Tilemap {
	m := &Tilemap{}
	cols := []byte{}
	for _, c := range s {
		if c == '\n' {
			if len(cols) == 0 {
				continue
			}
			m.r++
			m.data = append(m.data, cols)
			m.c = len(cols)
			cols = []byte{}
			continue
		}
		cols = append(cols, byte(c))
	}
	return m
}

// BuildFromGeoJSON ...
func BuildFromGeoJSON(file string) *Tilemap {
	return nil
}

// Get ...
func (m *Tilemap) Get(x, y int) byte {
	return m.data[x][y]
}

// Set ...
func (m *Tilemap) Set(x, y int, data byte) {
	m.data[x][y] = data
}

// GetRow ...
func (m *Tilemap) GetRow() int {
	return m.r
}

// GetCol ...
func (m *Tilemap) GetCol() int {
	return m.c
}

// SetWall ...
func (m *Tilemap) SetWall(x, y int) {
	if x < 0 || x >= m.r || y < 0 || y >= m.c {
		return
	}
	m.data[x][y] = '#'
}

// String ...
func (m *Tilemap) String() string {
	var s string
	for x := 0; x < m.r; x++ {
		for y := 0; y < m.c; y++ {
			s += fmt.Sprintf("%c", m.data[x][y])
		}
		s += "\n"
	}
	return s
}

// Draw ...
func (m *Tilemap) Draw() {
	fmt.Print(m)
}

// Clone ...
func (m *Tilemap) Clone() *Tilemap {
	data := make([][]byte, m.r)
	for x := 0; x < m.r; x++ {
		data[x] = make([]byte, m.c)
		copy(data[x], m.data[x])
	}
	return &Tilemap{
		r:    m.r,
		c:    m.c,
		data: data,
	}
}
