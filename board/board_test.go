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
