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
		{"fboueg.i.ulsoomg.e.loemna/gwwggw.w.wgygyyw.y.wgyywg", true}, // 001
		{"eqebla.m.eupirel.n.mdlwal/ggywgw.w.ywygwww.g.wgyywg", true}, // 509
		{"tuaehl.r.emrdcnu.i.heoeby/gwgygw.w.wyygwww.g.wgywyg", true}, // 513
		{"bexkrd.c.aemarih.k.geasat/gywygy.w.ywygyww.g.wgwywg", true}, // 514
		{"smkupm.w.nknbeui.e.rgaiey/gyywgw.w.ywygwyw.y.wgwyyg", true}, // 515

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
