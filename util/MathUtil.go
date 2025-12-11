package util

func Max(a int64, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func MaxI(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Sum(m map[int]int) int {
	sum := 0
	for _, v := range m {
		sum += v
	}
	return sum
}
