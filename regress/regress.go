package main

import (
	"flag"
	"fmt"
	"github.com/erikbryant/waffle/board"
	"github.com/erikbryant/waffle/pathfinder"
	"github.com/erikbryant/waffle/solver"
	"log"
	"os"
	"runtime/pprof"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	testCases5 = []string{
		// Daily Waffles
		"fboueg i ulsoomg e loemna/gwwggw w wgygyyw y wgyywg", // 001
		"scgoln n dindeer i uffare/gwwwgg w yyggwyw y wgwyyg", // 002
		"speeda t ptocirn e mpeiey/gggwgw w yywgwyw g wgyywg", // 003
		"ndeeye e ltraeck a idnsks/gygygy y wwwgwyy w ygwywg", // 004
		"crmvpe r glaivye n belouy/ggwwgw w yywgwyw w wgwywg", // 005
		"mmkoye u iaomerr o pahcln/gwywgw y wyggwyw w wgwwwg", // 100
		"lieirs w riipese n rcouye/gwywgw y wygggyw w wgwywg", // 200
		"agdrml r ueianeu l oeibrr/gwywgw y wwygywg w ggwywg", // 300
		"bsmcye s eaiarsl l adeeks/gwywgy y yywgwyw y wgwgwg", // 400
		"daieoe e couvowr g glnene/gwywgy g yywgwyw w wgwywg", // 500
		"eqebla m eupirel n mdlwal/ggywgw w ywygwww g wgyywg", // 509
		"bnueet p moeioaz o sluwer/gwywgw g wywgwyw y wgywyg", // 510
		"weoakm e hocberm i nprerl/gwwygw y wgwgyyw w ygyywg", // 511
		"fiaclt u eatolhu t ltarae/gwywgw y wyygyyw y wgwywg", // 512
		"tuaehl r emrdcnu i heoeby/gwgygw w wyygwww g wgywyg", // 513
		"bexkrd c aemarih k geasat/gywygy w ywygyww g wgwywg", // 514
		"smkupm w nknbeui e rgaiey/gyywgw w ywygwyw y wgwyyg", // 515
		// "cetrlu i haalcrt n epncih/gyyygw y wwwgwww y wgwywg", // 516
		"feeonn w rewdinl g edolly/gwyygw w wgygwyw w ygwywg", // 517
		"mulnsa i ilrboue o tfuily/gyyygw w wywgwyw y wgwywg", // 518
		"puiweu o toeieeq t hrsaor/gyywgw w yywgwyw y wgwwgg", // 519
		"airdne v irtneeo m icvnae/gyyygw w wyygyyw y wgywyg", // 520
		"damnyg g ererame n ikuled/gwgwgw y yyygyww w ygwwyg", // 521
		"tgheea i osreemt a htigrn/gwywgw w wwgggww y wgyyyg", // 522
		"snilna r eueiblc a lmwigh/gwyygy w gwwggyw y wgwwwg", // 523
	}
	testCases7 = []string{
		// Deluxe Waffles
		"eifdstal i p apertislt e e senithvte m t ueuedrra/yygygwyw w w wgwgggwgw w y wgwgggwgw w y wywgygww", // 056
		"deoitnai e i oaamcnrcd o n bdtsraida s e iglsdeel/yygwgwyy w w wgygwgwgw g g wgwgygygw y w ywwgygww", // 057
	}
)

func TestSolve(testCases []string) {
	total := 0
	count := 0
	for _, testCase := range testCases {
		waffle := board.Parse(testCase)
		s := solver.New(waffle)
		if s.Solve() {
			path := pathfinder.New(s)
			path.Find()
			count++
			total += path.PathLen()
			fmt.Println("Average steps:", float64(total)/float64(count))
		} else {
			fmt.Println("Unable to solve:", testCase)
			s.Print()
			fmt.Println()
		}
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

	TestSolve(testCases7)
	fmt.Println()
	TestSolve(testCases5)
}
