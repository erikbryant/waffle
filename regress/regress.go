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
		"socopr t ceatsen t amfeoy/gyyggy y wywgwww w ygwywg", // 006
		"onwigr o nrrtcra v rtxpee/gyyygy y wywgwyw w wgwywg", // 007
		"slropa m orodcht r ontreh/gwywgw w gyygwww w wgygwg", // 008
		"tjmilo r agoailn l kneaia/gwwggg w ywggwyw w wgwyyg", // 009
		"bnebkn e avlidll n vemalt/gyywgw w wyggwyy w wgwywg", // 010
		"cldara h dxopeul i vreeey/gwwwgw w ywwgwyw w wgwywg", // 011
		// "arcxet b ezmonje n aezodt/gwwwgw w yywggyw g wgwywg", // 012 <-- Ambiguous
		"dggoam e gioinwn h tllelt/gwgygw w gwwgwww w wgwwwg", // 013
		"siatno y odhdver h etrmoe/gyyygy y wwwgwyw w wgwywg", // 014
		"yiyltu e vharakl c isocpt/gwwwgw y wwwgwwy y ggwwgg", // 015
		"sdrder i oipdepf l igiray/gwywgy w wywgwww w ygggwg", // 016
		"boedts n ahoagya e rwalel/ggywgw w ywwgwwg w wgwwgg", // 017
		"clrial i aurpewe h ayzrcd/gwgygw w ywwgwyw w wgwwwg", // 018
		"eothss y olrluoo x ctgsit/gyyygy y wwwgwgw w ygwywg", // 019
		"kaayka s elultat n idctey/ggyygy y wwwgyyw w ygwywg", // 020
		"vuemtr r eeetgio e reovrt/gwyygy y wyggwww g ygwwyg", // 021
		"pxamen t urrbrlo w odiabl/gwwwgw w wwwgwww w ggwgwg", // 022
		"idsvrc n nlmoria r aegfee/gwwwgw w ywwggyy g wgwywg", // 023
		"frjodn r teihhic i udttey/gyyygy y wgygwww w ygyywg", // 024
		"setnei o hoetteb r ndrlay/gwwggw w wgygyww y wgywwg", // 025
		"elcutr i ahxourl m atcghe/gwwwgw w ywwgwyw w wgwywg", // 026
		"fainlu e oseorve e esrdty/gyyygy y wwygyyy w ygwywg", // 027
		"mldaht o rnotial e utnuty/gwwwgw w yywgwyw w wgwywg", // 028
		"aeaayo a lelrvls c rretet/gwywgw g wywgwww w wggwwg", // 029
		"luuhet e apgraut e olsmny/gwyggw w wgygwyw w wgywwg", // 030
		"mmkoye u iaomerr o pahcln/gwywgw y wyggwyw w wgwwwg", // 100
		"lieirs w riipese n rcouye/gwywgw y wygggyw w wgwywg", // 200
		"agdrml r ueianeu l oeibrr/gwywgw y wwygywg w ggwywg", // 300
		"bsmcye s eaiarsl l adeeks/gwywgy y yywgwyw y wgwgwg", // 400
		"miidcr g aognone m alirie/ggywgw w wyygwyy y wgwgwg", // 401
		"mmubyc t eecnoui i dreuda/gwywgw w wyygyyw g wgwywg", // 402
		"tpalti g msourea i rriynd/gwwwgy y wwwgwgy w yggywg", // 403
		"thfrmr u ulcocat o etnteh/gwwwgg w gywgwyw w wgyyyg", // 404
		"txrite i iiramct a realpl/gwyggw w yyygwwy w ygwwyg", // 405
		"ttolbo i uiraego n spharo/gwywgw w wyygyyw w wgwywg", // 406
		"pealnr n ltosaro l aeonyy/gwwygw w wyygyyw y ygwwwg", // 407
		"idrare d gkliaot u elnnee/gwywgy y yywgwyw w wgwwwg", // 408
		"ssnotl r acataer n atcoxt/gwywgw y gyyggww w wgwywg", // 409
		"hplntt i ganisea r geeear/gwwwgy w yyygyyw w wgywyg", // 410
		"woeadt l reabulh r efmwll/ggywgw w wyygwyy y ygwwwg", // 411
		"pemtob n acvorgn e vhiyaa/gwwwgw w wyygyyw y wgwywg", // 412
		"peroeo o akndnnl o baenry/gwyygw y wywgwyw w ygwgwg", // 413
		"snaeem r etletio v henrvr/gwwwgy y yywgwyw y wgwwwg", // 414
		"vahmro c isiapan l arlery/gyywgy y wgwgwyw w wgwywg", // 415
		"bbrtma o froimuw s pmoeoo/gwwwgw w wyygyyw y wgwywg", // 416
		"gosadr a tenisra v gelnue/ggywgg w wwygwwy w wgyywg", // 417
		"crnsgi e emsarui s nplgns/gwywgy y yywgwyw w wgwwwg", // 418
		"hamele e uoamivt o olyetr/gwwggy w wywgwwy y wgwyyg", // 419
		"sielkr t uuadoin n epcwml/gwywgw w wyygyyw y wgwwwg", // 420
		"daieoe e couvowr g glnene/gwywgy g yywgwyw w wgwywg", // 500
		"seostl l seohatn e upeecr/gyywgy w wywgyww y ygywwg", // 501
		"cmnsar i haochwi y etthot/gwywgw y wygggyw y wgwywg", // 502
		"wleerr o sapogts e arapry/gwwggw y yywgyww w ygwgwg", // 503
		"baarkk t lcaerxn p redola/gwgwgy y ywygyww g wgwwwg", // 504
		"paldes g eenugtu t noroer/gwwwgy w yyygyyw w wgwygg", // 505
		"cmhett e aagaalp r rtaret/gwywgw y wywgwyw w wgyyyg", // 506
		"reuori e lodewcd e areiro/ggywgw w wyygyww g ygywwg", // 507
		"sokrrn e nnaigpl u aehgee/gwywgw y wyygyyw w wgwgwg", // 508
		"eqebla m eupirel n mdlwal/ggywgw w ywygwww g wgyywg", // 509
		"bnueet p moeioaz o sluwer/gwywgw g wywgwyw y wgywyg", // 510
		"weoakm e hocberm i nprerl/gwwygw y wgwgyyw w ygyywg", // 511
		"fiaclt u eatolhu t ltarae/gwywgw y wyygyyw y wgwywg", // 512
		"tuaehl r emrdcnu i heoeby/gwgygw w wyygwww g wgywyg", // 513
		"bexkrd c aemarih k geasat/gywygy w ywygyww g wgwywg", // 514
		"smkupm w nknbeui e rgaiey/gyywgw w ywygwyw y wgwyyg", // 515
		"cetrlu i haalcrt n epncih/gyyygw y wwwgwww y wgwywg", // 516
		"feeonn w rewdinl g edolly/gwyygw w wgygwyw w ygwywg", // 517
		"mulnsa i ilrboue o tfuily/gyyygw w wywgwyw y wgwywg", // 518
		"puiweu o toeieeq t hrsaor/gyywgw w yywgwyw y wgwwgg", // 519
		"airdne v irtneeo m icvnae/gyyygw w wyygyyw y wgywyg", // 520
		"damnyg g ererame n ikuled/gwgwgw y yyygyww w ygwwyg", // 521
		"tgheea i osreemt a htigrn/gwywgw w wwgggww y wgyyyg", // 522
		"snilna r eueiblc a lmwigh/gwyygy w gwwggyw y wgwwwg", // 523
		"ctcrms n drnoaal h rmaihy/gwwwgw y wyygyyw w wgwywg", // 524
		"frbnhl o ooemeat o dmiiuy/gwywgw y gywggyw w wgwywg", // 525
		"adluyk i iclnbio a eelery/gwgwgw g wywgwyw w wgyyyg", // 526
		"sctrpm e tatobio u sftauy/ggywgw w ywygwww g wgyywg", // 527
		"wikwnt n nnoonhr j ehceaa/gwgwgy g yywgwyw w wgwywg", // 528
		"cecuke r oocdorn t aaamha/gwyygw y wyygwgw w ygwwwg", // 529
		"belbtr u aatmorn e utewrl/gwywgy y yywgwyw g wgwwwg", // 530
		"caceki e lltsaoe i nmhlnl/gwywgy y yywgyww w ygwwwg", // 531
		"jreeto e sdwieed n eltruy/gwywgw w wyygyyw y wgwywg", // 532
		"asalwp e wropjea c helagt/gwyygw y wywgwgy w ygwwwg", // 533
		"aselto e cmradrl g awnvlt/gwgwgy y yywgwyy w ygwywg", // 534
		"wihnfh s rzetino r rtfere/gyywgw w gwwgyyw y wgwyyg", // 535
	}
	testCases7 = []string{
		// Deluxe Waffles
		"taenergt r e perpsersd m x iaretaler m n vscsrevd/ywgygyww y w wgygygwgy w w ygygwgwgy w w wwwgygwy", // 001
		"abnserou i l ialmeeist f b rarryaseu m y sellsdam/yygygwww w y wgygwgygw w w ygygwgwgw w w wyygygwy", // 002
		"dsvsosgu r e rmdbrtors t c issaiines u b oyieueoe/wygwgwyw y w ygwgwgygy y w wgygwgwgy w w wywgwgww", // 003
		"ruorglcu p u gorttaaen h t rsephouto t r uetsearh/yygwgwwy w w wgygwgwgw w w ygwgwgygw w y wyygygyy", // 004
		"muinoeum c m siaputtsn r c naneiorec t s mrdscesw/wygygwww w y wgwgygwgy g g ygwgwgwgw w w wyygygyw", // 005
		"cespnbsa n e saeoatdds a n irilrapem l e nstdgeee/yygwgyyy w w ygygwgygw w w wgwgwgwgy w w yywgwgwy", // 006
		"inmoacrm e e iaigolncc h e dstaeerdw s e lemeitna/yygygyyw w w wgygwgygy w w ygwgygwgw w w wywgwgwy", // 007
		"hiawtcsn u o irydaaete r c esteltdeo d i srrsarpl/yygwgyyw w y wgwgygwgy g g wgwgwgwgw y w wwwgwgwy", // 008
		"ceninrai t c uaatretsa c e stfrmines p i uaelired/yygygwyy w w wgwgggwgw w y ygwgggwgw w w wwwgygww", // 009
		"autasedr n i onfrrafap d y eooieiynn i v senrpgoa/ywgwgwyy w w wgwgwgwgw g g ygygwgwgw w w wywgwgyy", // 010
		"eifdstal i p apertislt e e senithvte m t ueuedrra/yygygwyw w w wgwgggwgw w y wgwgggwgw w y wywgygww", // 056
		"deoitnai e i oaamcnrcd o n bdtsraida s e iglsdeel/yygwgwyy w w wgygwgwgw g g wgwgygygw y w ywwgygww", // 057
		"ussputbi r e mftnrsper b t leduiasee t c yelnhetl/wwgygyyw y w wgwgwgwgw g g wgwgwgwgy w w yyygwgyw", // 058
		"uotsntid n x delpvadno p r senieeete a e ghlgetis/ywgwgyww y w ygygwgwgw g g wgygwgwgw w y wywgygyy", // 059
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
			fmt.Printf("Game: %3d Steps: %3d Average: %3.2f\n", count, path.PathLen(), float64(total)/float64(count))
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
