package main

import (
	"fmt"
	"strings"

	"github.com/dakyskye/AoC/2023/utils"
)

func main() {
	input := utils.Read("./input.txt")
	part1(input)
	part2(input)
}

type Card struct {
	id             int
	winningNumbers []int
	otherNumbers   []int
}

func part1(input string) {
	scanner := utils.NewScanner(input)

	cards := make([]*Card, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		cards = append(cards, parseCard(line))
	}

	sum := 0
	for _, card := range cards {
		sum += calculateCardWorth(*card)
	}

	fmt.Println(sum)
}

func part2(input string) {
	scanner := utils.NewScanner(input)

	cards := make([]*Card, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		cards = append(cards, parseCard(line))
	}

	cardCounts := make(map[int]int)
	copiesCache := make(map[int][]int)

	for i := 0; i < len(cards); i++ {
		count, ok := cardCounts[cards[i].id]
		if !ok {
			cardCounts[cards[i].id] = 1
			copiesCache[cards[i].id] = getCopiesFromCard(*cards[i])
			i--
			continue
		}

		cache, ok := copiesCache[cards[i].id]
		if !ok {
			count++
			cardCounts[cards[i].id]++
			copiesCache[cards[i].id] = getCopiesFromCard(*cards[i])
			cache = copiesCache[cards[i].id]
		}

		for _, c := range cache {
			cardCounts[c] += count
		}
	}

	sum := 0

	for _, count := range cardCounts {
		sum += count
	}

	fmt.Println(sum)
}

func parseCard(line string) *Card {
	card := new(Card)

	cardAndNumbers := strings.Split(line, ":")
	cardInfo := strings.Split(cardAndNumbers[0], " ")
	numbersSet := strings.Split(cardAndNumbers[1], " | ")

	card.id = utils.StringToInt(cardInfo[len(cardInfo)-1])

	for _, numStr := range strings.Split(numbersSet[0][1:], " ") {
		if numStr == "" {
			continue
		}

		card.winningNumbers = append(card.winningNumbers, utils.StringToInt(numStr))
	}

	for _, numStr := range strings.Split(numbersSet[1], " ") {
		if numStr == "" {
			continue
		}

		card.otherNumbers = append(card.otherNumbers, utils.StringToInt(numStr))
	}

	return card
}

func calculateCardWorth(card Card) int {
	points := 0

	for _, winningNum := range card.winningNumbers {
		for _, otherNum := range card.otherNumbers {
			if winningNum == otherNum {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}
	}

	return points
}

func getCopiesFromCard(card Card) []int {
	copies := make([]int, 0)

	counter := 1

	for _, winningNum := range card.winningNumbers {
		for _, otherNum := range card.otherNumbers {
			if winningNum == otherNum {
				copies = append(copies, card.id+counter)
				counter++
			}
		}
	}

	return copies
}
