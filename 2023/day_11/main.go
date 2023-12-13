package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Pair struct {
	x, y int
}

var rows, cols int
var seenRow, seenCol []bool
var galaxies []Pair

const multiplier = int64(999999)

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func parseFile(path string) {
	data, _ := os.ReadFile(path)
	lines := strings.Split(string(data), "\n")

	rows = len(lines)
	seenRow = make([]bool, rows)

	cols = len(lines[0])
	seenCol = make([]bool, cols)

	galaxies = []Pair{}
	for i, line := range lines {
		line = strings.TrimSpace(line)
		allChars := strings.Split(line, "")
		for j, char := range allChars {
			if char == "#" {
				galaxies = append(galaxies, Pair{i, j})
				seenRow[i] = true
				seenCol[j] = true
			}
		}
	}
}

func getEmptyRows(i, j int) int {
	if i > j {
		i, j = j, i
	}
	empty := 0
	for r := i + 1; r < j; r++ {
		if seenRow[r] == false {
			empty++
		}
	}
	return empty
}

func getEmptyCols(i, j int) int {
	if i > j {
		i, j = j, i
	}
	empty := 0
	for c := i + 1; c < j; c++ {
		if seenCol[c] == false {
			empty++
		}
	}
	return empty
}

func main() {
	startTime := time.Now()
	parseFile("./input.txt")

	ans := int64(0)
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			emptyRows := int64(getEmptyRows(galaxies[i].x, galaxies[j].x))
			emptyCols := int64(getEmptyCols(galaxies[i].y, galaxies[j].y))

			ans += int64(abs(galaxies[i].x-galaxies[j].x)) + (emptyRows * multiplier)
			ans += int64(abs(galaxies[i].y-galaxies[j].y)) + (emptyCols * multiplier)
		}
	}
	fmt.Println(ans)
	endTime := time.Now()
	fmt.Printf("Execution Time: %v\n", endTime.Sub(startTime))
}
