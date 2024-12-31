package util

import "os"

const (
	mode       = 0600
	outputFile = "output.txt"
)

func Create(content string) {
	CreateFile(content, outputFile)
}

func CreateFile(content, file string) {
	write(content, os.O_WRONLY|os.O_CREATE, file)
}

func Append(content string) {
	AppendFile(content, outputFile)
}

func AppendFile(content, file string) {
	write(content, os.O_APPEND|os.O_WRONLY|os.O_CREATE, file)
}

func write(content string, flag int, file string) {
	f, err := os.OpenFile(file, flag, mode)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(content); err != nil {
		panic(err)
	}
}
