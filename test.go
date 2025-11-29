package main

import (
	"fmt"
	"strconv"
	"strings"
)

const data = `
1 2
3 4
10 20
`

func main() {
	lines := strings.Split(strings.TrimSpace(data), "\n")

	// map: parse sums
	sums := make([]int, 0, len(lines))
	for _, l := range lines {
		p := strings.Fields(l)
		a, _ := strconv.Atoi(p[0])
		b, _ := strconv.Atoi(p[1])
		sums = append(sums, a+b)
	}

	// filter: keep > 5
	filtered := make([]int, 0, len(sums))
	for _, v := range sums {
		if v > 5 {
			filtered = append(filtered, v)
		}
	}

	// reduce: sum them
	total := 0
	for _, v := range filtered {
		total += v
	}

	fmt.Println(total)
}
