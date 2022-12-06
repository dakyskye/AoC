package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"

	"github.com/dakyskye/AoC/2022/utils"
)

type stack struct {
	crates []byte
}

type instruction struct {
	amount int
	from   int
	to     int
}

type crateMover int8

const (
	crateMover9000 crateMover = iota
	crateMover9001
)

var separatorPattern = regexp.MustCompile(`(?m)^$\n`)

func main() {
	input := utils.Read("./input.txt")

	part1(input)
	part2(input)
}

func part1(data string) {
	stacks, instructions := getStacksAndInstructions(data)

	for _, inst := range instructions {
		alterStacks(stacks, inst, crateMover9000)
	}

	var topCrates []byte
	for _, stack := range stacks {
		topCrates = append(topCrates, stack.crates[len(stack.crates)-1])
	}

	fmt.Println(string(topCrates))
}

func part2(data string) {
	stacks, instructions := getStacksAndInstructions(data)

	for _, inst := range instructions {
		alterStacks(stacks, inst, crateMover9001)
	}

	var topCrates []byte
	for _, stack := range stacks {
		topCrates = append(topCrates, stack.crates[len(stack.crates)-1])
	}

	fmt.Println(string(topCrates))
}

func alterStacks(stacks []stack, inst instruction, cm crateMover) {
	var (
		from      = inst.from - 1
		to        = inst.to - 1
		takeIndex = len(stacks[from].crates) - inst.amount
	)

	taken := stacks[from].crates[takeIndex:]
	stacks[from].crates = stacks[from].crates[:takeIndex]

	if cm == crateMover9000 {
		utils.ReverseSlice(taken)
	}

	stacks[to].crates = append(stacks[to].crates, taken...)
}

func getStacksAndInstructions(data string) ([]stack, []instruction) {
	chunks := separatorPattern.Split(data, -1)
	if len(chunks) != 2 {
		panic("malformed input")
	}

	return getStacks(strings.TrimRight(chunks[0], "\n")), getInstructions(strings.TrimRight(chunks[1], "\n"))
}

func getStacks(data string) (stacks []stack) {
	lastLineBegIdx := strings.LastIndex(data, "\n") + 1
	lastLine := strings.ReplaceAll(data[lastLineBegIdx:], " ", "")

	stackAmount := len(lastLine)
	if stackAmount > 9 {
		panic("malformed input")
	}

	lineLength := -1 + stackAmount + 3*stackAmount
	stacks = make([]stack, stackAmount)

	scanner := bufio.NewScanner(strings.NewReader(data[:lastLineBegIdx-1]))

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) != lineLength {
			panic("malformed input")
		}

		for i := 1; i <= lineLength; i += 4 {
			crate := line[i]
			if crate == ' ' {
				continue
			}

			idx := 0
			for crateIndex := 1; crateIndex < lineLength; crateIndex += 4 {
				if crateIndex == i {
					break
				}
				idx += 1
			}

			stacks[idx].crates = append(stacks[idx].crates, crate)
		}
	}

	for i := range stacks {
		utils.ReverseSlice(stacks[i].crates)
	}

	return
}

func getInstructions(data string) (instructions []instruction) {
	scanner := bufio.NewScanner(strings.NewReader(data))

	for scanner.Scan() {
		line := scanner.Text()

		var (
			idx   = 0
			parts []string
		)

		for {
			idx += strings.IndexRune(line[idx:], ' ')
			nextIdx := idx + strings.IndexRune(line[idx+1:], ' ') + 1

			if nextIdx == idx {
				parts = append(parts, line[idx+1:])
				break
			}

			parts = append(parts, line[idx+1:nextIdx])
			idx = nextIdx + 1
		}

		if len(parts) != 3 {
			panic("malformed input")
		}

		inst := instruction{
			amount: utils.StringToInt(string(parts[0])),
			from:   utils.StringToInt(string(parts[1])),
			to:     utils.StringToInt(string(parts[2])),
		}

		instructions = append(instructions, inst)
	}

	return
}
