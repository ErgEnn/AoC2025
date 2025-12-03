package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const example = `
987654321111111
811111111111119
234234234234278
818181911112111
`

const real = ``

func main() {
	part1(example)
	part2(example)
}

func prep(input string) []string {
	return strings.Split(input, "\n")
}

func part1(input string) {
	lines := prep(input)
	sum := 0
	for _, line := range lines {
		msDigit := 0
		msDigitIdx := -1
		for i := 0; i < len(line)-1; i++ {
			char := line[i : i+1]
			val, _ := strconv.Atoi(char)
			if val > msDigit {
				msDigit = val
				msDigitIdx = i
			}
		}

		lsDigit := 0
		for i := len(line) - 1; i > msDigitIdx; i-- {
			char := line[i : i+1]
			val, _ := strconv.Atoi(char)
			if val > lsDigit {
				lsDigit = val
			}
		}

		val := msDigit*10 + lsDigit
		sum += val
	}
	fmt.Println(sum)
}

func analyze(line []int, digits int, start int) uint64 {
	if digits < 1 {
		return 0
	}

	maxDigit := 0
	maxDigitIdx := -1
	for i := start; i < len(line)-(digits-1); i++ {
		val := line[i]
		if val > maxDigit {
			maxDigit = val
			maxDigitIdx = i
		}
	}

	val := uint64(maxDigit) * uint64(math.Pow10(digits-1))
	return val + analyze(line, digits-1, maxDigitIdx+1)
}

func part2(input string) {
	lines := prep(input)
	var sum uint64 = 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		list := make([]int, len(line))
		for i := 0; i < len(line); i++ {
			val, _ := strconv.Atoi(line[i : i+1])
			list[i] = val
		}
		val := analyze(list, 12, 0)
		sum += val
	}
	fmt.Println(sum)
}
