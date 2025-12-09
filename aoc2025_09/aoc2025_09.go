package main

import (
	"AoC2025/util"
	"fmt"
	"maps"
	"math"
	"slices"
	"sort"
	"strings"
)

const example = `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`

func main() {
	//real := util.ReadInput()

	part1(example)
	part2(example)
}

func prep(input string) []string {
	return strings.Split(input, "\n")
}

func part1(input string) {
	lines := prep(input)
	max := 0
	for _, line := range lines {
		x1, y1 := util.Split2Func(line, func(s string) int {
			return util.ToInt(s)
		})
		for _, line := range lines {
			x2, y2 := util.Split2Func(line, func(s string) int { return util.ToInt(s) })
			max = int(math.Max(math.Abs(float64(x1-x2)+1)*math.Abs(float64(y1-y2)+1), float64(max)))
		}
	}
	fmt.Println(max)
}

func normalize(m *map[int]int) int {
	keys := slices.Collect(maps.Keys(*m))
	sort.Ints(keys)
	for i, k := range keys {
		(*m)[k] = i * 2
	}
	return (*m)[keys[len(keys)-1]]
}

func inverseMap(m map[int]int) map[int]int {
	temp := make(map[int]int)
	for k, v := range m {
		temp[v] = k
	}
	return temp
}

func part2(input string) {
	lines := prep(input)

	vertices := make([]util.Coord, len(lines))

	xNormalized := make(map[int]int)
	yNormalized := make(map[int]int)
	for i, line := range lines {
		x, y := util.Split2Func(line, func(s string) int {
			return util.ToInt(s)
		})
		vertices[i] = util.Coord{y, x}
		xNormalized[x] = 0
		yNormalized[y] = 0
	}
	maxX := normalize(&xNormalized)
	maxY := normalize(&yNormalized)
	verticesNormalized := make([]util.Coord, len(lines))
	for i, coord := range vertices {
		verticesNormalized[i] = util.Coord{yNormalized[coord.Row], xNormalized[coord.Col]}
	}
	m := util.Map{}
	m.Empty(maxX+1, maxY+1, '.')
	prev := verticesNormalized[len(verticesNormalized)-1]
	for _, coord := range verticesNormalized {
		m.Put(coord, '#')
		for areaCoord := range coord.Area(prev) {
			if areaCoord == prev || areaCoord == coord {
				continue
			}
			m.Put(areaCoord, 'X')
		}
		prev = coord
	}
	invX := inverseMap(xNormalized)
	invY := inverseMap(yNormalized)
	max := 0
	var maxPair util.CoordPair
	for coordPair := range util.PairwiseCombinations(verticesNormalized) {
		area := int((math.Abs(float64(invX[coordPair.A.Col]-invX[coordPair.B.Col])) + 1) * (math.Abs(float64(invY[coordPair.A.Row]-invY[coordPair.B.Row])) + 1))
		for coord := range coordPair.A.InnerArea(coordPair.B) {
			tile := m.Get(coord)
			if tile != '.' {
				area = 0
				break
			}
		}
		if area > max {
			max = area
			maxPair = coordPair
		}
	}
	for coord := range maxPair.A.Area(maxPair.B) {
		tile := m.Get(coord)
		if tile == '.' {
			m.Put(coord, 'O')
		}

	}
	m.Print()
	fmt.Printf("%d,%d\t%d,%d\n", invX[maxPair.A.Col], invY[maxPair.A.Row], invX[maxPair.B.Col], invY[maxPair.B.Row])
	fmt.Println(max)

}
