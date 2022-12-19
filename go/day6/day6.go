package main

import (
	"fmt"
	"github.com/samber/lo"
	"os"
	"strings"
)

func getInputLines() []string {
	inputPath := "/Users/shane/source/advent-of-code-2022/problems/day6/input.txt"
	inputBytes, _ := os.ReadFile(inputPath)
	inputString := string(inputBytes)
	inputLines := strings.Split(inputString, "\n")
	return lo.Filter(inputLines, func(line string, _ int) bool {
		return line != "" // remove blank lines
	})
}

func main() {
	lines := getInputLines()
	line := lines[0]
	chars := strings.Split(line, "")

	// Part 1
	startOfPackerPosition := getFirstMarkerPosition(chars, 4)
	fmt.Printf("First start-of-packet marker position: %v\n", startOfPackerPosition)

	// Part 2
	messagePosition := getFirstMarkerPosition(chars, 14)
	fmt.Printf("First message marker position: %v\n", messagePosition)
}

func getFirstMarkerPosition(chars []string, distinctCharCount int) int {
	var previousChars []string
	for index, char := range chars {
		previousChars = append(previousChars, char)
		if len(previousChars) >= distinctCharCount {
			lastChars := previousChars[len(previousChars)-distinctCharCount:]
			if len(lo.Uniq(lastChars)) == distinctCharCount {
				return index + 1
			}
		}
	}
	return -1
}
