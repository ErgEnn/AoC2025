package util

import "iter"

func Distinct(s iter.Seq[int]) map[int]int {
	m := make(map[int]int)
	for i := range s {
		m[i]++
	}
	return m
}
