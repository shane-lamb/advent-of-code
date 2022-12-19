package main

import (
	stack "advent/common"
	"fmt"
	"github.com/samber/lo"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getInputString() string {
	inputPath := "/Users/shane/source/advent-of-code-2022/problems/day5/input.txt"
	inputBytes, _ := os.ReadFile(inputPath)
	return string(inputBytes)
}

type instruction struct {
	count int
	from  int
	to    int
}

func main() {
	sections := strings.Split(getInputString(), "\n\n")

	instructions := parseInstructions(sections[1])

	// Part 1
	beginStacks := parseState(sections[0])
	endStacks := runInstructionsOld(instructions, beginStacks)
	containersAtTop := lo.Map(endStacks, func(stack *stack.Stack, _ int) string {
		container := stack.Peek()
		return container.(string)
	})
	fmt.Printf("Part 1, Containers at top of stacks: %v\n", strings.Join(containersAtTop, ""))

	// Part 2
	beginStacks = parseState(sections[0])
	endStacks = runInstructionsNew(instructions, beginStacks)
	containersAtTop = lo.Map(endStacks, func(stack *stack.Stack, _ int) string {
		container := stack.Peek()
		return container.(string)
	})
	fmt.Printf("Part 2, Containers at top of stacks: %v\n", strings.Join(containersAtTop, ""))
}

func runInstructionsOld(instructions []instruction, stacks []*stack.Stack) []*stack.Stack {
	for _, instruction := range instructions {
		for i := 0; i < instruction.count; i++ {
			popped := stacks[instruction.from].Pop()
			stacks[instruction.to].Push(popped)
		}
	}
	return stacks
}

func runInstructionsNew(instructions []instruction, stacks []*stack.Stack) []*stack.Stack {
	for _, instruction := range instructions {
		var temp = make([]interface{}, 0)
		for i := 0; i < instruction.count; i++ {
			temp = append(temp, stacks[instruction.from].Pop())
		}
		for i := len(temp) - 1; i > -1; i-- {
			stacks[instruction.to].Push(temp[i])
		}
	}
	return stacks
}

func parseInstructions(text string) []instruction {
	lines := strings.Split(text, "\n")
	withoutLastLine := lines[:len(lines)-1]

	// example instruction: move 1 from 5 to 2
	reg := regexp.MustCompile(`(move )|( from )|( to )`)
	return lo.Map(withoutLastLine, func(instructionText string, _ int) instruction {
		split := reg.Split(instructionText, -1)
		count, _ := strconv.Atoi(split[1])
		from, _ := strconv.Atoi(split[2])
		to, _ := strconv.Atoi(split[3])
		return instruction{count, from - 1, to - 1}
	})
}

func parseState(text string) []*stack.Stack {
	lines := strings.Split(text, "\n")
	withoutLastLine := lines[:len(lines)-1]

	grid := lo.Map(withoutLastLine, func(line string, _ int) []string {
		line = strings.ReplaceAll(line, "    ", " [ ]")
		line = strings.ReplaceAll(line, " ", "")
		line = strings.ReplaceAll(line, "[", "")
		columns := strings.Split(line, "]")
		return columns[:len(columns)-1]
	})
	flippedGrid := lo.Reverse(grid)
	stacks := lo.Map(grid[0], func(_ string, _ int) *stack.Stack {
		return stack.New()
	})
	for _, row := range flippedGrid {
		for index, part := range row {
			if part != "" {
				stacks[index].Push(part)
			}
		}
	}
	return stacks
}
