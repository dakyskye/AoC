package utils

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
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

func ReverseSlice[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func NewScanner(input string) *bufio.Scanner {
	return bufio.NewScanner(strings.NewReader(input))
}
