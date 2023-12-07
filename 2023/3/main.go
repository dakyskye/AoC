package main

import (
	"fmt"

	"github.com/dakyskye/AoC/2023/utils"
)

type Schema struct {
	lines []Line
}

type Line struct {
	length  int
	symbols []Symbol
	numbers []Number
}

type Symbol struct {
	symbol   rune
	position int
}

type Number struct {
	number   int
	startIdx int
	endIdx   int
}

func main() {
	input := utils.Read("./input.txt")
	part1(input)
	part2(input)
}

func part1(input string) {
	scanner := utils.NewScanner(input)

	var schema Schema

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		schema.lines = append(schema.lines, *parseLine(line))
	}

	sum := 0
	for i, line := range schema.lines {
		if len(line.numbers) == 0 {
			continue
		}

		for _, number := range line.numbers {
			lineBeforeIdx := i - 1
			if lineBeforeIdx < 0 {
				lineBeforeIdx = 0
			}
			lineAfterIdx := i + 1
			if lineAfterIdx > len(schema.lines)-1 {
				lineAfterIdx = len(schema.lines) - 1
			}

			isAdjacent := false
			for j := lineBeforeIdx; j <= lineAfterIdx; j++ {
				if lineHasSymbolInRange(schema.lines[j], number) {
					isAdjacent = true
					break
				}
			}

			if isAdjacent {
				sum += number.number
			}
		}
	}

	fmt.Println(sum)
}

func part2(input string) {
	scanner := utils.NewScanner(input)

	var schema Schema

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		schema.lines = append(schema.lines, *parseLine(line))
	}

	sum := 0
	for i, line := range schema.lines {
		if len(line.symbols) == 0 {
			continue
		}

		for _, symbol := range line.symbols {
			if string(symbol.symbol) != "*" {
				continue
			}

			lineBeforeIdx := i - 1
			if lineBeforeIdx < 0 {
				lineBeforeIdx = 0
			}
			lineAfterIdx := i + 1
			if lineAfterIdx > len(schema.lines)-1 {
				lineAfterIdx = len(schema.lines) - 1
			}

			var adjacentNumbers []Number
			for j := lineBeforeIdx; j <= lineAfterIdx; j++ {
				nums := getNumbersAdjacentToSymbol(schema.lines[j], symbol)
				if len(nums) > 0 {
					adjacentNumbers = append(adjacentNumbers, nums...)
				}
			}

			if len(adjacentNumbers) == 2 {
				sum += adjacentNumbers[0].number * adjacentNumbers[1].number
			}
		}
	}

	fmt.Println(sum)
}

func lineHasSymbolInRange(line Line, number Number) bool {
	acceptedIndexStart := number.startIdx - 1
	if acceptedIndexStart < 0 {
		acceptedIndexStart = 0
	}

	acceptedIndexEnd := number.endIdx + 1
	if acceptedIndexEnd > line.length-1 {
		acceptedIndexEnd = line.length - 1
	}

	for _, symbol := range line.symbols {
		if symbol.position < acceptedIndexStart {
			continue
		}

		if symbol.position > acceptedIndexEnd {
			break
		}

		return true
	}

	return false
}

func getNumbersAdjacentToSymbol(line Line, symbol Symbol) (nums []Number) {
	acceptedIndexStart := symbol.position - 1
	if acceptedIndexStart < 0 {
		acceptedIndexStart = 0
	}
	acceptedIndexEnd := symbol.position + 1
	if acceptedIndexEnd > line.length-1 {
		acceptedIndexEnd = line.length - 1
	}

	if len(line.numbers) == 0 {
		return nil
	}

	for _, number := range line.numbers {
		if number.endIdx < acceptedIndexStart {
			continue
		}

		if number.startIdx > acceptedIndexEnd {
			break
		}

		nums = append(nums, number)
	}

	return
}

func parseLine(line string) *Line {
	res := new(Line)

	res.length = len(line)

	tempNum := ""

	for i, char := range line {
		if char >= '0' && char <= '9' {
			tempNum += string(char)
			continue
		}

		if tempNum != "" {
			res.numbers = append(res.numbers, Number{
				number:   utils.StringToInt(tempNum),
				startIdx: i - len(tempNum),
				endIdx:   i - 1,
			})

			tempNum = ""
		}

		if char != '.' {
			res.symbols = append(res.symbols, Symbol{
				symbol:   char,
				position: i,
			})
		}

	}

	if tempNum != "" {
		res.numbers = append(res.numbers, Number{
			number:   utils.StringToInt(tempNum),
			startIdx: len(line) - len(tempNum),
			endIdx:   len(line),
		})
	}

	return res
}
