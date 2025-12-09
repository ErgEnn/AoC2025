package util

import (
	"fmt"
	"iter"
	"strings"
)

type Map struct {
	rows   [][]rune
	height int
	width  int
}

type Coord struct {
	Row int
	Col int
}

func (c *Coord) String() string {
	return fmt.Sprintf("%d,%d", c.Col, c.Row)
}

func (a *Coord) Area(b Coord) iter.Seq[Coord] {
	r1, r2 := a.Row, b.Row
	if r1 > r2 {
		r1, r2 = r2, r1
	}

	c1, c2 := a.Col, b.Col
	if c1 > c2 {
		c1, c2 = c2, c1
	}

	return func(yield func(Coord) bool) {
		for r := r1; r <= r2; r++ {
			for c := c1; c <= c2; c++ {
				if !yield(Coord{Row: r, Col: c}) {
					return
				}
			}
		}
	}
}

func (a *Coord) InnerArea(b Coord) iter.Seq[Coord] {
	r1, r2 := a.Row, b.Row
	if r1 > r2 {
		r1, r2 = r2, r1
	}

	c1, c2 := a.Col, b.Col
	if c1 > c2 {
		c1, c2 = c2, c1
	}

	return func(yield func(Coord) bool) {
		r1i, r2i := r1+1, r2-1
		c1i, c2i := c1+1, c2-1

		if r1i > r2i || c1i > c2i {
			return
		}

		for r := r1i; r <= r2i; r++ {
			for c := c1i; c <= c2i; c++ {
				if !yield(Coord{Row: r, Col: c}) {
					return
				}
			}
		}
	}
}

func (m *Map) Empty(width int, height int, char rune) {
	m.rows = make([][]rune, m.height)
	m.width = width
	m.height = height
	for row := 0; row < height; row++ {
		m.rows = append(m.rows, make([]rune, width))
		for col := 0; col < width; col++ {
			m.rows[row][col] = char
		}
	}
}

func (m *Map) Hydrate(input string) {
	temp := strings.Split(input, "\n")
	m.rows = make([][]rune, len(temp))
	for row := 0; row < len(temp); row++ {
		m.rows[row] = make([]rune, len(temp[row]))
		for col := 0; col < len(temp[row]); col++ {
			m.rows[row][col] = rune(temp[row][col])
		}
	}
	m.height = len(m.rows)
	m.width = len(m.rows[0])
}

func (m *Map) Walk() iter.Seq2[Coord, rune] {
	return func(yield func(Coord, rune) bool) {
		for rIdx, row := range m.rows {
			for cIdx, char := range row {
				if (!yield(Coord{rIdx, cIdx}, char)) {
					return
				}
			}
		}
	}
}

func (m *Map) Get(c Coord) rune {
	row := m.rows[c.Row]
	return row[c.Col]
}

func (m *Map) Iter8(c Coord) iter.Seq2[Coord, rune] {
	return func(yield func(Coord, rune) bool) {
		if c.Row-1 >= 0 {
			if c.Col-1 >= 0 {
				newC := Coord{c.Row - 1, c.Col - 1}
				yield(newC, m.Get(newC))
			}
			{
				newC := Coord{c.Row - 1, c.Col}
				yield(newC, m.Get(newC))
			}
			if c.Col+1 < m.width {
				newC := Coord{c.Row - 1, c.Col + 1}
				yield(newC, m.Get(newC))
			}
		}
		if c.Col+1 < m.width {
			newC := Coord{c.Row, c.Col + 1}
			yield(newC, m.Get(newC))
		}
		if c.Row+1 < m.height {
			if c.Col+1 < m.width {
				newC := Coord{c.Row + 1, c.Col + 1}
				yield(newC, m.Get(newC))
			}
			{
				newC := Coord{c.Row + 1, c.Col}
				yield(newC, m.Get(newC))
			}
			if c.Col-1 >= 0 {
				newC := Coord{c.Row + 1, c.Col - 1}
				yield(newC, m.Get(newC))
			}
		}
		if c.Col-1 >= 0 {
			newC := Coord{c.Row, c.Col - 1}
			yield(newC, m.Get(newC))
		}
	}
}

func (m *Map) Copy() Map {
	newMap := Map{
		rows:   make([][]rune, m.height),
		width:  m.width,
		height: m.height,
	}

	for row := 0; row < m.height; row++ {
		newMap.rows[row] = make([]rune, m.width)
		copy(newMap.rows[row], m.rows[row])
	}

	return newMap
}

func (m *Map) Put(c Coord, r rune) {
	m.rows[c.Row][c.Col] = r
}

func (m *Map) Print() {
	for row := 0; row < m.height; row++ {
		for col := 0; col < m.width; col++ {
			fmt.Print(string(m.rows[row][col]))
		}
		fmt.Println()
	}
}
