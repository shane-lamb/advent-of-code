package main

import (
	"fmt"
	"github.com/samber/lo"
	"os"
	"strconv"
	"strings"
)

func getInputLines() []string {
	inputPath := "/Users/shane/source/advent-of-code-2022/problems/day4/input.txt"
	inputBytes, _ := os.ReadFile(inputPath)
	inputString := string(inputBytes)
	inputLines := strings.Split(inputString, "\n")
	return lo.Filter(inputLines, func(line string, _ int) bool {
		return line != "" // remove blank lines
	})
}

type assignment struct {
	start int
	end   int
}

func main() {
	lines := getInputLines()

	pairs := lo.Map(lines, func(line string, _ int) []assignment {
		assignmentStrings := strings.Split(line, ",")
		return lo.Map(assignmentStrings, func(assignmentString string, _ int) assignment {
			return toAssignment(assignmentString)
		})
	})

	// Part 1
	overlappingPairsCount := lo.CountBy(pairs, func(pair []assignment) bool {
		return doesCompletelyOverlap(pair[0], pair[1])
	})
	fmt.Printf("Number of completely overlapping pairs: %v\n", overlappingPairsCount)

	// Part 2
	partiallyOverlappingPairsCount := lo.CountBy(pairs, func(pair []assignment) bool {
		return doesOverlapAtAll(pair[0], pair[1])
	})
	fmt.Printf("Number of partially overlapping pairs: %v\n", partiallyOverlappingPairsCount)
}

func doesCompletelyOverlap(a1 assignment, a2 assignment) bool {
	return (a1.start >= a2.start && a1.end <= a2.end) || (a2.start >= a1.start && a2.end <= a1.end)
}

func doesOverlapAtAll(a1 assignment, a2 assignment) bool {
	return !(a1.start > a2.end || a2.start > a1.end)
}

func toAssignment(assignmentString string) assignment {
	parts := strings.Split(assignmentString, "-")
	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])
	return assignment{start, end}
}
