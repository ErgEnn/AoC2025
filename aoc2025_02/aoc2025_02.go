package main

import (
	"fmt"
	"strconv"
	"strings"
)

const example = `
11-22,95-115,998-1012,1188511880-1188511890,222220-222224,
1698522-1698528,446443-446449,38593856-38593862,565653-565659,
824824821-824824827,2121212118-2121212124
`

const real = `
`

func main() {
	part1(example)
	part2(example)
}

func prep(input string) []string {
	return strings.Split(input, ",")
}

func part1(input string) {
	id_pairs := prep(input)
	sum := 0
	for _, id_pair := range id_pairs {
		id_pair = strings.TrimSpace(id_pair)
		limits := strings.Split(id_pair, "-")
		limitStart, _ := strconv.Atoi(limits[0])
		limitEnd, _ := strconv.Atoi(limits[1])
		for i := limitStart; i <= limitEnd; i++ {
			if isTwice(strconv.Itoa(i)) {
				sum += i
			}
		}

	}
	fmt.Println(sum)
}

func isTwice(part string) bool {
	if len(part)%2 == 1 {
		return false
	}
	for i := 0; i < len(part)/2; i++ {
		if part[i] != part[len(part)/2+i] {
			return false
		}
	}
	return true
}

func part2(input string) {
	id_pairs := prep(input)
	sum := 0
	for _, id_pair := range id_pairs {
		id_pair = strings.TrimSpace(id_pair)
		limits := strings.Split(id_pair, "-")
		limitStart, _ := strconv.Atoi(limits[0])
		limitEnd, _ := strconv.Atoi(limits[1])
		for i := limitStart; i <= limitEnd; i++ {
			if hasRepeatingSection(strconv.Itoa(i)) {
				sum += i
			}
		}

	}
	fmt.Println(sum)
}

func hasRepeatingSection(part string) bool {
	for secLen := len(part) / 2; secLen > 0; secLen-- {
		secCount := len(part) / secLen
		if secLen*secCount != len(part) {
			continue
		}
		if allSecsEqual(part, secLen, secCount) {
			return true
		}
	}

	return false
}

func allSecsEqual(part string, secLen int, secCount int) bool {
	firstSec := part[:secLen]
	for secIdx := 1; secIdx < secCount; secIdx++ {
		sec := part[secIdx*secLen : secIdx*secLen+secLen]
		if sec != firstSec {
			return false
		}
	}
	return true
}
