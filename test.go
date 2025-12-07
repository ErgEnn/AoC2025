package main

import "fmt"

func main() {
	m := map[int]bool{
		1: false,
		2: true,
	}

	if !m[0] {
		fmt.Println("0")
	}
	if !m[1] {
		fmt.Println("1")
	}
	if !m[2] {
		fmt.Println("2")
	}
	foo(m[0])

	m2 := map[int]*string{}
	a := "a"
	m2[1] = &a

	if m2[0] == nil {
		fmt.Println("0s")
	}
	if m2[1] == &a {
		fmt.Println("1s")
	}
}

func foo(b bool) {
	fmt.Printf("%v %v\n", "foo", b)
}
