package main

import (
	"fmt"
	"iter"
	"strconv"
	"strings"
)

const example = `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `

const real = ``

func main() {
	part1(example)
	part2(example)
}

func prep(input string) iter.Seq[iter.Seq2[int, string]] {
	lines := strings.Split(input, "\n")
	return func(yield func(iter.Seq2[int, string]) bool) {
		for i := len(lines) - 1; i >= 0; i-- {
			line := lines[i]
			fields := strings.Fields(line)
			yield(func(yield func(int, string) bool) {
				for i, field := range fields {
					yield(i, field)
				}
			})
		}
	}
}

func part1(input string) {
	next, _ := iter.Pull(prep(input))
	colIsAdd := make(map[int]bool)
	opsLine, _ := next()
	for i, op := range opsLine {
		if op == "+" {
			colIsAdd[i] = true
		} else {
			colIsAdd[i] = false
		}
	}
	cols := make([]int64, len(colIsAdd))
	line, _ := next()
	for i, field := range line {
		cols[i], _ = strconv.ParseInt(field, 10, 64)
	}
	for {
		line, ok := next()
		if !ok {
			break
		}
		for i, field := range line {
			val, _ := strconv.ParseInt(field, 10, 64)
			if colIsAdd[i] {
				cols[i] = cols[i] + val
			} else {
				cols[i] = cols[i] * val
			}
		}
	}
	sum := int64(0)
	for _, col := range cols {
		sum += col
	}
	fmt.Println(sum)
}

func part2(input string) {
	lines := strings.Split(input, "\n")
	total := int64(0)
	buffer := make([]int, 0)
	for colIdx := len(lines[0]) - 1; colIdx >= 0; colIdx-- {
		inner(lines, &buffer, &colIdx, &total)
	}
	fmt.Println(total)
}

func inner(lines []string, buffer *[]int, colIdx *int, total *int64) {
	colVal := 0
	for _, line := range lines {
		col := line[*colIdx]
		if col == ' ' {
			continue
		}
		if col == '+' {
			*colIdx--
			for _, val := range *buffer {
				colVal += val
			}
			*total += int64(colVal)
			colVal = 0
			*buffer = make([]int, 0)
			return
		}
		if col == '*' {
			*colIdx--
			for _, val := range *buffer {
				colVal *= val
			}
			*total += int64(colVal)
			colVal = 0
			*buffer = make([]int, 0)
			return
		}
		colAsInt, _ := strconv.Atoi(string(col))
		colVal = colVal*10 + colAsInt
	}
	*buffer = append(*buffer, colVal)
}
