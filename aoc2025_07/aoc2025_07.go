package main

import (
	"AoC2025/util"
	"fmt"
	"strings"
)

const example = `.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............`

const real = ``

func main() {
	part1(example)
	part2(example)
}

func part1(input string) {
	lines := strings.Split(input, "\n")
	splits := make(map[int]bool)
	cnt := 0
	for _, line := range lines {
		for i := 0; i < len(line); i++ {
			if line[i] == '^' {
				if splits[i] {
					splits[i] = false
					splits[i-1] = true
					splits[i+1] = true
					cnt++
					continue
				}
			}
			if line[i] == 'S' {
				splits[i] = true
			}
		}
	}
	fmt.Println(cnt)
}

func part2(input string) {
	lines := strings.Split(input, "\n")
	splits := make(map[int]int)
	for _, line := range lines {
		for i := 0; i < len(line); i++ {
			if line[i] == '^' {
				if splits[i] > 0 {
					splits[i-1] = splits[i-1] + splits[i]
					splits[i+1] = splits[i+1] + splits[i]
					splits[i] = 0
				}
			}
			if line[i] == 'S' {
				splits[i] = 1
			}
		}
	}
	fmt.Println(util.Sum(splits))
}
