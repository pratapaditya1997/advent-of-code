package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseLine(line string) []int {
	nums := []int{}
	for _, numStr := range strings.Fields(line) {
		num, _ := strconv.Atoi(numStr)
		nums = append(nums, num)
	}
	return nums
}

func reverse(nums []int) []int {
	reversed := []int{}
	for i := len(nums) - 1; i >= 0; i-- {
		reversed = append(reversed, nums[i])
	}
	return reversed
}

func isAllZero(nums []int) bool {
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}
	return true
}

// returns single integer, which is the extrapolated number for the give nums
// array.
func extrapolateList(nums []int) int {
	if isAllZero(nums) {
		return 0
	}
	diffNums := []int{}
	for i := 1; i < len(nums); i++ {
		diffNums = append(diffNums, nums[i]-nums[i-1])
	}
	return nums[len(nums)-1] + extrapolateList(diffNums)
}

func main() {
	data, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(data), "\n")

	ans := 0
	for _, line := range lines {
		nums := parseLine(line)
		// ans += extrapolateList(nums) // solves part one
		ans += extrapolateList(reverse(nums)) // solves part two
	}
	fmt.Println(ans)
}
