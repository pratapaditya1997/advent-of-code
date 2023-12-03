package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ------------------------------------
// Common Stuff used in both the parts
// ------------------------------------
var grid = [][]string{}
var m, n int

var dx = []int{-1, -1, -1, 1, 1, 1, 0, 0}
var dy = []int{-1, 0, 1, -1, 0, 1, -1, 1}

type NumRange struct {
	row, leftPointer, rightPointer int
}

func (nr *NumRange) numericalValue() int {
	return convertToInt(nr.row, nr.leftPointer, nr.rightPointer)
}

func convertToInt(i, lp, rp int) int {
	ans := 0
	for j := lp; j <= rp; j++ {
		num, _ := strconv.Atoi(grid[i][j])
		ans = ans*10 + num
	}
	return ans
}

func parseFileToGrid(path string) {
	data, _ := os.ReadFile(path)
	lines := strings.Split(string(data), "\n")
	m = len(lines)
	n = len(lines[0])

	temp_grid := make([][]string, m)
	for i := 0; i < m; i++ {
		line := lines[i]
		chars := strings.Split(line, "")
		temp_grid[i] = make([]string, n)
		for j := 0; j < n; j++ {
			temp_grid[i][j] = chars[j]
		}
	}

	grid = make([][]string, m+2)
	for i := 0; i < m+2; i++ {
		grid[i] = make([]string, n+2)
		for j := 0; j < n+2; j++ {
			if i == 0 || i == m+1 || j == 0 || j == n+1 {
				grid[i][j] = "."
			} else {
				grid[i][j] = temp_grid[i-1][j-1]
			}
		}
	}
}

func isNumber(char string) bool {
	_, err := strconv.Atoi(char)
	return err == nil
}

// returns a slice of all the NumRanges present in the grid
func buildNumRanges() []NumRange {
	var numRanges = []NumRange{}
	for i := 1; i < m; i++ {
		j := 0
		lp, rp := j, j
		for j < n {
			if isNumber(grid[i][j]) == true {
				lp = j
				for j < n {
					if isNumber(grid[i][j]) == false {
						rp = j - 1
						break
					}
					j++
				}
				numRanges = append(numRanges, NumRange{i, lp, rp})
			}
			j++
		}
	}
	return numRanges
}

// ------------------------------------
// Stuff used in only the Part 1
// ------------------------------------

func isCharSpecial(i, j int) bool {
	_, err := strconv.Atoi(grid[i][j])
	if err == nil {
		// char is digit, hence not special
		return false
	}
	return grid[i][j] != "."
}

// returns true if any surrounding character is special.
func isCharValid(i, j int) bool {
	for idx := 0; idx < len(dx); idx++ {
		if isCharSpecial(i+dx[idx], j+dy[idx]) {
			return true
		}
	}
	return false
}

func isRangeValid(i, lp, rp int) bool {
	for j := lp; j <= rp; j++ {
		if isCharValid(i, j) == true {
			return true
		}
	}
	return false
}

func solveForPartOne() int {
	ans := 0
	numRanges := buildNumRanges()
	for _, numRange := range numRanges {
		if isRangeValid(numRange.row, numRange.leftPointer, numRange.rightPointer) {
			ans += int(numRange.numericalValue())
		}
	}
	return ans
}

// ------------------------------------
// Stuff used in only the Part 2
// ------------------------------------

// a given char at index (i,j) will exactly be in one range
func charInNumRange(i int, j int, numRanges []NumRange) (bool, NumRange) {
	for _, numRange := range numRanges {
		if i == numRange.row && j >= numRange.leftPointer && j <= numRange.rightPointer {
			return true, numRange
		}
	}
	return false, NumRange{}
}

func numbersAroundGear(i int, j int, numRanges []NumRange) []NumRange {
	var ranges = map[NumRange]bool{}
	var uniqueRanges = []NumRange{}
	for idx := 0; idx < len(dx); idx++ {
		res, validRange := charInNumRange(i+dx[idx], j+dy[idx], numRanges)
		if res {
			ranges[validRange] = true
		}
	}
	for k := range ranges {
		uniqueRanges = append(uniqueRanges, k)
	}
	return uniqueRanges
}

func solveForPartTwo() int {
	ans := 0
	numRanges := buildNumRanges()

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == "*" {
				validRanges := numbersAroundGear(i, j, numRanges)
				if len(validRanges) == 2 {
					val1 := validRanges[0].numericalValue()
					val2 := validRanges[1].numericalValue()
					ans += val1 * val2
				}
			}
		}
	}
	return ans
}

// ------------------------------------
// Main function
// ------------------------------------

func main() {
	parseFileToGrid("./input.txt")
	m = len(grid)
	n = len(grid[0])
	// fmt.Println(solveForPartOne())
	fmt.Println(solveForPartTwo())
}
