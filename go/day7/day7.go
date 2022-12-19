package main

import (
	stack "advent/common"
	"fmt"
	"github.com/samber/lo"
	"os"
	"strconv"
	"strings"
)

func getInputLines() []string {
	inputPath := "/Users/shane/source/advent-of-code-2022/problems/day7/input.txt"
	inputBytes, _ := os.ReadFile(inputPath)
	inputString := string(inputBytes)
	inputLines := strings.Split(inputString, "\n")
	return lo.Filter(inputLines, func(line string, _ int) bool {
		return line != "" // remove blank lines
	})
}

type file struct {
	name string
	size int
}

type dir struct {
	name       string
	childDirs  []*dir
	childFiles []file
}

func main() {
	promptLines := getInputLines()
	rootDir := parseConsoleOutput(promptLines)
	rootDirSize, dirSizes := getDirectorySizes(rootDir)

	// Part 1
	sumOfSizes := lo.SumBy(dirSizes, func(dirSize int) int {
		if dirSize > 100_000 {
			return 0
		}
		return dirSize
	})
	fmt.Printf("Sum of directory sizes of size 100,000 or less: %v\n", sumOfSizes)

	// Part 2
	totalDiskSpace := 70000000
	unusedSpaceNeeded := 30000000
	spaceUsedBeforeDelete := rootDirSize
	currentUnusedSpace := totalDiskSpace - spaceUsedBeforeDelete
	spaceToFree := unusedSpaceNeeded - currentUnusedSpace
	candidateDirSizes := lo.Filter(dirSizes, func(dirSize int, _ int) bool {
		return dirSize >= spaceToFree
	})
	dirSizeToDelete := lo.Min(candidateDirSizes)
	fmt.Printf("Size of directory to delete: %v\n", dirSizeToDelete)
}

func getDirectorySizes(thisDir *dir) (int, []int) {
	filesTotalSize := lo.SumBy(thisDir.childFiles, func(file file) int {
		return file.size
	})
	if thisDir.childDirs == nil {
		return filesTotalSize, []int{}
	}
	immediateChildDirSizes := lo.Map(thisDir.childDirs, func(childDir *dir, _ int) int {
		dir, _ := getDirectorySizes(childDir)
		return dir
	})
	immediateChildDirsTotalSize := lo.SumBy(immediateChildDirSizes, func(dirSize int) int {
		return dirSize
	})
	childDirSizes := lo.FlatMap(thisDir.childDirs, func(childDir *dir, _ int) []int {
		_, dirs := getDirectorySizes(childDir)
		return dirs
	})
	return filesTotalSize + immediateChildDirsTotalSize, append(childDirSizes, immediateChildDirSizes...)
}

func parseConsoleOutput(promptLines []string) *dir {
	dirStack := stack.New()
	for _, line := range promptLines {
		words := strings.Split(line, " ")

		switch words[0] {
		case "$":
			switch words[1] {
			case "ls":
				continue
			case "cd":
				changeDirectory(words[2], dirStack)
			default:
				panic(fmt.Sprintf("'%s' is not a recognised command!", words[1]))
			}
		case "dir":
			continue
		default:
			currentDir := dirStack.Peek().(*dir)
			addFile(currentDir, words)
		}
	}
	return dirStack.GetItems()[0].(*dir)
}

func addFile(currentDir *dir, words []string) {
	size, _ := strconv.Atoi(words[0])
	name := words[1]
	file := file{name, size}
	currentDir.childFiles = append(currentDir.childFiles, file)
}

func changeDirectory(dirName string, dirStack *stack.Stack) {
	if dirName == ".." {
		dirStack.Pop()
		return
	}
	childDir := dir{name: dirName}
	currentDir := dirStack.Peek()
	if currentDir != nil {
		currentDir.(*dir).childDirs = append(currentDir.(*dir).childDirs, &childDir)
	}
	dirStack.Push(&childDir)
}
