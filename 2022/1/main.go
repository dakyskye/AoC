package main

import (
	"bufio"
	"fmt"
	"sort"
	"strings"

	"github.com/dakyskye/AoC/2022/utils"
)

func main() {
	input := utils.Read("./input.txt")

	part1(input)
	part2(input)
}

func part1(data string) {
	calories := calculateCalories(data)
	mostCalories := findMax(calories)
	fmt.Println(mostCalories)
}

func part2(data string) {
	calories := calculateCalories(data)
	sumMost3Calories := utils.Sum(findNMax(calories, 3))
	fmt.Println(sumMost3Calories)
}

func calculateCalories(data string) (res []int) {
	scanner := bufio.NewScanner(strings.NewReader(data))

	lastSum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			res = append(res, lastSum)
			lastSum = 0
			continue
		}
		lastSum += utils.StringToInt(line)
	}
	res = append(res, lastSum)
	return
}

func findMax(ints []int) int {
	max := ints[0]

	if len(ints) == 1 {
		return max
	}

	for i := 1; i < len(ints); i++ {
		if ints[i] > max {
			max = ints[i]
		}
	}

	return max
}

func findNMax(ints []int, n int) []int {
	if n > len(ints) {
		panic("invalid n")
	}

	intsCopy := make([]int, len(ints))
	copy(intsCopy, ints)

	sort.Ints(intsCopy)

	return intsCopy[len(intsCopy)-n:]
}
