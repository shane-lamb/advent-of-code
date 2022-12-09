package main

import (
	"fmt"
	"github.com/samber/lo"
	"os"
	"strings"
)

func getInput() string {
	inputPath := "/Users/shane/source/advent-of-code-2022/problems/day2/input.txt"
	inputBytes, _ := os.ReadFile(inputPath)
	return string(inputBytes)
}

func main() {
	lines := strings.Split(getInput(), "\n")
	turns := lo.Filter(lines, func(line string, _ int) bool {
		return line != "" // remove blank lines
	})

	// Part 1
	wrongTurnScores := lo.Map(turns, func(turn string, _ int) int {
		return calculateTotalScoreOld(turn)
	})
	wrongTotalScore := lo.Sum(wrongTurnScores)
	fmt.Printf("If everything goes to plan, my score will be: %v\n", wrongTotalScore)

	// Part 2
	actualTurnScores := lo.Map(turns, func(turn string, _ int) int {
		return calculateTotalScoreNew(turn)
	})
	actualTotalScore := lo.Sum(actualTurnScores)
	fmt.Printf("Correction: If everything goes to plan, my score will be: %v\n", actualTotalScore)
}

func calculateTotalScoreNew(turn string) int {
	symbols := strings.Split(turn, " ")
	opponentMove := getMove(symbols[0])
	outcome := getOutcome(symbols[1])
	outcomeScore := (outcome + 1) * 3
	return outcomeScore + calculateShapeScoreNew(opponentMove, outcome)
}

func calculateShapeScoreNew(opponentMove int, outcome int) int {
	myMove := mod(opponentMove+outcome, 3)
	return myMove + 1
}

// equivalent to % in python
func mod(x, d int) int {
	x = x % d
	if x >= 0 {
		return x
	}
	if d < 0 {
		return x - d
	}
	return x + d
}

func getOutcome(outcomeSymbol string) int {
	switch outcomeSymbol {
	case "X":
		return -1
	case "Y":
		return 0
	case "Z":
		return 1
	default:
		panic("Invalid input")
	}
}

func getMove(moveSymbol string) int {
	switch moveSymbol {
	case "A":
		return 0
	case "B":
		return 1
	case "C":
		return 2
	default:
		panic("Invalid input")
	}
}

func calculateTotalScoreOld(turn string) int {
	return calculateOutcomeScoreOld(turn) + calculateShapeScoreOld(turn)
}

func calculateOutcomeScoreOld(turn string) int {
	switch turn {
	case "A X":
		return 3
	case "A Y":
		return 6 // opponent has rock, I have paper
	case "B Y":
		return 3
	case "B Z":
		return 6 // opponent has paper, I have scissors
	case "C X":
		return 6 // opponent has scissors, I have rock
	case "C Z":
		return 3
	default:
		return 0
	}
}

func calculateShapeScoreOld(turn string) int {
	moves := strings.Split(turn, " ")
	switch moves[1] {
	case "X":
		return 1
	case "Y":
		return 2
	case "Z":
		return 3
	default:
		panic("Invalid input")
	}
}
