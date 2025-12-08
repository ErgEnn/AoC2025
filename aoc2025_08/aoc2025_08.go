package main

import (
	"AoC2025/util"
	"fmt"
	"maps"
	"math"
	"slices"
	"strings"
)

const example = `162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`

func main() {
	real := util.ReadInput()

	//part1(real, 1000)
	part2(real)
}

func prep(input string) []Coord3D {
	lines := strings.Split(input, "\n")
	junctionBoxes := make([]Coord3D, 0)
	for _, line := range lines {
		x, y, z := util.Split3(line, ",")
		junctionBoxes = append(junctionBoxes, Coord3D{
			util.ToInt(x),
			util.ToInt(y),
			util.ToInt(z),
		})
	}
	return junctionBoxes
}

type Coord3D struct {
	x int
	y int
	z int
}

func (c Coord3D) String() string {
	return fmt.Sprintf("%d,%d,%d", c.x, c.y, c.z)
}

func (a *Coord3D) DistTo(b Coord3D) int {
	return int(math.Pow(float64(a.x-b.x), 2) + math.Pow(float64(a.y-b.y), 2) + math.Pow(float64(a.z-b.z), 2))
}

type Coord3DPair struct {
	a Coord3D
	b Coord3D
}

func (p Coord3DPair) String() string {
	return fmt.Sprintf("%d,%d,%d <-> %d,%d,%d", p.a.x, p.a.y, p.a.z, p.b.x, p.b.y, p.b.z)
}

func part1(input string, noOfConnections int) {
	junctionBoxes := prep(input)

	distances := make(map[Coord3DPair]int)
	for _, a := range junctionBoxes {
		for _, b := range junctionBoxes {
			if a == b {
				continue
			}
			_, existsab := distances[Coord3DPair{a, b}]
			if existsab {
				continue
			}
			_, existsba := distances[Coord3DPair{b, a}]
			if existsba {
				continue
			}
			distances[Coord3DPair{a, b}] = a.DistTo(b)
		}
	}
	pairs := slices.SortedFunc(maps.Keys(distances), func(a, b Coord3DPair) int {
		return distances[a] - distances[b]
	})
	circuits := make(map[Coord3D]int)
	circuitIdxGen := 0
	for i := 0; i < noOfConnections; i++ {
		connection := pairs[i]
		aIdx, a := circuits[connection.a]
		bIdx, b := circuits[connection.b]
		if a && b && aIdx != bIdx {
			for k, v := range circuits {
				if v == bIdx {
					circuits[k] = aIdx
				}
			}
		}
		if a {
			circuits[connection.b] = aIdx
			continue
		}

		if b {
			circuits[connection.a] = bIdx
			continue
		}
		circuits[connection.a] = circuitIdxGen
		circuits[connection.b] = circuitIdxGen
		circuitIdxGen++
	}
	circuitSizes := make(map[int]int)
	for _, circuitIdx := range circuits {
		circuitSizes[circuitIdx] = circuitSizes[circuitIdx] + 1
	}
	orderedCircuitSizes := slices.SortedFunc(maps.Values(circuitSizes), func(i int, j int) int {
		return j - i
	})
	fmt.Println(circuitSizes)
	fmt.Println(orderedCircuitSizes[0] * orderedCircuitSizes[1] * orderedCircuitSizes[2])
}

func part2(input string) {
	junctionBoxes := prep(input)

	distances := make(map[Coord3DPair]int)
	for _, a := range junctionBoxes {
		for _, b := range junctionBoxes {
			if a == b {
				continue
			}
			_, existsab := distances[Coord3DPair{a, b}]
			if existsab {
				continue
			}
			_, existsba := distances[Coord3DPair{b, a}]
			if existsba {
				continue
			}
			distances[Coord3DPair{a, b}] = a.DistTo(b)
		}
	}
	pairs := slices.SortedFunc(maps.Keys(distances), func(a, b Coord3DPair) int {
		return distances[a] - distances[b]
	})
	circuits := make(map[Coord3D]int)
	circuitIdxGen := 0
	for i := 0; true; i++ {
		if len(circuits) == len(junctionBoxes) && len(util.Distinct(maps.Values(circuits))) == 1 {
			connection := pairs[i-1]
			fmt.Println(connection.a.String())
			fmt.Println(connection.b.String())
			fmt.Println(connection.a.x * connection.b.x)
			break
		}
		connection := pairs[i]
		aIdx, a := circuits[connection.a]
		bIdx, b := circuits[connection.b]
		if a && b && aIdx != bIdx {
			for k, v := range circuits {
				if v == bIdx {
					circuits[k] = aIdx
				}
			}
		}
		if a {
			circuits[connection.b] = aIdx
			continue
		}

		if b {
			circuits[connection.a] = bIdx
			continue
		}
		circuits[connection.a] = circuitIdxGen
		circuits[connection.b] = circuitIdxGen
		circuitIdxGen++
	}
}
