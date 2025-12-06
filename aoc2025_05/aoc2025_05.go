package main

import (
	"AoC2025/util"
	"fmt"
	"iter"
	"sort"
	"strconv"
	"strings"
)

const example = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

const real = ``

func main() {
	part1(example)
	part2(example)
}

func prep(input string) iter.Seq[iter.Seq[string]] {
	return func(yield func(iter.Seq[string]) bool) {
		lines := strings.Split(input, "\n")
		next, stop := iter.Pull(func(yield func(string) bool) {
			for _, line := range lines {
				yield(line)
			}
		})
		yield(func(yield func(string) bool) {
			for {
				line, _ := next()
				if line == "" {
					break
				}
				yield(line)
			}
		})
		yield(func(yield func(string) bool) {
			for {
				line, ok := next()
				if !ok {
					break
				}
				yield(line)
			}
		})
		stop()
	}

}

func part1(input string) {
	next, _ := iter.Pull(prep(input))
	ranges, _ := next()
	l_list := make([]int64, 0)
	h_list := make([]int64, 0)
	for line := range ranges {
		ls, hs := util.Split2(line, "-")
		l, _ := strconv.ParseInt(ls, 10, 64)
		h, _ := strconv.ParseInt(hs, 10, 64)
		l_list = append(l_list, l)
		h_list = append(h_list, h)
	}
	produce, _ := next()
	cnt := 0
	for product := range produce {
		productId, _ := strconv.ParseInt(product, 10, 64)
		if inner(productId, l_list, h_list) {
			cnt++
		}
	}
	fmt.Printf("total: %d\n", cnt)
}

func inner(productId int64, l_list []int64, h_list []int64) bool {
	for i, l := range l_list {
		h := h_list[i]
		if productId >= l && productId <= h {
			//fmt.Printf("%d-%d\t%d\n", l, h, productId)
			return true
		}
	}
	//fmt.Printf("NO MATCH %d\n", productId)
	return false
}

type LoHiPair struct {
	lo int64
	hi int64
}

func part2(input string) {
	next, _ := iter.Pull(prep(input))
	rangesIter, _ := next()
	ranges := make([]LoHiPair, 0)
	for line := range rangesIter {
		ls, hs := util.Split2(line, "-")
		l, _ := strconv.ParseInt(ls, 10, 64)
		h, _ := strconv.ParseInt(hs, 10, 64)
		ranges = append(ranges, LoHiPair{lo: l, hi: h})
	}
	sort.Slice(ranges, func(i, j int) bool { return ranges[i].lo < ranges[j].lo })
	prevRange := ranges[0]
	cnt := prevRange.hi - prevRange.lo + 1
	for _, currentRange := range ranges[1:] {
		if currentRange.hi <= prevRange.hi {
			continue
		}
		if currentRange.lo <= prevRange.hi {
			cnt += currentRange.hi - prevRange.hi
			prevRange = currentRange
			continue
		}
		if currentRange.lo > prevRange.hi {
			cnt += currentRange.hi - currentRange.lo + 1
			prevRange = currentRange
			continue
		}
	}
	//fmt.Println(ranges)
	fmt.Println(cnt)
}
