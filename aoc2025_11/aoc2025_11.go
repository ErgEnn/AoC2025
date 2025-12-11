package main

import (
	"AoC2025/util"
	"fmt"
	"strings"
)

const example = `aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out`

const example2 = `svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out`

func main() {
	real := util.ReadInput()

	//part1(real)
	part2(real)
}

func prep(input string) map[string]*Vertex {
	lines := strings.Split(input, "\n")
	index := make(map[string]*Vertex)
	for _, line := range lines {
		parts := strings.Fields(line)
		current := parts[0][:len(parts[0])-1]
		v, exists := index[current]
		if !exists {
			v = &Vertex{
				name: current,
			}

		}
		for _, child := range parts[1:] {
			c, exists := index[child]
			if !exists {
				c = &Vertex{
					name: child,
				}
				index[child] = c
			}
			v.Children = append(v.Children, c)
		}
		index[current] = v
	}
	return index
}

type Vertex struct {
	name     string
	Children []*Vertex
}

type VertexPair struct {
	First  *Vertex
	Second *Vertex
}

func recurse(current *Vertex, end *Vertex, memory map[VertexPair]int) int {
	if current == end {
		return 1
	}
	sum := 0
	for _, child := range current.Children {
		val, exists := memory[VertexPair{child, end}]
		if !exists {
			val = recurse(child, end, memory)
			memory[VertexPair{child, end}] = val
			sum += val
		} else {
			sum += val
		}

	}
	return sum
}

func part1(input string) {
	index := prep(input)
	start := index["you"]
	end := index["out"]
	memory := make(map[VertexPair]int)
	fmt.Println(recurse(start, end, memory))
}

func part2(input string) {
	index := prep(input)
	start := index["svr"]
	fft := index["fft"]
	dac := index["dac"]
	out := index["out"]
	memory := make(map[VertexPair]int)
	sToF := recurse(start, fft, memory)
	fToD := recurse(fft, dac, memory)
	dtoO := recurse(dac, out, memory)
	fmt.Println(sToF * fToD * dtoO)
}
