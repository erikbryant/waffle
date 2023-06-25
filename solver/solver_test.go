package solver

import (
	"github.com/erikbryant/waffle/board"
	"testing"
)

func TestSolve(t *testing.T) {
	testCases := []struct {
		serial   string
		solvable bool
	}{
		// Deluxe Waffles
		{"eifdstal.i.p.apertislt.e.e.senithvte.m.t.ueuedrra/yygygwyw.w.w.wgwgggwgw.w.y.wgwgggwgw.w.y.wywgygww", true}, // 056

		// Daily Waffles
		{"mmkoye.u.iaomerr.o.pahcln/gwywgw.y.wyggwyw.w.wgwwwg", true}, // 100
		{"daieoe.e.couvowr.g.glnene/gwywgy.g.yywgwyw.w.wgwywg", true}, // 500

		// Unsolvable Waffles
		{"abcd.e.fghi/wwww.w.wwww", false},
	}

	for _, testCase := range testCases {
		waffle := board.Parse(testCase.serial)
		s := New(waffle)
		solved := s.Solve()
		if solved != testCase.solvable {
			t.Errorf("ERROR: For %s expected %t got %t", testCase.serial, testCase.solvable, solved)
		}
	}
}
