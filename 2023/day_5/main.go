package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Why no math.Min() function for int types :(
func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

// Why no math.Max() function for int types :(
func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

type SeedRange struct {
	start, end int
}

type Range struct {
	destinationStart, sourceStart, sourceEnd, length int
}

type Ranges = []Range

var seeds []SeedRange
var allMaps []Ranges

func strToNum(numStr string) int {
	num, _ := strconv.Atoi(numStr)
	return num
}

func strToNumArray(numStr string) []int {
	var nums []int
	for _, num := range strings.Fields(strings.TrimSpace(numStr)) {
		nums = append(nums, strToNum(num))
	}
	return nums
}

func parseRange(line string) Range {
	t := strings.Fields(line)
	destinationStart := strToNum(t[0])
	sourceStart := strToNum(t[1])
	length := strToNum(t[2])
	sourceEnd := sourceStart + length - 1
	return Range{destinationStart, sourceStart, sourceEnd, length}
}

func parseFile(path string) {
	data, _ := os.ReadFile(path)
	lines := strings.Split(string(data), "\n")

	seedsInput := strToNumArray(lines[0][6:])
	for i := 0; i < len(seedsInput); i += 2 {
		seeds = append(seeds, SeedRange{seedsInput[i], seedsInput[i] + seedsInput[i+1] - 1})
	}
	currentSection := Ranges{}
	for i := 3; i < len(lines); i++ {
		if lines[i] == "" {
			// a map section has ended.
			i += 2
			allMaps = append(allMaps, currentSection)
			currentSection = Ranges{}
			if i >= len(lines) {
				break
			}
		}
		currentSection = append(currentSection, parseRange(lines[i]))
	}
	allMaps = append(allMaps, currentSection)

	// Critical for the code implementation (i.e. applyMap function)
	for _, curMap := range allMaps {
		sort.Slice(curMap, func(i, j int) bool {
			return curMap[i].sourceStart < curMap[j].sourceStart
		})
	}
}

func applyMap(seedRange SeedRange, curMap Ranges) []SeedRange {
	var seedRanges []SeedRange

	// x1 is the left pointer of the seed range, processed so far. with each
	// iteration of map range, it will only increase, as the map ranges are sorted.
	x1 := seedRange.start
	x2 := seedRange.end

	// the ranges are sorted in the current map
	for _, rng := range curMap {
		y1, y2 := rng.sourceStart, rng.sourceEnd
		destStart := rng.destinationStart

		os := max(x1, y1) // os = overlap start
		oe := min(x2, y2) // oe = overlap end

		if os <= oe {
			// left portion to overlap, keep this intact.
			// as no future map range can overlap this.
			if x1 < os {
				seedRanges = append(seedRanges, SeedRange{x1, os - 1})
			}
			// process the overlap portion.
			seedRanges = append(seedRanges, SeedRange{os - y1 + destStart, oe - y1 + destStart})
			// uppdate the value of x1 on the basis of seed range portion that
			// has been covered so far.
			if oe < x2 {
				x1 = oe + 1
			} else {
				// whole seed range has been covered.
				x1 = math.MaxInt
				break
			}
		}
	}
	// if some seed range is left, even after processing all the map ranges.
	if x1 <= x2 {
		seedRanges = append(seedRanges, SeedRange{x1, x2})
	}
	return seedRanges
}

func solve() int {
	ans := math.MaxInt
	currentSeeds := seeds
	for _, curMap := range allMaps {
		var newSeeds []SeedRange
		for _, seed := range currentSeeds {
			tempSeedRanges := applyMap(seed, curMap)
			for _, tempSeedRange := range tempSeedRanges {
				newSeeds = append(newSeeds, tempSeedRange)
			}
		}
		currentSeeds = newSeeds
	}
	for _, currentSeed := range currentSeeds {
		ans = min(ans, currentSeed.start)
	}
	return ans
}

func main() {
	parseFile("./input_2.txt")
	fmt.Println(solve())
}
