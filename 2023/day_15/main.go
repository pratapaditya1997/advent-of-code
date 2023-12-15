package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Value struct {
	focalLength, counter int
}

type Box map[string]Value

var steps []string
var boxes []Box

func parseFile(path string) {
	data, _ := os.ReadFile(path)
	line := strings.TrimSpace(strings.Split(string(data), "\n")[0])
	steps = strings.Split(line, ",")
}

func hash(label string) int {
	ans := 0
	for _, char := range label {
		ans = ((ans + int(char)) * 17) % 256
	}
	return ans
}

func remove(boxIdx int, label string) {
	_, hasKey := boxes[boxIdx][label]
	if !hasKey {
		return
	}
	delete(boxes[boxIdx], label)
}

func add(boxIdx int, label string, focalLength int, counter int) {
	_, hasKey := boxes[boxIdx][label]
	if hasKey {
		value := boxes[boxIdx][label]
		value.focalLength = focalLength
		boxes[boxIdx][label] = value
	} else {
		boxes[boxIdx][label] = Value{focalLength: focalLength, counter: counter}
	}
}

func solve() int {
	ans := 0
	for i := 0; i < 256; i++ {
		curBoxAns := 0
		boxSlice := []Value{}
		for _, value := range boxes[i] {
			boxSlice = append(boxSlice, value)
		}
		sort.Slice(boxSlice, func(i, j int) bool {
			return boxSlice[i].counter < boxSlice[j].counter
		})
		for idx, value := range boxSlice {
			curBoxAns += ((i + 1) * (idx + 1) * value.focalLength)
		}
		ans += curBoxAns
	}
	return ans
}

func main() {
	parseFile("./input.txt")
	boxes = make([]Box, 256)
	for i := 0; i < 256; i++ {
		boxes[i] = Box{}
	}

	counter := 0
	for _, step := range steps {
		if step[len(step)-1] == '-' {
			// remove
			label := step[0 : len(step)-1]
			boxIdx := hash(label)
			remove(boxIdx, label)
		} else {
			// add
			label := step[0 : len(step)-2]
			focalLength, _ := strconv.Atoi(string(step[len(step)-1]))
			boxIdx := hash(label)
			add(boxIdx, label, focalLength, counter)
		}
		counter++
	}

	fmt.Println(solve())
}
