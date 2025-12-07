package util

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func ReadInput() string {
	_, thisFile, _, _ := runtime.Caller(1)
	base := filepath.Dir(thisFile)

	data, err := os.ReadFile(filepath.Join(base, "input.txt"))
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
