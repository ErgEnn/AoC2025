package util

import (
	"strconv"
	"strings"
)

func Split2(s string, sep string) (string, string) {
	parts := strings.Split(s, sep)
	return parts[0], strings.TrimSpace(parts[1])
}

func Split2Func[T any](s string, f func(string) T) (T, T) {
	p1, p2 := Split2(s, ",")
	return f(p1), f(p2)
}

func Split3(s string, sep string) (string, string, string) {
	parts := strings.Split(s, sep)
	return parts[0], parts[1], strings.TrimSpace(parts[2])
}

func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(-1)
	}
	return i
}
