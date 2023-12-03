package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// needed only for part one
var maxPerColor = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func ReadLines(path string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("can't read file %s because of error: %s", path, err))
	}
	return strings.Split(string(data), "\n")
}

func BreakGameIntoSets(game string) []string {
	game = strings.TrimSpace(game)
	setsCombined := strings.Split(game, ":")[1]
	sets := strings.Split(setsCombined, ";")
	return sets
}

func BreakDownConfig(config string) (int, string) {
	config = strings.TrimSpace(config)
	number, _ := strconv.Atoi(strings.Split(config, " ")[0])
	color := strings.TrimSpace(strings.Split(config, " ")[1])
	return number, color
}

// needed only for part one
func IsConfigPossible(set string) bool {
	set = strings.TrimSpace(set)
	configs := strings.Split(set, ",")
	for _, config := range configs {
		number, color := BreakDownConfig(config)
		if number > maxPerColor[color] {
			return false
		}
	}
	return true
}

// needed only for part one
func SolveForPartOne(id int, sets []string) int {
	isSetPossible := true
	for _, set := range sets {
		if IsConfigPossible(set) == false {
			isSetPossible = false
		}
		if isSetPossible == false {
			break
		}
	}
	if isSetPossible == true {
		return id + 1
	}
	return 0
}

// needed only for part two
// This returns a game' score.
func SolveForPartTwo(sets []string) int {
	minNeeded := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for _, set := range sets {
		set = strings.TrimSpace(set)
		configs := strings.Split(set, ",")
		for _, config := range configs {
			number, color := BreakDownConfig(config)
			if number > minNeeded[color] {
				minNeeded[color] = number
			}
		}
	}

	gameScore := 1
	for _, v := range minNeeded {
		gameScore *= v
	}
	return gameScore
}

func main() {
	ans := 0
	lines := ReadLines("./input_2.txt")
	for _, line := range lines {
		sets := BreakGameIntoSets(line)
		// ans += SolveForPartOne(id, sets)
		ans += SolveForPartTwo(sets)
	}
	fmt.Println(ans)
}
