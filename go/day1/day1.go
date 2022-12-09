package main

import (
	"fmt"
	"github.com/samber/lo"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getInput() string {
	inputPath := "/Users/shane/source/advent-of-code-2022/problems/day1/input.txt"
	inputBytes, _ := os.ReadFile(inputPath)
	return string(inputBytes)
}

func main() {
	// Part 1
	elves := strings.Split(getInput(), "\n\n")
	elvesTotalCalories := lo.Map(elves, func(elf string, _ int) int {
		lines := strings.Split(elf, "\n")
		return lo.SumBy(lines, func(line string) int {
			calories, _ := strconv.ParseInt(line, 10, 64)
			return int(calories)
		})
	})
	mostCaloriesForSingleElf := lo.Max(elvesTotalCalories)
	fmt.Printf("Most calories for a single elf: %v\n", mostCaloriesForSingleElf)

	// Part 2
	// Sort the elves in ascending order of total calories held
	sort.Ints(elvesTotalCalories)
	top3Elves := elvesTotalCalories[len(elvesTotalCalories)-3:]
	top3ElvesCalories := lo.Sum(top3Elves)
	fmt.Printf("Total calories for top 3 elves: %v\n", top3ElvesCalories)
}
