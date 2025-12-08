package util

import (
	"strconv"
	"strings"
)

func Split2(s string, sep string) (string, string) {
	parts := strings.Split(s, sep)
	return parts[0], parts[1]
}

func Split3(s string, sep string) (string, string, string) {
	parts := strings.Split(s, sep)
	part2 := strings.TrimSpace(parts[2])
	return parts[0], parts[1], part2
}

func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(-1)
	}
	return i
}
