package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadLines(path string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("can't read file %s because of error: %s", path, err))
	}
	return strings.Split(string(data), "\n")
}

func LineScore(line string) int {
	lp, rp := -1, -1
	chars := strings.Split(line, "")
	for _, char := range chars {
		num, err := strconv.Atoi(char)
		if err == nil {
			if lp == -1 {
				lp = num
			}
			rp = num
		}
	}
	return lp*10 + rp
}

func LineReplacer(line string) string {
	mapper := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "4",
		"five":  "5e",
		"six":   "6",
		"seven": "7n",
		"eight": "e8t",
		"nine":  "9e",
	}
	for k, v := range mapper {
		line = strings.ReplaceAll(line, k, v)
	}
	return line
}

func main() {
	lines := ReadLines("./input_second.txt")
	ans := 0
	for _, line := range lines {
		line = LineReplacer(line)
		ans += LineScore(line)
	}
	fmt.Println(ans)
}
