package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/dakyskye/AoC/2022/utils"
)

type rucksack struct {
	firstCompartment  []rune
	secondCompartment []rune
}

var itemPriority = make(map[rune]int)

func init() {
	priority := 0

	for elem := 'a'; elem <= 'z'; elem++ {
		priority++
		itemPriority[elem] = priority
	}

	for elem := 'A'; elem <= 'Z'; elem++ {
		priority++
		itemPriority[elem] = priority
	}
}

func main() {
	input := utils.Read("./input.txt")

	part1(input)
	part2(input)
}

func part1(data string) {
	rucksacks := getRucksacks(data)

	var sharedItems []rune
	for _, rucksack := range rucksacks {
		sharedItems = append(sharedItems, findSharedItem(rucksack))
	}

	prioritySum := sumItemsPriority(sharedItems)
	fmt.Println(prioritySum)
}

func part2(data string) {
	rucksacks := getRucksacks(data)

	var groupBadges []rune
	for _, group := range groupsOfThree(rucksacks) {
		groupBadges = append(groupBadges, findSharedItemInGroup(group))
	}

	prioritySum := sumItemsPriority(groupBadges)
	fmt.Println(prioritySum)
}

func findSharedItemInGroup(rucksackItems [3]string) (res rune) {
	for _, r1Item := range rucksackItems[0] {
		if strings.ContainsRune(rucksackItems[1], r1Item) && strings.ContainsRune(rucksackItems[2], r1Item) {
			return r1Item
		}
	}
	return
}

func groupsOfThree(rucksacks []rucksack) (res [][3]string) {
	group := [3]string{}
	for index, rucksack := range rucksacks {
		group[index%3] = fmt.Sprintf("%s%s", string(rucksack.firstCompartment), string(rucksack.secondCompartment))
		if (index+1)%3 == 0 {
			res = append(res, group)
			group = [3]string{}
		}
	}

	return
}

func sumItemsPriority(items []rune) (sum int) {
	for _, item := range items {
		sum += itemPriority[item]
	}
	return
}

func findSharedItem(r rucksack) (res rune) {
	for _, i1 := range r.firstCompartment {
		for _, i2 := range r.secondCompartment {
			if i1 == i2 {
				res = i1
				break
			}
		}
	}

	return
}

func getRucksacks(data string) (res []rucksack) {
	scanner := bufio.NewScanner(strings.NewReader(data))

	for scanner.Scan() {
		line := scanner.Text()
		if len(line)%2 != 0 {
			panic("malformed input")
		}

		rucksack := rucksack{
			firstCompartment:  []rune(line[:len(line)/2]),
			secondCompartment: []rune(line[len(line)/2:]),
		}

		res = append(res, rucksack)
	}

	return
}
