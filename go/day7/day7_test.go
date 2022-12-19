package main

import (
	"github.com/samber/lo"
	"testing"
)

func TestCase(t *testing.T) {
	consoleLines := []string{
		"$ cd /",
		"$ ls",
		"dir a",
		"dir b",
		"1000 afile1",
		"$ cd a",
		"$ ls",
		"2000 afile2.cs",
		"dir c",
		"dir d",
		"$ cd c",
		"$ ls",
		"4000 afile3",
		"$ cd ..",
		"$ cd d",
		"$ ls",
		"$ cd ..",
		"$ cd ..",
		"$ cd b",
		"$ ls",
		"dir e",
		"8000 afile4.cs",
		"1000 afile5",
		"dir f",
		"$ cd e",
		"$ ls",
		"16000 afile6",
		"$ cd ..",
		"$ cd f",
		"$ ls",
		"32000 afile7",
		"500 afile8",
	}
	rootDir := parseConsoleOutput(consoleLines)
	totalSize, sizes := getDirectorySizes(rootDir)

	folderF := 32000 + 500
	folderE := 16000
	folderB := folderE + folderF + 8000 + 1000
	folderD := 0
	folderC := 4000
	folderA := folderD + folderC + 2000
	folderRoot := folderA + folderB + 1000

	expectedTotal := folderRoot + folderA + folderB + folderC + folderD + folderE + folderF
	actualTotal := totalSize + lo.Sum(sizes)

	if actualTotal != expectedTotal {
		t.Errorf("got %v, wanted %v", actualTotal, expectedTotal)
	}
}
