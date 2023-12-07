package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var cardValues = map[string]int{
	"J": 0,
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"Q": 10,
	"K": 11,
	"A": 12,
}

const (
	UNKNOWN = iota
	HIGH_CARD
	ONE_PAIR
	TWO_PAIR
	THREE_OF_KIND
	FULL_HOUSE
	FOUR_OF_KIND
	FIVE_OF_KIND
)

type hand struct {
	cards    []string
	bid      int
	handType int // based on above enums
}

func sortHands(hands []hand) []hand {
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].handType != hands[j].handType {
			return hands[i].handType < hands[j].handType
		}
		for k := 0; k < 5; k++ {
			if cardValues[hands[i].cards[k]] != cardValues[hands[j].cards[k]] {
				return cardValues[hands[i].cards[k]] < cardValues[hands[j].cards[k]]
			}
		}
		return true
	})
	return hands
}

func calculateHandType(cards []string) int {
	count := map[string]int{}
	jCount := 0
	for _, card := range cards {
		if card == "J" {
			jCount++
			continue
		}
		count[card]++
	}

	maxKey := ""
	maxCount := 0

	for key, val := range count {
		if val >= maxCount {
			maxCount = val
			maxKey = key
		}
	}

	count[maxKey] += jCount

	switch len(count) {
	case 1:
		return FIVE_OF_KIND
	case 2:
		for _, val := range count {
			if val == 4 {
				return FOUR_OF_KIND
			}
		}
		return FULL_HOUSE
	case 3:
		for _, val := range count {
			if val == 3 {
				return THREE_OF_KIND
			}
		}
		return TWO_PAIR
	case 4:
		return ONE_PAIR
	case 5:
		return HIGH_CARD
	}
	return HIGH_CARD
}

func calculateHandTypes(hands []hand) []hand {
	var updatedHands []hand
	for _, curHand := range hands {
		curHand.handType = calculateHandType(curHand.cards)
		updatedHands = append(updatedHands, curHand)
	}
	return updatedHands
}

func parseFile(path string) []hand {
	data, _ := os.ReadFile(path)
	lines := strings.Split(string(data), "\n")

	var hands []hand
	for _, line := range lines {
		cardsStr := strings.Fields(line)[0]
		bidStr := strings.Fields(line)[1]
		cards := strings.Split(cardsStr, "")
		bid, _ := strconv.Atoi(bidStr)
		hands = append(hands, hand{cards, bid, UNKNOWN})
	}
	return hands
}

func main() {
	startTime := time.Now()
	hands := parseFile("input.txt")
	valuedHands := calculateHandTypes(hands)
	sortedHands := sortHands(valuedHands)

	ans := 0
	for idx, curHand := range sortedHands {
		ans += (idx + 1) * curHand.bid
	}
	fmt.Println(ans)
	endTime := time.Now()
	fmt.Printf("Execution Time: %v\n", endTime.Sub(startTime))
}
