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
		{"scgoln.n.dindeer.i.uffare/gwwwgg.w.yyggwyw.y.wgwyyg", true}, // 002
		{"mmkoye.u.iaomerr.o.pahcln/gwywgw.y.wyggwyw.w.wgwwwg", true}, // 100
		{"lieirs.w.riipese.n.rcouye/gwywgw.y.wygggyw.w.wgwywg", true}, // 200
		{"agdrml.r.ueianeu.l.oeibrr/gwywgw.y.wwygywg.w.ggwywg", true}, // 300
		{"bsmcye.s.eaiarsl.l.adeeks/gwywgy.y.yywgwyw.y.wgwgwg", true}, // 400
		{"daieoe.e.couvowr.g.glnene/gwywgy.g.yywgwyw.w.wgwywg", true}, // 500
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
