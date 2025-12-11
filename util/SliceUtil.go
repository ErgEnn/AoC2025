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

func SumSlice(i []int) int {
	sum := 0
	for _, v := range i {
		sum += v
	}
	return sum
}

func Compositions(sum int, maxes []int) iter.Seq[[]int] {
	return func(yield func([]int) bool) {
		n := len(maxes)
		a := make([]int, n)

		// Scratch buffer for output
		out := make([]int, n)

		// Recursive generator in a closure
		var dfs func(i, remaining int) bool
		dfs = func(i, remaining int) bool {
			if i == n {
				if remaining == 0 {
					copy(out, a)
					return yield(out)
				}
				return true
			}

			// max allowed at this position
			hi := maxes[i]
			if hi > remaining {
				hi = remaining
			}

			for v := 0; v <= hi; v++ {
				a[i] = v
				if !dfs(i+1, remaining-v) {
					return false
				}
			}
			return true
		}

		dfs(0, sum)
	}
}
