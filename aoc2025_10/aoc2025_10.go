package main

import (
	"AoC2025/util"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

const example = `[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}`

func main() {
	real := util.ReadInput()

	//part1(real)
	part2(real)
}

func prep(input string) []string {
	return strings.Split(input, "\n")
}

func parseInitialLightsState(s string) int {
	val := 0
	for i := 1; i < len(s)-1; i++ {
		val = val << 1
		isOn := s[i] == '#'
		if isOn {
			val = val | 1
		}
	}
	return val
}

func parseButtons(buttons []string, noOfLights int) []int {
	masks := make([]int, len(buttons))
	for i, button := range buttons {
		numbers := strings.Split(button[1:len(button)-1], ",")
		mask := 0
		for i := 0; i < noOfLights; i++ {
			mask = mask << 1
			iasS := strconv.Itoa(i)
			if slices.Contains(numbers, iasS) {
				mask = mask | 1
			}
		}
		masks[i] = mask
	}
	return masks
}

func leastPressesForLights(state int, buttons []int) int {
	queue := make([]int, 1)
	queue[0] = state
	depths := make([]int, 1)
	depths[0] = 0
	for i := 0; i < len(queue); i++ {
		current := queue[i]
		currentDepth := depths[i]
		if current == 0 {
			return currentDepth
		}
		for _, mask := range buttons {
			newState := current ^ mask
			if !slices.Contains(queue, newState) {
				queue = append(queue, newState)
				depths = append(depths, currentDepth+1)
			}
		}
	}
	panic("aaaaa")
}

func part1(input string) {
	lines := prep(input)
	sum := 0
	for _, line := range lines {
		parts := strings.Fields(line)
		state := parseInitialLightsState(parts[0])
		buttons := parseButtons(parts[1:len(parts)-1], len(parts[0])-2)
		sum += leastPressesForLights(state, buttons)
	}
	fmt.Println(sum)
}

func parseRequiredJoltageState(s string) []int {
	fields := strings.Split(s[1:len(s)-1], ",")
	ints := make([]int, len(fields))
	for i, field := range fields {
		ints[i] = util.ToInt(field)
	}
	return ints
}

func parseButtonMultiplierMasks(buttons []string, noOfLights int) [][]int {
	masks := make([][]int, len(buttons))
	for i, button := range buttons {
		numbers := strings.Split(button[1:len(button)-1], ",")
		mask := make([]int, noOfLights)
		for _, idx := range numbers {
			mask[util.ToInt(idx)] = 1
		}
		masks[i] = mask
	}
	return masks
}

func allZero(i []int) bool {
	for _, v := range i {
		if v != 0 {
			return false
		}
	}
	return true
}

func areEqual(a, b []int) bool {
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func contains(haystack [][]int, needle []int) int {
	for i, v := range haystack {
		if areEqual(v, needle) {
			return i
		}
	}
	return -1
}

func anyBelowZero(i []int) bool {
	for _, v := range i {
		if v < 0 {
			return true
		}
	}
	return false
}

func multiply(i []int) int {
	total := 1
	for _, v := range i {
		total *= v
	}
	return total
}

func maxValAndCount(i []int) (int, int) {
	max := 0
	count := 0
	for _, v := range i {
		if v > max {
			max = v
			count = 1
		} else if v == max {
			count++
		}
	}
	return max, count
}

func leastPressesForJoltage(desiredState []int, buttons [][]int) int {
	maxPressesForButton := make([]int, len(buttons))
	for i := range maxPressesForButton {
		maxPressesForButton[i] = math.MaxInt
	}
	for i, val := range desiredState {
		for j, mask := range buttons {
			if mask[i] == 1 {
				maxPressesForButton[j] = util.Min(maxPressesForButton[j], val)
			}
		}
	}
	pressedButtons := make([]int, len(buttons))
	for i := range pressedButtons {
		pressedButtons[i] = -1
	}
	result := recurse(desiredState, 0, buttons, maxPressesForButton, pressedButtons)
	fmt.Println(result)
	return util.SumSlice(result)
}

func recurse(desiredStates []int, desiredStateIdx int, buttons [][]int, maxPressesForButton []int, pressedButtons []int) []int {
	if desiredStateIdx == len(desiredStates) {
		return pressedButtons
	}
	desiredVal := desiredStates[desiredStateIdx]
	existingVal := 0
	maxPressesForButtons := make([]int, 0)
	for i, mask := range buttons {
		if mask[desiredStateIdx] == 1 {
			if pressedButtons[i] == -1 {
				maxPressesForButtons = append(maxPressesForButtons, maxPressesForButton[i])
			} else {
				existingVal += pressedButtons[i]
			}
		}
	}
	newDesiredVal := desiredVal - existingVal
	if newDesiredVal < 0 {
		return []int{math.MaxInt}
	}
	if newDesiredVal == 0 {
		newPressedButtons := make([]int, len(pressedButtons))
		copy(newPressedButtons, pressedButtons)
		for i, mask := range buttons {
			if mask[desiredStateIdx] == 1 && pressedButtons[i] == -1 {
				newPressedButtons[i] = 0
			}
		}
		return recurse(desiredStates, desiredStateIdx+1, buttons, maxPressesForButton, newPressedButtons)
	}
	combinationIter := util.Compositions(desiredVal-existingVal, maxPressesForButtons)
	bestPresses := []int{math.MaxInt}
	for combination := range combinationIter {
		ci := 0
		newPressedButtons := make([]int, len(pressedButtons))
		copy(newPressedButtons, pressedButtons)
		for i, mask := range buttons {
			if mask[desiredStateIdx] == 1 && pressedButtons[i] == -1 {
				newPressedButtons[i] = combination[ci]
				ci++
			}
		}
		result := recurse(desiredStates, desiredStateIdx+1, buttons, maxPressesForButton, newPressedButtons)
		if len(result) != 0 && util.SumSlice(result) < util.SumSlice(bestPresses) {
			bestPresses = result
		}
	}
	return bestPresses
}

func test(desiredState []int, buttons [][]int) int {
	for i, state := range desiredState {
		fmt.Print("{")
		for _, button := range buttons {
			fmt.Print(button[i], ",")
		}
		fmt.Print(state)
		fmt.Print("},")
	}
	return 0
}

func part2(input string) {
	lines := prep(input)
	sum := 0
	for i, line := range lines {
		parts := strings.Fields(line)
		state := parseRequiredJoltageState(parts[len(parts)-1])
		buttons := parseButtonMultiplierMasks(parts[1:len(parts)-1], len(parts[0])-2)
		fmt.Print(i)
		test(state, buttons)
		fmt.Println()
		//least := leastPressesForJoltage(state, buttons)
		//fmt.Printf("%d: %d \n", i, least)
		//sum += least
	}
	fmt.Println(sum)
}
