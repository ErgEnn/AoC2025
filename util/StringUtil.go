package util

import "strings"

func Split2(s string, sep string) (string, string) {
	parts := strings.Split(s, sep)
	return parts[0], parts[1]
}
