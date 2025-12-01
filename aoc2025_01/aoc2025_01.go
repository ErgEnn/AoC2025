package main

import (
	"fmt"
	"strconv"
	"strings"
)

const example = `
L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
`

const real = `
`

func main() {
	part1(real)
	part2(real)
}

func prep(input string) []string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	return lines
}

func part1(input string) {
	lines := prep(input)
	current := 50
	answer := 0
	for _, l := range lines {
		offset, _ := strconv.Atoi(l[1:])
		switch l[0] {
		case 'L':
			current = (current - offset + 100) % 100
		case 'R':
			current = (current + offset + 100) % 100
		}
		if current == 0 {
			answer++
		}
	}
	fmt.Println(answer)
}

func part2(input string) {
	lines := prep(input)
	current := 50
	answer := 0
	for _, l := range lines {
		offset, _ := strconv.Atoi(l[1:])
		switch l[0] {
		case 'L':
			current = current - offset
		case 'R':
			current = current + offset
		}
		for {
			if current == 0 {
				break
			}
			if current < 0 {
				answer++
				current += 100
				continue
			}
			if current >= 100 {
				answer++
				current -= 100
				continue
			}
			break
		}
	}
	fmt.Println(answer)
}
