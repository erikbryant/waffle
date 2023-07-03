package pathfinder

import (
	"testing"
)

func TestFindDouble(t *testing.T) {
	testCases := []struct {
		swappable []Swappable
		have      rune
		want      rune
		expected  int
	}{
		{[]Swappable{}, 'a', 'b', -999},
		{[]Swappable{{1, 1, 'b', 'a'}}, 'a', 'b', 0},
		{[]Swappable{{0, 1, 'b', 'c'}, {1, 1, 'b', 'a'}}, 'a', 'b', 1},
		{[]Swappable{{0, 1, 'b', 'c'}, {0, 1, 'c', 'a'}}, 'a', 'b', -999},
	}

	for _, testCase := range testCases {
		index := findDouble(testCase.want, testCase.have, testCase.swappable)
		if index != testCase.expected {
			t.Errorf("ERROR: For %v expected %d got %d", testCase.swappable, testCase.expected, index)
		}
	}
}

func TestFindSingle(t *testing.T) {
	testCases := []struct {
		swappable []Swappable
		want      rune
		expected  int
	}{
		{[]Swappable{}, 'b', -999},
		{[]Swappable{{1, 1, 'b', 'a'}}, 'b', 0},
		{[]Swappable{{0, 1, 'b', 'c'}, {1, 1, 'b', 'a'}}, 'b', 0},
		{[]Swappable{{0, 1, 'd', 'c'}, {1, 1, 'b', 'a'}}, 'b', 1},
		{[]Swappable{{0, 1, 'x', 'c'}, {0, 1, 'c', 'a'}}, 'b', -999},
	}

	for _, testCase := range testCases {
		index := findSingle(testCase.want, testCase.swappable)
		if index != testCase.expected {
			t.Errorf("ERROR: For %v expected %d got %d", testCase.swappable, testCase.expected, index)
		}
	}
}
