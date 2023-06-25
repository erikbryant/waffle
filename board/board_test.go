package board

import (
	"testing"
)

func TestSize(t *testing.T) {
	testCases := []struct {
		size int
	}{
		{0},
		{1},
		{5},
		{7},
	}

	for _, testCase := range testCases {
		waffle := New(testCase.size)
		size := waffle.Size()
		if size != testCase.size {
			t.Errorf("ERROR: Expected %d got %d", testCase.size, size)
		}
	}
}

func TestNew(t *testing.T) {
	waffle := New(4)

	for row := 0; row < waffle.Size(); row++ {
		for col := 0; col < waffle.Size(); col++ {
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

	waffle := New(5)

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

func TestParse(t *testing.T) {
	testCases := []struct {
		serial     string
		expectSize int
		row        int
		col        int
		expectL    rune
		expectC    rune
	}{
		{"fboueg i ulsoomg e loemna/gwwggw w wgygyyw y wgyywg", 5, 4, 4, 'a', Green},
		{"eifdstal i p apertislt e e senithvte m t ueuedrra/yygygwyw w w wgwgggwgw w y wgwgggwgw w y wywgygww", 7, 1, 2, 'i', White},
	}

	for _, testCase := range testCases {
		waffle := Parse(testCase.serial)

		size := waffle.Size()
		if size != testCase.expectSize {
			t.Errorf("ERROR: Expected %d got %d", testCase.expectSize, size)
		}

		letter, color := waffle.Get(testCase.row, testCase.col)
		if letter != testCase.expectL {
			t.Errorf("ERROR: For (%d, %d) letter expected '%c' got '%c'", testCase.row, testCase.col, testCase.expectL, letter)
		}
		if color != testCase.expectC {
			t.Errorf("ERROR: For (%d, %d) color expected '%c' got '%c'", testCase.row, testCase.col, testCase.expectC, color)
		}
	}
}

func TestSerialize(t *testing.T) {
	testCases := []struct {
		serial string
	}{
		{"fboueg i ulsoomg e loemna/gwwggw w wgygyyw y wgyywg"},
		{"eifdstal i p apertislt e e senithvte m t ueuedrra/yygygwyw w w wgwgggwgw w y wgwgggwgw w y wywgygww"},
	}

	for _, testCase := range testCases {
		waffle := Parse(testCase.serial)
		serial := waffle.Serialize()
		if serial != testCase.serial {
			t.Errorf("ERROR: Expected %s got %s", testCase.serial, serial)
		}
	}
}
