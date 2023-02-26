package main

import (
	"fmt"
	"github.com/samber/lo"
	"os"
	"strconv"
	"strings"
)

func getInputLines() []string {
	inputPath := "/Users/shane/source/advent-of-code-2022/problems/day8/input.txt"
	inputBytes, _ := os.ReadFile(inputPath)
	inputString := string(inputBytes)
	inputLines := strings.Split(inputString, "\n")
	return lo.Filter(inputLines, func(line string, _ int) bool {
		return line != "" // remove blank lines
	})
}

type Tree struct {
	row    int
	col    int
	height int
}

func main() {
	lines := getInputLines()

	treeRows := lo.Map(lines, func(line string, row int) []Tree {
		chars := strings.Split(line, "")
		return lo.Map(chars, func(char string, col int) Tree {
			height, _ := strconv.Atoi(char)
			return Tree{row, col, height}
		})
	})
	treeCols := lo.Map(treeRows[0], func(_ Tree, col int) []Tree {
		var colTrees []Tree
		for row := 0; row < len(treeRows); row += 1 {
			colTrees = append(colTrees, treeRows[row][col])
		}
		return colTrees
	})

	// Part 1
	allLines := append(treeRows, treeCols...)
	reversedLines := lo.Map(allLines, func(treeLine []Tree, _ int) []Tree {
		return lo.Reverse(append([]Tree{}, treeLine...))
	})
	allLines = append(allLines, reversedLines...)
	visibleTreeDuplicates := lo.FlatMap(allLines, func(treeLine []Tree, _ int) []Tree {
		return getVisibleTrees(treeLine)
	})
	visibleTrees := lo.Uniq(visibleTreeDuplicates)
	fmt.Printf("Number of visible trees: %v\n", len(visibleTrees))

	// Part 2
	allScores := lo.Map(lo.Flatten(treeRows), func(tree Tree, _ int) int {
		return getScore(tree, treeRows, treeCols)
	})
	fmt.Printf("Highest score of any tree: %v\n", lo.Max(allScores))
}

func getScore(tree Tree, treeRows [][]Tree, treeCols [][]Tree) int {
	treeRow := treeRows[tree.row]
	treeCol := treeCols[tree.col]

	left := lo.Reverse(append([]Tree{}, treeRow[:tree.col]...))
	right := treeRow[tree.col+1:]
	up := lo.Reverse(append([]Tree{}, treeCol[:tree.row]...))
	down := treeCol[tree.row+1:]
	allDirections := [][]Tree{left, right, up, down}

	scores := lo.Map(allDirections, func(trees []Tree, _ int) int {
		return getDirectionalScore(trees, tree.height)
	})
	totalScore := lo.Reduce(scores, func(agg int, item int, _ int) int {
		return agg * item
	}, 1)
	return totalScore
}

func getDirectionalScore(treesInDirection []Tree, height int) int {
	for index, tree := range treesInDirection {
		if tree.height >= height {
			return index + 1
		}
	}
	return len(treesInDirection)
}

func getVisibleTrees(treeLine []Tree) []Tree {
	biggestHeight := -1
	var visibleTrees []Tree
	for _, tree := range treeLine {
		if tree.height > biggestHeight {
			visibleTrees = append(visibleTrees, tree)
			biggestHeight = tree.height
		}
	}
	return visibleTrees
}
