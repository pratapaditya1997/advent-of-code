package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Cell struct {
	x, y int
}

// direction enum
const (
	N = iota
	E
	W
	S
)

type Ray struct {
	cell Cell
	d    int // direction enum
}

var m, n int
var grid [][]rune
var seen map[Ray]bool

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func parseFile(path string) {
	data, _ := os.ReadFile(path)
	lines := strings.Split(string(data), "\n")

	m = len(lines)
	grid = make([][]rune, m)
	for i, line := range lines {
		n = len(line)
		grid[i] = make([]rune, n)
		for j, char := range line {
			grid[i][j] = char
		}
	}
}

func getNextDirs(ray Ray) []int {
	cell := ray.cell
	char := grid[cell.x][cell.y]

	switch char {
	case '/':
		switch ray.d {
		case N:
			return []int{E}
		case E:
			return []int{N}
		case W:
			return []int{S}
		case S:
			return []int{W}
		}
	case '\\':
		switch ray.d {
		case N:
			return []int{W}
		case E:
			return []int{S}
		case W:
			return []int{N}
		case S:
			return []int{E}
		}
	case '|':
		switch ray.d {
		case N:
			return []int{N}
		case E:
			return []int{N, S}
		case W:
			return []int{N, S}
		case S:
			return []int{S}
		}
	case '-':
		switch ray.d {
		case N:
			return []int{E, W}
		case E:
			return []int{E}
		case W:
			return []int{W}
		case S:
			return []int{E, W}
		}
	}
	return []int{ray.d}
}

func getNextRays(cell Cell, nextDirs []int) []Ray {
	nextRays := []Ray{}
	for _, dirs := range nextDirs {
		nx, ny := cell.x, cell.y
		switch dirs {
		case N:
			nx, ny = cell.x-1, cell.y
		case E:
			nx, ny = cell.x, cell.y+1
		case W:
			nx, ny = cell.x, cell.y-1
		case S:
			nx, ny = cell.x+1, cell.y
		}
		nextRays = append(nextRays, Ray{Cell{nx, ny}, dirs})
	}
	return nextRays
}

func cellOutOfBounds(cell Cell) bool {
	if cell.x < 0 || cell.x >= m {
		return true
	}
	if cell.y < 0 || cell.y >= n {
		return true
	}
	return false
}

func dfs(ray Ray) {
	if cellOutOfBounds(ray.cell) {
		return
	}
	if seen[ray] {
		return
	}
	seen[ray] = true

	nextDirs := getNextDirs(ray)
	nextRays := getNextRays(ray.cell, nextDirs)
	for _, nextRay := range nextRays {
		dfs(nextRay)
	}
}

func solve(ray Ray) int {
	seen = map[Ray]bool{}
	dfs(ray)

	cellSeen := map[Cell]bool{}
	for k, v := range seen {
		if v {
			cellSeen[k.cell] = true
		}
	}

	ans := 0
	for _, v := range cellSeen {
		if v {
			ans++
		}
	}
	return ans
}

func main() {
	startTime := time.Now()

	parseFile("./input.txt")
	ans, curAns := 0, 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				curAns = max(curAns, solve(Ray{Cell{i, j}, S}))
			}
			if i == m-1 {
				curAns = max(curAns, solve(Ray{Cell{i, j}, N}))
			}
			if j == 0 {
				curAns = max(curAns, solve(Ray{Cell{i, j}, E}))
			}
			if j == n-1 {
				curAns = max(curAns, solve(Ray{Cell{i, j}, W}))
			}
			ans = max(ans, curAns)
		}
	}
	fmt.Println(ans)

	endTime := time.Now()
	fmt.Printf("Execution Time: %v\n", endTime.Sub(startTime))
}
