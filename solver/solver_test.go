package solver

import (
	"github.com/erikbryant/waffle/board"
	"testing"
)

func TestSolve(t *testing.T) {
	testCases := []struct {
		serial string
	}{
		// Deluxe Waffles
		{"eifdstal.i.p.apertislt.e.e.senithvte.m.t.ueuedrra/yygygwyw.w.w.wgwgggwgw.w.y.wgwgggwgw.w.y.wywgygww"}, // 056

		// Daily Waffles
		{"mmkoye.u.iaomerr.o.pahcln/gwywgw.y.wyggwyw.w.wgwwwg"}, // 100
		{"daieoe.e.couvowr.g.glnene/gwywgy.g.yywgwyw.w.wgwywg"}, // 500
	}

	for _, testCase := range testCases {
		waffle := board.Parse(testCase.serial)
		s := New(waffle)
		solved := s.Solve()
		if !solved {
			t.Errorf("ERROR: Failed to solve %s", testCase.serial)
		}
	}
}
