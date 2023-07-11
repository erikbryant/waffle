package main

import (
	"fmt"
	"github.com/erikbryant/waffle/board"
	"github.com/erikbryant/waffle/pathfinder"
	"github.com/erikbryant/waffle/solver"
)

type TestCase struct {
	serial string
	index  int
}

var (
	deluxeWaffles = []TestCase{
		{"taenergt r e perpsersd m x iaretaler m n vscsrevd/ywgygyww y w wgygygwgy w w ygygwgwgy w w wwwgygwy", 1},
		{"abnserou i l ialmeeist f b rarryaseu m y sellsdam/yygygwww w y wgygwgygw w w ygygwgwgw w w wyygygwy", 2},
		{"dsvsosgu r e rmdbrtors t c issaiines u b oyieueoe/wygwgwyw y w ygwgwgygy y w wgygwgwgy w w wywgwgww", 3},
		{"ruorglcu p u gorttaaen h t rsephouto t r uetsearh/yygwgwwy w w wgygwgwgw w w ygwgwgygw w y wyygygyy", 4},
		{"muinoeum c m siaputtsn r c naneiorec t s mrdscesw/wygygwww w y wgwgygwgy g g ygwgwgwgw w w wyygygyw", 5},
		{"cespnbsa n e saeoatdds a n irilrapem l e nstdgeee/yygwgyyy w w ygygwgygw w w wgwgwgwgy w w yywgwgwy", 6},
		{"inmoacrm e e iaigolncc h e dstaeerdw s e lemeitna/yygygyyw w w wgygwgygy w w ygwgygwgw w w wywgwgwy", 7},
		{"hiawtcsn u o irydaaete r c esteltdeo d i srrsarpl/yygwgyyw w y wgwgygwgy g g wgwgwgwgw y w wwwgwgwy", 8},
		{"ceninrai t c uaatretsa c e stfrmines p i uaelired/yygygwyy w w wgwgggwgw w y ygwgggwgw w w wwwgygww", 9},
		{"autasedr n i onfrrafap d y eooieiynn i v senrpgoa/ywgwgwyy w w wgwgwgwgw g g ygygwgwgw w w wywgwgyy", 10},
		{"pnrbevvt n g tldmbiegc e o ienasiaec u i lvetanxo/wygygwww y w ygwgggwgw w y wgwgggwgw w w yyygwgyw", 11},
		{"nesrfaut e g ssaahitne p l ndtpmuael u i oolrbeer/wwgwgyyw y w wgwgwgygy g g ygwgygwgw w w wywgwgwy", 12},
		{"asiipeon l m saecetrcs n e gdrsmiase m i noaritda/yygwgywy w w wgwgggwgw y w wgwgggwgw w w ywygwgyy", 13},
		{"twsruldr l r ferumiied e l elngrosku e a abeevmrg/ywgygwyw y y wgwgwgwgy w w ygwgygwgy w w ywygwgyw", 14},
		{"maoxeyvu o p sertvattv d s rerooicea r t yrnrtnbn/wwgwgwyw y w wgygwgygw g g ygwgwgwgw w w yyygygww", 15},
		{"eemoenom g a tceiubmrl s r tnrtoasec c a iptnneto/ywgygwww w w ygygwgygw g g wgygwgwgy w w wwygygyw", 16},
		{"rtoiamtg i n naredomel r r nsnaeivap n f dellhgen/ywgwgwyw w y ygwgwgygw g g wgygwgwgy w w wywgygww", 17},
		{"iiaiadme l s umtnsoanb o e mreskilre s e dndaehoa/ywgwgyyw w y ygwgggwgw w w wgygggwgw w w ywygygww", 18},
		{"eiitoael g n pcrptujed u n yclnteetb a a oimertlc/ywgygyyy y y wgygggwgw w w wgwgggwgy w w wywgwgwy", 19},
		{"thfeccun h t isistasni n i husitisgl g o ndigceiu/ywgygwww y w ygygggwgy y w ygwgggwgw w w wywgwgwy", 20},
		{"henreubq n i tmerluree e q aduuotado o d yieearnc/wwgygyww w w wgygwgwgw w y wgwgygwgy w w yyygwgwy", 21},
		{"tcvcsrou o r dbpneares o i itsueiett e g coiehkar/ywgwgyyw w w wgwgwgwgw w w wgygwgwgw y w wywgygwy", 22},
		{"ucpmonop r l tatategru y r eipflasea l k nrulreyn/wygwgyyy w w wgwgggwgw g g wgwgggwgw y w yywgwgyw", 23},
		{"besnrxat c v agerolral v i lntaeeyto c r eivtiese/yygwgwww w y wgwgygwgw g g ygygygwgy w w yywgwgww", 24},
		{"mirtieib p n latgllece a n apsteeetu i e xdnalsrn/ywgygyyw w w wgwgwgygy g g wgwgwgwgw w w yywgygww", 25},
		{"eifdstal i p apertislt e e senithvte m t ueuedrra/yygygwyw w w wgwgggwgw w y wgwgggwgw w y wywgygww", 56},
		{"deoitnai e i oaamcnrcd o n bdtsraida s e iglsdeel/yygwgwyy w w wgygwgwgw g g wgwgygygw y w ywwgygww", 57},
		{"ussputbi r e mftnrsper b t leduiasee t c yelnhetl/wwgygyyw y w wgwgwgwgw g g wgwgwgwgy w w yyygwgyw", 58},
		{"uotsntid n x delpvadno p r senieeete a e ghlgetis/ywgwgyww y w ygygwgwgw g g wgygwgwgw w y wywgygyy", 59},
	}
	dailyWaffles = []TestCase{
		{"fboueg i ulsoomg e loemna/gwwggw w wgygyyw y wgyywg", 1},
		{"scgoln n dindeer i uffare/gwwwgg w yyggwyw y wgwyyg", 2},
		{"speeda t ptocirn e mpeiey/gggwgw w yywgwyw g wgyywg", 3},
		{"ndeeye e ltraeck a idnsks/gygygy y wwwgwyy w ygwywg", 4},
		{"crmvpe r glaivye n belouy/ggwwgw w yywgwyw w wgwywg", 5},
		{"socopr t ceatsen t amfeoy/gyyggy y wywgwww w ygwywg", 6},
		{"onwigr o nrrtcra v rtxpee/gyyygy y wywgwyw w wgwywg", 7},
		{"slropa m orodcht r ontreh/gwywgw w gyygwww w wgygwg", 8},
		{"tjmilo r agoailn l kneaia/gwwggg w ywggwyw w wgwyyg", 9},
		{"bnebkn e avlidll n vemalt/gyywgw w wyggwyy w wgwywg", 10},
		{"cldara h dxopeul i vreeey/gwwwgw w ywwgwyw w wgwywg", 11},
		// {"arcxet b ezmonje n aezodt/gwwwgw w yywggyw g wgwywg", 12}, // <-- Ambiguous
		{"dggoam e gioinwn h tllelt/gwgygw w gwwgwww w wgwwwg", 13},
		{"siatno y odhdver h etrmoe/gyyygy y wwwgwyw w wgwywg", 14},
		{"yiyltu e vharakl c isocpt/gwwwgw y wwwgwwy y ggwwgg", 15},
		{"sdrder i oipdepf l igiray/gwywgy w wywgwww w ygggwg", 16},
		{"boedts n ahoagya e rwalel/ggywgw w ywwgwwg w wgwwgg", 17},
		{"clrial i aurpewe h ayzrcd/gwgygw w ywwgwyw w wgwwwg", 18},
		{"eothss y olrluoo x ctgsit/gyyygy y wwwgwgw w ygwywg", 19},
		{"kaayka s elultat n idctey/ggyygy y wwwgyyw w ygwywg", 20},
		{"vuemtr r eeetgio e reovrt/gwyygy y wyggwww g ygwwyg", 21},
		{"pxamen t urrbrlo w odiabl/gwwwgw w wwwgwww w ggwgwg", 22},
		{"idsvrc n nlmoria r aegfee/gwwwgw w ywwggyy g wgwywg", 23},
		{"frjodn r teihhic i udttey/gyyygy y wgygwww w ygyywg", 24},
		{"setnei o hoetteb r ndrlay/gwwggw w wgygyww y wgywwg", 25},
		{"elcutr i ahxourl m atcghe/gwwwgw w ywwgwyw w wgwywg", 26},
		{"fainlu e oseorve e esrdty/gyyygy y wwygyyy w ygwywg", 27},
		{"mldaht o rnotial e utnuty/gwwwgw w yywgwyw w wgwywg", 28},
		{"aeaayo a lelrvls c rretet/gwywgw g wywgwww w wggwwg", 29},
		{"luuhet e apgraut e olsmny/gwyggw w wgygwyw w wgywwg", 30},
		{"mmkoye u iaomerr o pahcln/gwywgw y wyggwyw w wgwwwg", 100},
		{"lieirs w riipese n rcouye/gwywgw y wygggyw w wgwywg", 200},
		{"agdrml r ueianeu l oeibrr/gwywgw y wwygywg w ggwywg", 300},
		{"bsmcye s eaiarsl l adeeks/gwywgy y yywgwyw y wgwgwg", 400},
		{"miidcr g aognone m alirie/ggywgw w wyygwyy y wgwgwg", 401},
		{"mmubyc t eecnoui i dreuda/gwywgw w wyygyyw g wgwywg", 402},
		{"tpalti g msourea i rriynd/gwwwgy y wwwgwgy w yggywg", 403},
		{"thfrmr u ulcocat o etnteh/gwwwgg w gywgwyw w wgyyyg", 404},
		{"txrite i iiramct a realpl/gwyggw w yyygwwy w ygwwyg", 405},
		{"ttolbo i uiraego n spharo/gwywgw w wyygyyw w wgwywg", 406},
		{"pealnr n ltosaro l aeonyy/gwwygw w wyygyyw y ygwwwg", 407},
		{"idrare d gkliaot u elnnee/gwywgy y yywgwyw w wgwwwg", 408},
		{"ssnotl r acataer n atcoxt/gwywgw y gyyggww w wgwywg", 409},
		{"hplntt i ganisea r geeear/gwwwgy w yyygyyw w wgywyg", 410},
		{"woeadt l reabulh r efmwll/ggywgw w wyygwyy y ygwwwg", 411},
		{"pemtob n acvorgn e vhiyaa/gwwwgw w wyygyyw y wgwywg", 412},
		{"peroeo o akndnnl o baenry/gwyygw y wywgwyw w ygwgwg", 413},
		{"snaeem r etletio v henrvr/gwwwgy y yywgwyw y wgwwwg", 414},
		{"vahmro c isiapan l arlery/gyywgy y wgwgwyw w wgwywg", 415},
		{"bbrtma o froimuw s pmoeoo/gwwwgw w wyygyyw y wgwywg", 416},
		{"gosadr a tenisra v gelnue/ggywgg w wwygwwy w wgyywg", 417},
		{"crnsgi e emsarui s nplgns/gwywgy y yywgwyw w wgwwwg", 418},
		{"hamele e uoamivt o olyetr/gwwggy w wywgwwy y wgwyyg", 419},
		{"sielkr t uuadoin n epcwml/gwywgw w wyygyyw y wgwwwg", 420},
		{"daieoe e couvowr g glnene/gwywgy g yywgwyw w wgwywg", 500},
		{"seostl l seohatn e upeecr/gyywgy w wywgyww y ygywwg", 501},
		{"cmnsar i haochwi y etthot/gwywgw y wygggyw y wgwywg", 502},
		{"wleerr o sapogts e arapry/gwwggw y yywgyww w ygwgwg", 503},
		{"baarkk t lcaerxn p redola/gwgwgy y ywygyww g wgwwwg", 504},
		{"paldes g eenugtu t noroer/gwwwgy w yyygyyw w wgwygg", 505},
		{"cmhett e aagaalp r rtaret/gwywgw y wywgwyw w wgyyyg", 506},
		{"reuori e lodewcd e areiro/ggywgw w wyygyww g ygywwg", 507},
		{"sokrrn e nnaigpl u aehgee/gwywgw y wyygyyw w wgwgwg", 508},
		{"eqebla m eupirel n mdlwal/ggywgw w ywygwww g wgyywg", 509},
		{"bnueet p moeioaz o sluwer/gwywgw g wywgwyw y wgywyg", 510},
		{"weoakm e hocberm i nprerl/gwwygw y wgwgyyw w ygyywg", 511},
		{"fiaclt u eatolhu t ltarae/gwywgw y wyygyyw y wgwywg", 512},
		{"tuaehl r emrdcnu i heoeby/gwgygw w wyygwww g wgywyg", 513},
		{"bexkrd c aemarih k geasat/gywygy w ywygyww g wgwywg", 514},
		{"smkupm w nknbeui e rgaiey/gyywgw w ywygwyw y wgwyyg", 515},
		{"cetrlu i haalcrt n epncih/gyyygw y wwwgwww y wgwywg", 516},
		{"feeonn w rewdinl g edolly/gwyygw w wgygwyw w ygwywg", 517},
		{"mulnsa i ilrboue o tfuily/gyyygw w wywgwyw y wgwywg", 518},
		{"puiweu o toeieeq t hrsaor/gyywgw w yywgwyw y wgwwgg", 519},
		{"airdne v irtneeo m icvnae/gyyygw w wyygyyw y wgywyg", 520},
		{"damnyg g ererame n ikuled/gwgwgw y yyygyww w ygwwyg", 521},
		{"tgheea i osreemt a htigrn/gwywgw w wwgggww y wgyyyg", 522},
		{"snilna r eueiblc a lmwigh/gwyygy w gwwggyw y wgwwwg", 523},
		{"ctcrms n drnoaal h rmaihy/gwwwgw y wyygyyw w wgwywg", 524},
		{"frbnhl o ooemeat o dmiiuy/gwywgw y gywggyw w wgwywg", 525},
		{"adluyk i iclnbio a eelery/gwgwgw g wywgwyw w wgyyyg", 526},
		{"sctrpm e tatobio u sftauy/ggywgw w ywygwww g wgyywg", 527},
		{"wikwnt n nnoonhr j ehceaa/gwgwgy g yywgwyw w wgwywg", 528},
		{"cecuke r oocdorn t aaamha/gwyygw y wyygwgw w ygwwwg", 529},
		{"belbtr u aatmorn e utewrl/gwywgy y yywgwyw g wgwwwg", 530},
		{"caceki e lltsaoe i nmhlnl/gwywgy y yywgyww w ygwwwg", 531},
		{"jreeto e sdwieed n eltruy/gwywgw w wyygyyw y wgwywg", 532},
		{"asalwp e wropjea c helagt/gwyygw y wywgwgy w ygwwwg", 533},
		{"aselto e cmradrl g awnvlt/gwgwgy y yywgwyy w ygwywg", 534},
		{"wihnfh s rzetino r rtfere/gyywgw w gwwgyyw y wgwyyg", 535},
		{"dttini a eeotmhc r ehdeoe/gwwwgw w wywgwyw y wgyyyg", 536},
	}
)

func TestSolve(testCases []TestCase) {
	total := 0
	count := 0
	for _, testCase := range testCases {
		waffle := board.Parse(testCase.serial)
		s := solver.New(waffle)
		if s.Solve() {
			path := pathfinder.New(s)
			path.Find()
			count++
			total += path.PathLen()
			fmt.Printf("Game: %3d Steps: %3d Average: %3.2f\n", testCase.index, path.PathLen(), float64(total)/float64(count))
		} else {
			fmt.Println("Unable to solve:", testCase.serial)
			s.Print()
			fmt.Println()
		}
	}
}

func main() {
	fmt.Printf("Welcome to waffle regression tests!\n")

	fmt.Printf("\nDeluxe Waffles\n")
	TestSolve(deluxeWaffles)

	fmt.Printf("\nDaily Waffles\n")
	TestSolve(dailyWaffles)
}
