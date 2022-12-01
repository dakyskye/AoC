package utils

import (
	"io"
	"os"
	"strconv"
)

func Read(path string) string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func StringToInt(str string) int {
	number, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return number
}

func Sum(ints []int) int {
	sum := 0
	for _, i := range ints {
		sum += i
	}
	return sum
}
