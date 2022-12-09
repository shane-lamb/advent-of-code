package main

import (
	"fmt"
	"github.com/samber/lo"
	"os"
	"strings"
)

func getInputLines() []string {
	inputPath := "/Users/shane/source/advent-of-code-2022/problems/day3/input.txt"
	inputBytes, _ := os.ReadFile(inputPath)
	inputString := string(inputBytes)
	inputLines := strings.Split(inputString, "\n")
	return lo.Filter(inputLines, func(line string, _ int) bool {
		return line != "" // remove blank lines
	})
}

func main() {
	rucksacks := getInputLines()

	// Part 1
	sharedItemPriorities := lo.Map(rucksacks, func(rucksack string, _ int) int {
		compartment1, compartment2 := getCompartments(rucksack)
		sharedItem := getSharedItem([]string{compartment1, compartment2})
		return getPriority(sharedItem)
	})
	sumOfItemPriorities := lo.Sum(sharedItemPriorities)
	fmt.Printf("Sum of item priorities: %v\n", sumOfItemPriorities)

	// Part 2
	elfGroups := lo.Chunk(rucksacks, 3)
	badgePriorities := lo.Map(elfGroups, func(rucksacks []string, _ int) int {
		sharedItem := getSharedItem(rucksacks)
		return getPriority(sharedItem)
	})
	sumOfBadgePriorities := lo.Sum(badgePriorities)
	fmt.Printf("Sum of badge priorities: %v\n", sumOfBadgePriorities)
}

func getPriority(item string) int {
	// A to Z = 65 to 90
	// a to z = 97 to 122
	code := int(item[0])
	if code > 90 {
		return code - 96
	}
	return code - 38 // will make A = 27
}

func getSharedItem(groups []string) string {
	commonItems := lo.Reduce(groups, func(agg []string, group string, _ int) []string {
		items := strings.Split(group, "")
		return lo.Intersect(items, agg)
	}, strings.Split(groups[0], ""))
	return commonItems[0]
}

func getCompartments(rucksack string) (compartment1 string, compartment2 string) {
	compartmentSize := len(rucksack) / 2
	return rucksack[:compartmentSize], rucksack[compartmentSize:]
}
