package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/dakyskye/AoC/2023/utils"
)

func main() {
	input := utils.Read("./input.txt")
	part1(input)
	part2(input)
}

func part1(input string) {
	values := retrieveOriginalCalibrationValues(input)
	fmt.Println(utils.Sum(values))
}

func part2(input string) {
	values := retrieveOriginalCalibrationValuesEnhanced(input)
	fmt.Println(utils.Sum(values))
}

func retrieveOriginalCalibrationValues(data string) (res []int) {
	scanner := bufio.NewScanner(strings.NewReader(data))

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		var (
			idx      = 0
			numLeft  int
			numRight int
		)

		for ; idx < len(line); idx++ {
			if line[idx] >= '0' && line[idx] <= '9' {
				numLeft = utils.StringToInt(string(line[idx]))
				break
			}
		}

		for i := len(line) - 1; i >= idx; i-- {
			if line[i] >= '0' && line[i] <= '9' {
				numRight = utils.StringToInt(string(line[i]))
				break
			}
		}

		res = append(res, 10*numLeft+numRight)
	}

	return
}

var digitsAsLetters = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func retrieveOriginalCalibrationValuesEnhanced(data string) (res []int) {
	scanner := bufio.NewScanner(strings.NewReader(data))

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		var nums []int

		for i := 0; i < len(line); i++ {
			if line[i] >= '0' && line[i] <= '9' {
				nums = append(nums, utils.StringToInt(string(line[i])))
				continue
			}

			for word, digit := range digitsAsLetters {
				if strings.HasPrefix(line[i:], word) {
					nums = append(nums, digit)
					break
				}
			}
		}

		num := 10*nums[0] + nums[len(nums)-1]
		res = append(res, num)
	}

	return
}
