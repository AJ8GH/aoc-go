package util

import (
	"log"
	"os"
	"strings"
)

const (
	input         = "resources/input.txt"
	example       = "resources/example.txt"
	example2      = "resources/example-2.txt"
	lineDelimiter = "\n"
	errorMessage  = "Error reading file %q: %q"
)

func ReadInput() []string {
	return Read(input)
}

func ReadExample() []string {
	return Read(example)
}

func ReadExample2() []string {
	return Read(example2)
}

func Read(file string) []string {
	bytes, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf(errorMessage, file, err.Error())
	}

	return strings.Split(string(bytes), lineDelimiter)
}
