package main

import (
	"flag"
	"fmt"
	"github.com/erikbryant/waffle/board"
	"github.com/erikbryant/waffle/solver"
	"log"
	"os"
	"runtime/pprof"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
)

func TestSolve() {
	testCases := []string{
		// Deluxe Waffles
		"eifdstal.i.p.apertislt.e.e.senithvte.m.t.ueuedrra/yygygwyw.w.w.wgwgggwgw.w.y.wgwgggwgw.w.y.wywgygww", // 056

		// Daily Waffles
		"fboueg.i.ulsoomg.e.loemna/gwwggw.w.wgygyyw.y.wgyywg", // 001
		"scgoln.n.dindeer.i.uffare/gwwwgg.w.yyggwyw.y.wgwyyg", // 002
		"mmkoye.u.iaomerr.o.pahcln/gwywgw.y.wyggwyw.w.wgwwwg", // 100
		"lieirs.w.riipese.n.rcouye/gwywgw.y.wygggyw.w.wgwywg", // 200
		"agdrml.r.ueianeu.l.oeibrr/gwywgw.y.wwygywg.w.ggwywg", // 300
		"bsmcye.s.eaiarsl.l.adeeks/gwywgy.y.yywgwyw.y.wgwgwg", // 400
		"daieoe.e.couvowr.g.glnene/gwywgy.g.yywgwyw.w.wgwywg", // 500
		"eqebla.m.eupirel.n.mdlwal/ggywgw.w.ywygwww.g.wgyywg", // 509
		"tuaehl.r.emrdcnu.i.heoeby/gwgygw.w.wyygwww.g.wgywyg", // 513
		"bexkrd.c.aemarih.k.geasat/gywygy.w.ywygyww.g.wgwywg", // 514
		"smkupm.w.nknbeui.e.rgaiey/gyywgw.w.ywygwyw.y.wgwyyg", // 515
	}

	for _, testCase := range testCases {
		waffle := board.Parse(testCase)
		s := solver.New(waffle)
		s.Solve()
	}
}

func main() {
	fmt.Println("Welcome to waffle regression tests!")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	TestSolve()
}
