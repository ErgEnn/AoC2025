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
	row int
	col int
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
	row := m.rows[c.row]
	return row[c.col]
}

func (m *Map) Iter8(c Coord) iter.Seq2[Coord, rune] {
	return func(yield func(Coord, rune) bool) {
		if c.row-1 >= 0 {
			if c.col-1 >= 0 {
				newC := Coord{c.row - 1, c.col - 1}
				yield(newC, m.Get(newC))
			}
			{
				newC := Coord{c.row - 1, c.col}
				yield(newC, m.Get(newC))
			}
			if c.col+1 < m.width {
				newC := Coord{c.row - 1, c.col + 1}
				yield(newC, m.Get(newC))
			}
		}
		if c.col+1 < m.width {
			newC := Coord{c.row, c.col + 1}
			yield(newC, m.Get(newC))
		}
		if c.row+1 < m.height {
			if c.col+1 < m.width {
				newC := Coord{c.row + 1, c.col + 1}
				yield(newC, m.Get(newC))
			}
			{
				newC := Coord{c.row + 1, c.col}
				yield(newC, m.Get(newC))
			}
			if c.col-1 >= 0 {
				newC := Coord{c.row + 1, c.col - 1}
				yield(newC, m.Get(newC))
			}
		}
		if c.col-1 >= 0 {
			newC := Coord{c.row, c.col - 1}
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
	m.rows[c.row][c.col] = r
}

func (m *Map) Print() {
	for row := 0; row < m.height; row++ {
		for col := 0; col < m.width; col++ {
			fmt.Print(string(m.rows[row][col]))
		}
		fmt.Println()
	}
}
