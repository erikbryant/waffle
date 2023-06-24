package board

import (
	"testing"
)

func TestWidthHeight(t *testing.T) {
	testCases := []struct {
		width  int
		height int
	}{
		{0, 0},
		{1, 0},
		{0, 1},
		{1, 1},
		{5, 5},
		{7, 7},
	}

	for _, testCase := range testCases {
		waffle := New(testCase.width, testCase.height)
		val := waffle.Width()
		if val != testCase.width {
			t.Errorf("ERROR: For width expected %d got %d", testCase.width, val)
		}
		val = waffle.Height()
		if val != testCase.height {
			t.Errorf("ERROR: For height expected %d got %d", testCase.height, val)
		}
	}
}

func TestNew(t *testing.T) {
	waffle := New(3, 4)

	for row := 0; row < waffle.Height(); row++ {
		for col := 0; col < waffle.Width(); col++ {
			letter, color := waffle.Get(row, col)
			if letter != Empty {
				t.Errorf("ERROR: For letter expected '%c' got '%c'", Empty, letter)
			}
			if color != Empty {
				t.Errorf("ERROR: For color expected '%c' got '%c'", Empty, color)
			}
		}
	}
}

func TestSetGet(t *testing.T) {
	testCases := []struct {
		row     int
		col     int
		letter  rune
		color   rune
		expectL rune
		expectC rune
	}{
		{0, 0, 'e', Green, 'e', Green},
		{0, 0, 'f', Yellow, 'f', Yellow},
		{1, 1, 'p', Yellow, Empty, Empty},
		{3, 3, 'l', Yellow, Empty, Empty},
		{4, 4, 'k', White, 'k', White},
		{4, 4, 'z', Green, 'z', Green},
	}

	waffle := New(5, 5)

	for _, testCase := range testCases {
		waffle.Set(testCase.row, testCase.col, testCase.letter, testCase.color)
		letter, color := waffle.Get(testCase.row, testCase.col)
		if letter != testCase.expectL {
			t.Errorf("ERROR: For (%d, %d) letter expected '%c' got '%c'", testCase.row, testCase.col, testCase.expectL, letter)
		}
		if color != testCase.expectC {
			t.Errorf("ERROR: For (%d, %d) color expected '%c' got '%c'", testCase.row, testCase.col, testCase.expectC, color)
		}
	}
}
