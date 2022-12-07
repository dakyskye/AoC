package main

import (
	"fmt"

	"github.com/dakyskye/AoC/2022/utils"
)

func main() {
	input := utils.Read("./input.txt")

	part1(input)
	part2(input)
}

func part1(data string) {
	fmt.Println(findStartOfPacketMarkerIndex(data))
}

func part2(data string) {
	fmt.Println(findStartOfMessageMarkerIndex(data))
}

func findStartOfPacketMarkerIndex(data string) int {
	isUnique := func(data string) bool {
		runes := make(map[rune]bool)
		for _, r := range data {
			if _, ok := runes[r]; ok {
				return false
			}
			runes[r] = true
		}
		return true
	}

	for i := 3; i < len(data); i++ {
		if isUnique(data[i-3 : i+1]) {
			return i + 1
		}
	}

	return -1
}

func findStartOfMessageMarkerIndex(data string) int {
	isUnique := func(data string) bool {
		runes := make(map[rune]bool)
		for _, r := range data {
			if _, ok := runes[r]; ok {
				return false
			}
			runes[r] = true
		}
		return true
	}

	for i := 13; i < len(data); i++ {
		if isUnique(data[i-13 : i+1]) {
			return i + 1
		}
	}

	return -1
}
