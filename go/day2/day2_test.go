package main

import "testing"

var turnTestCases = []struct {
	turn     string
	expected int
}{
	{"A X", 3},
	{"A Y", 4},
	{"A Z", 8},
	{"B X", 1},
	{"B Y", 5},
	{"B Z", 9},
	{"C X", 2},
	{"C Y", 6},
	{"C Z", 7},
}

func TestCalculateTotalScoreNew(t *testing.T) {
	for _, testCase := range turnTestCases {
		t.Run(testCase.turn, func(t *testing.T) {
			result := calculateTotalScoreNew(testCase.turn)
			if result != testCase.expected {
				t.Errorf("%s: got %v, wanted %v",
					testCase.turn, result, testCase.expected)
			}
		})
	}
}
