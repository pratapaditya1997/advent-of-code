package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseStr(numStr string) float64 {
	numStr = strings.Split(numStr, ":")[1]
	numStr = strings.TrimSpace(numStr)

	tempStr := ""
	for _, char := range numStr {
		if char == ' ' {
			continue
		}
		tempStr += string(char)
	}

	val, _ := strconv.ParseFloat(tempStr, 64)
	return val
}

func parseFile(path string) (float64, float64) {
	data, _ := os.ReadFile(path)
	lines := strings.Split(string(data), "\n")
	time := parseStr(lines[0])
	distance := parseStr(lines[1])
	return time, distance
}

func quadraticSolver(a, b, c float64) (float64, float64) {
	sq := math.Sqrt(b*b - 4*a*c)
	root1 := (-b + sq) / (2 * a)
	root2 := (-b - sq) / (2 * a)
	return root1, root2
}

func main() {
	time, distance := parseFile("./input.txt")
	root1, root2 := quadraticSolver(-1, time, -distance)
	ans := int(math.Floor(root2) - math.Ceil(root1) + 1)
	fmt.Println(ans)
}
