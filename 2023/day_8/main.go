package main

import (
	"fmt"
	"os"
	"strings"
)

type Node struct {
	name                  string
	leftChild, rightChild int
}

var graph []Node
var instructions string
var nameToId map[string]int

func parseLineToNode(line string) Node {
	node := Node{}
	node.name = line[0:3]
	node.leftChild = nameToId[line[7:10]]
	node.rightChild = nameToId[line[12:15]]
	return node
}

func parseFile(path string) {
	data, _ := os.ReadFile(path)
	lines := strings.Split(string(data), "\n")

	instructions = strings.TrimSpace(lines[0])

	idCounter := 0
	nameToId = map[string]int{}
	for i := 2; i < len(lines); i++ {
		_, hasKey := nameToId[lines[i][0:3]]
		if hasKey == false {
			nameToId[lines[i][0:3]] = idCounter
			idCounter++
		}
	}

	graph = []Node{}
	for i := 2; i < len(lines); i++ {
		graph = append(graph, parseLineToNode(lines[i]))
	}
}

func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(nums []int64) int64 {
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		ans = ((nums[i] * ans) / gcd(nums[i], ans))
	}
	return ans
}

func main() {
	parseFile("input.txt")

	aNodes := []int{}
	for idx, node := range graph {
		if node.name[2] == 'A' {
			aNodes = append(aNodes, idx)
		}
	}

	stepsList := []int64{}
	for _, idx := range aNodes {
		curStep := 0
		curNode := idx
		ok := true

		for ok {
			if graph[curNode].name[2] == 'Z' {
				break
			}
			curInstruction := instructions[curStep%len(instructions)]
			if curInstruction == 'L' {
				curNode = graph[curNode].leftChild
			} else {
				curNode = graph[curNode].rightChild
			}
			curStep++
		}
		stepsList = append(stepsList, int64(curStep))
	}
	fmt.Println(lcm(stepsList))
}
