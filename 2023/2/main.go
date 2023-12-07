package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/dakyskye/AoC/2023/utils"
)

type Game struct {
	id   int
	sets [][]Set
}

type Set struct {
	amount int
	colour CubeColour
}

type CubeColour string

const (
	cubeColourRed   CubeColour = "red"
	cubeColourGreen CubeColour = "green"
	cubeColourBlue  CubeColour = "blue"
)

var cubeAmounts = map[CubeColour]int{
	cubeColourRed:   12,
	cubeColourGreen: 13,
	cubeColourBlue:  14,
}

func main() {
	input := utils.Read("./input.txt")
	part1(input)
	part2(input)
}

func part1(input string) {
	scanner := bufio.NewScanner(strings.NewReader(input))

	games := make([]*Game, 0)

	for scanner.Scan() {
		line := scanner.Text()

		games = append(games, parseGame(line))
	}

	sum := 0
	for _, game := range games {
		if isGamePossible(*game) {
			sum += game.id
		}
	}

	fmt.Println(sum)
}

func part2(input string) {
	scanner := bufio.NewScanner(strings.NewReader(input))

	games := make([]*Game, 0)

	for scanner.Scan() {
		line := scanner.Text()

		games = append(games, parseGame(line))
	}

	sum := 0
	for _, game := range games {
		sum += powerOfGame(*game)
	}

	fmt.Println(sum)
}

func parseGame(line string) (game *Game) {
	game = new(Game)

	gameAndSet := strings.Split(line, ":")
	sets := strings.Split(gameAndSet[1][1:], "; ")

	game.id = utils.StringToInt(strings.Split(gameAndSet[0], " ")[1])
	game.sets = make([][]Set, len(sets))

	for i, set := range sets {
		items := strings.Split(set, ", ")

		game.sets[i] = make([]Set, len(items))

		for j, item := range items {
			game.sets[i][j].amount = utils.StringToInt(strings.Split(item, " ")[0])
			game.sets[i][j].colour = CubeColour(strings.Split(item, " ")[1])
		}
	}

	return
}

func powerOfGame(game Game) int {
	reds, greens, blues := 0, 0, 0

	for _, set := range game.sets {
		for _, item := range set {
			switch item.colour {
			case cubeColourRed:
				if item.amount > reds {
					reds = item.amount
				}
			case cubeColourGreen:
				if item.amount > greens {
					greens = item.amount
				}
			case cubeColourBlue:
				if item.amount > blues {
					blues = item.amount
				}
			}
		}
	}

	return reds * greens * blues
}

func isGamePossible(game Game) bool {
	for _, set := range game.sets {
		reds, greens, blues := 0, 0, 0

		for _, item := range set {
			switch item.colour {
			case cubeColourRed:
				reds += item.amount
			case cubeColourGreen:
				greens += item.amount
			case cubeColourBlue:
				blues += item.amount
			}
		}

		if reds > cubeAmounts[cubeColourRed] {
			return false
		}

		if greens > cubeAmounts[cubeColourGreen] {
			return false
		}

		if blues > cubeAmounts[cubeColourBlue] {
			return false
		}
	}

	return true
}
