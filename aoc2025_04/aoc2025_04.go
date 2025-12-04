package main

import (
	"AoC2025/util"
	"fmt"
)

const example = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

const real = ``

func main() {
	part1(example)
	part2(example)
}

func prep(input string) util.Map {
	m := util.Map{}
	m.Hydrate(input)
	return m
}

func part1(input string) {
	m := prep(input)
	result := 0
	for coord, tile := range m.Walk() {
		if tile == '@' {
			cnt := 0
			for _, tile := range m.Iter8(coord) {
				if tile == '@' {
					cnt++
				}
			}
			if cnt < 4 {
				result += 1
			}
		}
	}
	fmt.Println(result)
}

func part2(input string) {
	m := prep(input)
	result := 0
	for {
		nextMap := m.Copy()
		intermResult := 0
		for coord, tile := range m.Walk() {
			if tile == '@' {
				cnt := 0
				for _, tile := range m.Iter8(coord) {
					if tile == '@' {
						cnt++
					}
				}
				if cnt < 4 {
					nextMap.Put(coord, '.')
					intermResult += 1
				}
			}
		}
		//nextMap.Print()

		m = nextMap
		result += intermResult
		//fmt.Println(intermResult)
		if intermResult == 0 {
			break
		}
	}

	fmt.Println(result)
}
