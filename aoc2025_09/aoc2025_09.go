package main

import (
	"AoC2025/util"
	"fmt"
	"strings"
)

const example = ``

func main() {
	_ = util.ReadInput()

	part1(example)
	//part2(example)
}

func prep(input string) []string {
	return strings.Split(input, "\r\n")
}

func part1(input string) {
	i := prep(input)
	fmt.Println(i)
}

func part2(input string) {
	i := prep(input)
	fmt.Println(i)
}
