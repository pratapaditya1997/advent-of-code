package main

import (
	"fmt"
	"os"
	"strings"
)

type pair struct {
	i, j int
}

var inputGrid map[pair]rune
var m, n int
var source pair

var loop map[pair]bool // using as set

var possiblePipes = []string{"|", "-", "L", "J", "7", "F"}

func parseFile(path string) {
	data, _ := os.ReadFile(path)
	lines := strings.Split(string(data), "\n")

	m = len(lines)
	inputGrid = map[pair]rune{}
	for i := 0; i < m; i++ {
		lines[i] = strings.TrimSpace(lines[i])
		allChars := []rune(lines[i])
		n = len(allChars)
		for j := 0; j < n; j++ {
			if allChars[j] == 'S' {
				source.i = i
				source.j = j
			}
			inputGrid[pair{i, j}] = allChars[j]
		}
	}
}

// takes a cell(i,j) and returns possible next cells that can eb visited as per
// the current cell's pipe type.
func getNextCells(cell pair) []pair {
	nextCells := []pair{}

	north := inputGrid[pair{cell.i - 1, cell.j}]
	south := inputGrid[pair{cell.i + 1, cell.j}]
	east := inputGrid[pair{cell.i, cell.j + 1}]
	west := inputGrid[pair{cell.i, cell.j - 1}]

	switch string(inputGrid[pair{cell.i, cell.j}]) {
	case "|":
		if north == '|' || north == '7' || north == 'F' {
			nextCells = append(nextCells, pair{cell.i - 1, cell.j})
		}
		if south == '|' || south == 'J' || south == 'L' {
			nextCells = append(nextCells, pair{cell.i + 1, cell.j})
		}
	case "-":
		if west == '-' || west == 'L' || west == 'F' {
			nextCells = append(nextCells, pair{cell.i, cell.j - 1})
		}
		if east == '-' || east == '7' || east == 'J' {
			nextCells = append(nextCells, pair{cell.i, cell.j + 1})
		}
	case "L":
		if north == '|' || north == '7' || north == 'F' {
			nextCells = append(nextCells, pair{cell.i - 1, cell.j})
		}
		if east == '-' || east == '7' || east == 'J' {
			nextCells = append(nextCells, pair{cell.i, cell.j + 1})
		}
	case "J":
		if north == '|' || north == '7' || north == 'F' {
			nextCells = append(nextCells, pair{cell.i - 1, cell.j})
		}
		if west == '-' || west == 'L' || west == 'F' {
			nextCells = append(nextCells, pair{cell.i, cell.j - 1})
		}
	case "7":
		if west == '-' || west == 'L' || west == 'F' {
			nextCells = append(nextCells, pair{cell.i, cell.j - 1})
		}
		if south == '|' || south == 'J' || south == 'L' {
			nextCells = append(nextCells, pair{cell.i + 1, cell.j})
		}
	case "F":
		if south == '|' || south == 'J' || south == 'L' {
			nextCells = append(nextCells, pair{cell.i + 1, cell.j})
		}
		if east == '-' || east == '7' || east == 'J' {
			nextCells = append(nextCells, pair{cell.i, cell.j + 1})
		}
	case ".":
		return nextCells
	}
	return nextCells
}

func isOutOfBounds(cell pair) bool {
	if cell.i < 0 || cell.i >= m {
		return true
	}
	if cell.j < 0 || cell.j >= n {
		return true
	}
	return false
}

// solves for given instance of the inputGrid and returns max distance point
// in the loop. Uses BFS (Breadth First Search).
func solve() int {
	queue := make([]pair, 0)
	dist := map[pair]int{}
	loop = map[pair]bool{}

	queue = append(queue, pair{source.i, source.j})
	dist[pair{source.i, source.j}] = 0
	maxDist := 0
	loop[source] = true

	for len(queue) != 0 {
		curCell := queue[0]
		queue = queue[1:] // pop the current element

		nextCells := getNextCells(curCell)
		for _, nextCell := range nextCells {
			if isOutOfBounds(nextCell) {
				continue
			}
			_, alreadyVisited := dist[pair{nextCell.i, nextCell.j}]
			if alreadyVisited {
				continue
			}
			loop[nextCell] = true
			queue = append(queue, pair{nextCell.i, nextCell.j})
			dist[pair{nextCell.i, nextCell.j}] = dist[pair{curCell.i, curCell.j}] + 1

			if dist[pair{nextCell.i, nextCell.j}] > maxDist {
				maxDist = dist[pair{nextCell.i, nextCell.j}]
			}
		}
	}

	return maxDist
}

func main() {
	parseFile("./sample_input.txt")
	for _, pipe := range possiblePipes {
		inputGrid[pair{source.i, source.j}] = []rune(pipe)[0]

		curSourceNextCells := getNextCells(source)
		if len(curSourceNextCells) == 2 {
			break
		}
	}

	solve()
	ans := 0
	for i := 0; i < m; i++ {
		inside := 0
		for j := 0; j < n; j++ {
			cellVal := inputGrid[pair{i, j}]
			if cellVal == '|' || cellVal == 'L' || cellVal == 'J' {
				if loop[pair{i, j}] {
					inside = 1 - inside
				}
			}
			if loop[pair{i, j}] == false && inside == 1 {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
