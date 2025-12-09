package util

import "iter"

func Distinct(s iter.Seq[int]) map[int]int {
	m := make(map[int]int)
	for i := range s {
		m[i]++
	}
	return m
}

type CoordPair struct {
	A, B Coord
}

func PairwiseCombinations(s []Coord) iter.Seq[CoordPair] {
	return func(yield func(CoordPair) bool) {
		for i, coord := range s {
			for j := i + 1; j < len(s); j++ {
				yield(CoordPair{coord, s[j]})
			}
		}
	}
}
