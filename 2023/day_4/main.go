package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type card struct {
	winningCards, myCards map[int]bool //set
}

var cards = []card{}

func stringToIntMap(str string) map[int]bool {
	var nums = map[int]bool{}
	for _, strNum := range strings.Fields(str) {
		strNum = strings.TrimSpace(strNum)
		num, _ := strconv.Atoi(strNum)
		nums[num] = true
	}
	return nums
}

func parseFileToCards(path string) {
	data, _ := os.ReadFile(path)
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		line = strings.Split(line, ":")[1]
		line = strings.TrimSpace(line)
		winningCards := strings.TrimSpace(strings.Split(line, "|")[0])
		myCards := strings.TrimSpace(strings.Split(line, "|")[1])

		card := card{stringToIntMap(winningCards), stringToIntMap(myCards)}
		cards = append(cards, card)
	}
}

func getIntersectionCount(winningCards, myCards map[int]bool) int {
	intersectionCount := 0
	for myNumber := range myCards {
		_, has_key := winningCards[myNumber]
		if has_key {
			intersectionCount++
		}
	}
	return intersectionCount
}

func solveForPartOne() int {
	ans := 0
	for _, card := range cards {
		intersectionCount := getIntersectionCount(card.winningCards, card.myCards)
		if intersectionCount > 0 {
			ans += (1 << (intersectionCount - 1))
		}
	}
	return ans
}

func solveForPartTwo() int {
	ans := 0
	cardsCount := make([]int, len(cards))
	for i := 0; i < len(cardsCount); i++ {
		cardsCount[i] = 1
	}
	for cardId, card := range cards {
		intersectionCount := getIntersectionCount(card.winningCards, card.myCards)
		for i := 1; i <= intersectionCount; i++ {
			cardsCount[cardId+i] += cardsCount[cardId]
		}
	}
	for _, v := range cardsCount {
		ans += v
	}
	return ans
}

func main() {
	parseFileToCards("./input_2.txt")
	// fmt.Println(solveForPartOne())
	fmt.Println(solveForPartTwo())
}
