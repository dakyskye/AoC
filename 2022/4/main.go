package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/dakyskye/AoC/2022/utils"
)

type elvesPair struct {
	first  section
	second section
}

type section struct {
	start int
	end   int
}

func main() {
	input := utils.Read("./input.txt")

	part1(input)
	part2(input)
}

func part1(data string) {
	pairs := getPairs(data)

	fullOverlapCounter := 0
	for _, pair := range pairs {
		if sectionsFullyOverlap(pair.first, pair.second) {
			fullOverlapCounter += 1
		}
	}

	fmt.Println(fullOverlapCounter)
}

func part2(data string) {
	pairs := getPairs(data)

	overlapCounter := 0
	for _, pair := range pairs {
		if sectionsOverlap(pair.first, pair.second) {
			overlapCounter += 1
		}
	}

	fmt.Println(overlapCounter)
}

func sectionsOverlap(sect1, sect2 section) bool {
	if sect1.start >= sect2.start && sect1.start <= sect2.end {
		return true
	}

	if sect2.start >= sect1.start && sect2.start <= sect1.end {
		return true
	}

	return false
}

func sectionsFullyOverlap(sect1, sect2 section) bool {
	if sect1.start >= sect2.start && sect1.end <= sect2.end {
		return true
	}

	if sect2.start >= sect1.start && sect2.end <= sect1.end {
		return true
	}

	return false
}

func getPairs(data string) (res []elvesPair) {
	scanner := bufio.NewScanner(strings.NewReader(data))

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			panic("malformed input")
		}

		var pair elvesPair

		for i := 0; i < 2; i++ {
			sectionParts := strings.Split(parts[i], "-")
			if len(sectionParts) != 2 {
				panic("malformed input")
			}

			sect := section{
				start: utils.StringToInt(sectionParts[0]),
				end:   utils.StringToInt(sectionParts[1]),
			}

			if i == 0 {
				pair.first = sect
			} else {
				pair.second = sect
			}
		}

		res = append(res, pair)
	}

	return
}
