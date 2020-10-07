package scene

import (
	"fmt"
)

// Scene ...
type Scene struct {
	r int
	c int
	s [][]byte
}

// New ...
func New(r, c int) *Scene {
	s := make([][]byte, r)
	for i := 0; i < r; i++ {
		s[i] = make([]byte, c)
		for j := 0; j < c; j++ {
			if i == 0 || i == r-1 || j == 0 || j == c-1 {
				s[i][j] = '#'
			} else {
				s[i][j] = ' '
			}
		}
	}
	return &Scene{
		r: r,
		c: c,
		s: s,
	}
}

// BuildFromGeoJSON ...
func BuildFromGeoJSON(file string) *Scene {
	return nil
}

// Get ...
func (s *Scene) Get(r, c int) byte {
	return s.s[r][c]
}

// Set ...
func (s *Scene) Set(r, c int, ctx byte) {
	s.s[r][c] = ctx
}

// SetWall ...
func (s *Scene) SetWall(r, c int) {
	if r < 0 || r > s.r || c < 0 || c > s.c {
		return
	}
	s.s[r][c] = '#'
}

// String ...
func (s *Scene) String() string {
	var str string
	for i := 0; i < s.r; i++ {
		for j := 0; j < s.c; j++ {
			str += fmt.Sprintf("%c", s.s[i][j])
		}
		str += fmt.Sprintf("\n")
	}
	return str
}
