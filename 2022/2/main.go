package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/dakyskye/AoC/2022/utils"
)

func main() {
	input := utils.Read("./input.txt")

	part1(input)
	part2(input)
}

func part1(data string) {
	rounds := extractRounds(data)
	score := 0

	for _, round := range rounds {
		enemy, me := enemyMove(round[0]), myMove(round[1])
		result := playRound(enemy, me)

		score += moveScore[me] + roundResultScore[result]
	}

	fmt.Println(score)
}

func part2(data string) {
	rounds := extractRounds(data)
	score := 0

	for _, round := range rounds {
		enemy, action := enemyMove(round[0]), actionToTake(round[1])
		move := playRoundSmart(enemy, action)

		score += moveScore[move] + actionScore[action]
	}

	fmt.Println(score)
}

func extractRounds(data string) (res [][2]byte) {
	scanner := bufio.NewScanner(strings.NewReader(data))

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			panic("malformed input")
		}

		part1 := parts[0][0]
		part2 := parts[1][0]

		switch enemyMove(part1) {
		case enemyMoveRock, enemyMovePaper, enemyMoveScrissors:
		default:
			panic("malformed input")
		}

		switch myMove(part2) {
		case myMoveRock, myMovePaper, myMoveScrissors:
		default:
			panic("malformed input")
		}

		res = append(res, [2]byte{part1, part2})
	}

	return
}
