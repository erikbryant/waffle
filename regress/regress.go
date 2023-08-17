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
		{"rdodtuve e e gseeciela b e vmarriuya r d asrnllap/yygwgwwy w w wgwgggwgy w w wgwgggwgw y w yywgwgyw", 64},
		{"ecietdpd e n noptopasi c i dpartiioc o r lermpgun/ywgwgyyy y y wgwgggwgw w y ygwgggwgw w w wywgwgww", 63},
		{"nutxqnns n t cdrnaiute k n tooiiisni m a eunguepe/yygwgwwy y w ygwgwgwgw g g ygygwgwgw w w yyygwgwy", 62},
		{"negeapar e o rdiippdda f p elmterdlu r a teadrlte/ywgwgwww y w ygwgggwgy w y ygwgggwgy y y yywgygwy", 61},
		{"memlidnv r c ssbspeute i o naciatera e r dvneoaci/wygwgwyw w w wgwgggygy y w wgwgggwgy w y ywwgwgww", 60},
		{"uotsntid n x delpvadno p r senieeete a e ghlgetis/ywgwgyww y w ygygwgwgw g g wgygwgwgw w y wywgygyy", 59},
		{"ussputbi r e mftnrsper b t leduiasee t c yelnhetl/wwgygyyw y w wgwgwgwgw g g wgwgwgwgy w w yyygwgyw", 58},
		{"deoitnai e i oaamcnrcd o n bdtsraida s e iglsdeel/yygwgwyy w w wgygwgwgw g g wgwgygygw y w ywwgygww", 57},
		{"eifdstal i p apertislt e e senithvte m t ueuedrra/yygygwyw w w wgwgggwgw w y wgwgggwgw w y wywgygww", 56},
		{"ntcmiiee t d iscnspitd n i uaeryluco e d aabledoo/yygwgyyw y w ygwgggwgy w w wgwgggwgw y w wwwgygwy", 55},
		{"nagntieg f n alotertlt b a uahilieyu e e idrgngae/wwgwgyyy w y ygwgggwgw w w wgwgggwgw w w wywgygyy", 54},
		{"isulelel r n aiostaitt i p lfnittara g a nwegucln/wwgygyyy w w ygwgggwgw w w ygwgggwgy w w wwwgwgyw", 53},
		{"ariiroed t e ndpntueeg t e sldtrisen s k trdraeae/yygwgwyy w w wgwgggygw w w ygwgggwgw w w ywygygww", 52},
		{"omaeiuoo g s ksmarmire r e eismmrdlt g n lmndneag/yygwgwww w y wgygwgwgw g g ygwgygwgw w w wwygwgwy", 51},
		{"olttatse c t tpherinrs a m taisiaseg r m uarnoenr/wwgygyww w y ygwgygwgy g g ygwgwgygw y w wywgwgwy", 50},
		{"eeirtvdt e h mbtpollrw n s ninoniegs i i rreemgaa/ywgwgwyy w w wgwgggwgy w w ygygggwgy w y wyygygwy", 49},
		{"srpcolav t s lpuintorn s t ieramaoey g m grredeel/ywgwgyyw y w ygwgggwgw w w wgwgggwgw w w yywgwgwy", 48},
		{"teceiehp i n iaoricutu f s efuolilht o o recsdroa/ywgwgwww y w wgygggwgy w w ygwgggwgw w w wywgygwy", 47},
		{"feorueue e n voeepdfen r r rirslactx t d crpevtrr/yygygwwy y w wgwgwgwgw g g ygwgwgwgw w w yywgwgww", 46},
		{"eeidpdum r c edpploras o s icroleeto k m raseedle/ywgwgwww y y ygwgggwgy w w wgwgggwgw w w yywgygww", 45},
		{"uenrbeua o l dnrediegi e h yonficire l v ggolfese/yygwgwww w w wgwgggygy y y ygwgggwgy w w wywgwgyy", 44},
		// {"reispoen l r pdsvelnpg n e dittrueer o d lesdeeio/yygwgwyw w w wgwgggwgy y w ygwgggwgw w w wyygwgww", 43}, <-- Ambiguous
		{"riterrro t o nthancuer w n eicspineu l e lvrteepn/ywgwgwyy w w ygygggwgw w y wgwgggygy w y wwygygww", 42},
		{"enfdurla l o lnltabuei n a amslaraat a o nacyitiy/ywgygwyy y w wgygggwgy w w wgwgggygw y w yyygwgww", 41},
		{"tmreueud n n hiaflaiem i i rtmskiagt u e afrlemfn/yygygwwy w y ygwgggwgy w w ygwgggygw y w wwwgygww", 40},
		{"ielirdsg e m hrpbuiadn i u lnotrisea t i ihartgnl/ywgygwyw y w wgwgggwgw w y wgwgggwgy w w wyygygwy", 39},
		{"oiceloue m e amsntgaro t g gerokienn l d aifthgea/wwgygwww w w ygwgwgygw g g ygwgwgwgw w w wyygwgyy", 38},
		{"crnsenrm k i dadreavyc e c rpuodudtr o p ltssreue/ywgygwyy w w wgygggwgw y w wgwgggwgw w w wywgygww", 37},
		{"ecnuooce o q aaouftuci t s dsvrjiuec s s lnlnseii/wygwgyww y w wgwgwgwgy g g wgygwgwgy w w ywwgygww", 36},
		{"daoarult r e uoatomenr b g eesablnrw m s uusterpa/wygwgwww w y ygygggwgy w w ygwgggygy w w wyygygwy", 35},
		{"nsbpalco e a tslorrnnb a i hidstamlg a w aiemelsg/yygwgwyw w w ygwgggwgw w y wgwgggwgy w w ywygwgwy", 34},
		{"pmanieid s u dsleciasl e d iaetriper d t siedllar/yygwgwyy w w ygwgggwgw y w wgwgggwgy w y wyygwgwy", 33},
		{"ecsouona o e dfrleiega u f gcureirle d p urdendsh/yygwgwww w y wgwgwgwgw g g wgwgwgwgw w w yyygwgyy", 32},
		{"sleaiust h v dnoccesrt l r ainoealeh n s tmretefl/ywgwgwyw w w wgwgwgwgw g g ygwgwgygw y w yywgygwy", 31},
		{"mllnpoel y g yibpliade o e oaeettlrn e a raidsrrc/wwgwgyyw w w ygwgggwgy w y ygwgggwgy w y ywwgwgww", 30},
		{"uuodpmpo i o rfetneese e s lsoortrrt d l aeekhduu/ywgygwyw w w ygwgggwgy w w wgwgggwgw w y wywgwgyy", 29},
		{"atbgrner i g moierdado e l yisducare l o aenrvlml/ywgygwyw w w wgygggwgw y w ygwgggwgy y w wwwgwgyy", 28},
		{"adotpwet a b rahscocde n e dtpibueen r h rtrgceuh/wwgwgwyy w w ygwgggwgy y w ygwgggwgw w w wwygwgwy", 27},
		{"smeeavps r x dsclaoftc i c taeppauee r t elleaosn/ywgygwyw w w wgwgwgwgw g g wgygygwgw w w ywwgygwy", 26},
		{"mirtieib p n latgllece a n apsteeetu i e xdnalsrn/ywgygyyw w w wgwgwgygy g g wgwgwgwgw w w yywgygww", 25},
		{"besnrxat c v agerolral v i lntaeeyto c r eivtiese/yygwgwww w y wgwgygwgw g g ygygygwgy w w yywgwgww", 24},
		{"ucpmonop r l tatategru y r eipflasea l k nrulreyn/wygwgyyy w w wgwgggwgw g g wgwgggwgw y w yywgwgyw", 23},
		{"tcvcsrou o r dbpneares o i itsueiett e g coiehkar/ywgwgyyw w w wgwgwgwgw w w wgygwgwgw y w wywgygwy", 22},
		{"henreubq n i tmerluree e q aduuotado o d yieearnc/wwgygyww w w wgygwgwgw w y wgwgygwgy w w yyygwgwy", 21},
		{"thfeccun h t isistasni n i husitisgl g o ndigceiu/ywgygwww y w ygygggwgy y w ygwgggwgw w w wywgwgwy", 20},
		{"eiitoael g n pcrptujed u n yclnteetb a a oimertlc/ywgygyyy y y wgygggwgw w w wgwgggwgy w w wywgwgwy", 19},
		{"iiaiadme l s umtnsoanb o e mreskilre s e dndaehoa/ywgwgyyw w y ygwgggwgw w w wgygggwgw w w ywygygww", 18},
		{"rtoiamtg i n naredomel r r nsnaeivap n f dellhgen/ywgwgwyw w y ygwgwgygw g g wgygwgwgy w w wywgygww", 17},
		{"eemoenom g a tceiubmrl s r tnrtoasec c a iptnneto/ywgygwww w w ygygwgygw g g wgygwgwgy w w wwygygyw", 16},
		{"maoxeyvu o p sertvattv d s rerooicea r t yrnrtnbn/wwgwgwyw y w wgygwgygw g g ygwgwgwgw w w yyygygww", 15},
		{"twsruldr l r ferumiied e l elngrosku e a abeevmrg/ywgygwyw y y wgwgwgwgy w w ygwgygwgy w w ywygwgyw", 14},
		{"asiipeon l m saecetrcs n e gdrsmiase m i noaritda/yygwgywy w w wgwgggwgw y w wgwgggwgw w w ywygwgyy", 13},
		{"nesrfaut e g ssaahitne p l ndtpmuael u i oolrbeer/wwgwgyyw y w wgwgwgygy g g ygwgygwgw w w wywgwgwy", 12},
		{"pnrbevvt n g tldmbiegc e o ienasiaec u i lvetanxo/wygygwww y w ygwgggwgw w y wgwgggwgw w w yyygwgyw", 11},
		{"autasedr n i onfrrafap d y eooieiynn i v senrpgoa/ywgwgwyy w w wgwgwgwgw g g ygygwgwgw w w wywgwgyy", 10},
		{"ceninrai t c uaatretsa c e stfrmines p i uaelired/yygygwyy w w wgwgggwgw w y ygwgggwgw w w wwwgygww", 9},
		{"hiawtcsn u o irydaaete r c esteltdeo d i srrsarpl/yygwgyyw w y wgwgygwgy g g wgwgwgwgw y w wwwgwgwy", 8},
		{"inmoacrm e e iaigolncc h e dstaeerdw s e lemeitna/yygygyyw w w wgygwgygy w w ygwgygwgw w w wywgwgwy", 7},
		{"cespnbsa n e saeoatdds a n irilrapem l e nstdgeee/yygwgyyy w w ygygwgygw w w wgwgwgwgy w w yywgwgwy", 6},
		{"muinoeum c m siaputtsn r c naneiorec t s mrdscesw/wygygwww w y wgwgygwgy g g ygwgwgwgw w w wyygygyw", 5},
		{"ruorglcu p u gorttaaen h t rsephouto t r uetsearh/yygwgwwy w w wgygwgwgw w w ygwgwgygw w y wyygygyy", 4},
		{"dsvsosgu r e rmdbrtors t c issaiines u b oyieueoe/wygwgwyw y w ygwgwgygy y w wgygwgwgy w w wywgwgww", 3},
		{"abnserou i l ialmeeist f b rarryaseu m y sellsdam/yygygwww w y wgygwgygw w w ygygwgwgw w w wyygygwy", 2},
		{"taenergt r e perpsersd m x iaretaler m n vscsrevd/ywgygyww y w wgygygwgy w w ygygwgwgy w w wwwgygwy", 1},
	}
	dailyWaffles = []TestCase{
		{"sheryp i aoccull u atidyl/ggwwgg w wwwgyyy y wgywwg", 573},
		{"sevepw e grdieio g nrtnmy/gwywgy y yywgwyw w wgwywg", 572},
		{"soofpf r aieftlc i ifsret/gwwwgy y wwwgwgy w ygygwg", 571},
		{"hvscna o ieugatu m rhulue/gwwwgw g wywgwyy y ygwgwg", 570},
		{"vangtd l iygdimi u elohle/ggywgw w wyygwyg y wgwwwg", 569},
		{"uimnlf a ehauten e gdlbnr/gyyygy w ywwgwwy w ygwgwg", 568},
		{"sneekc u tatucdr t pegliy/ggywgw w ywygwww g wgyywg", 567},
		{"dvegyo s eutamgr a uatvst/gwywgg y gywgwyw w wgwywg", 566},
		{"baeleg i creuego s secrlt/ggywgy w ywygwww g wgyywg", 565},
		{"dedipb m frmiuuo f utreoy/gwywgy y yywgwyw w wgwwwg", 564},
		{"srcawu e seoiiet w rtansk/gwwwgy y wwwgwgy w ygyywg", 563},
		{"chnehl s vnaivca c ufiafh/gwgwgw w wyygyyy w ygwywg", 562},
		{"slhekg v fplinla a etrent/gwwwgw w wwggwyy y ygwyyg", 561},
		{"unnepw e eesdrru s udsgrs/gwywgy y yywgwyw w wgywyg", 560},
		{"mtesti u ercoead n ksuclr/gwyggw w ywygwyw g ygwwwg", 559},
		{"wzeoyu a daoccnw r enaomy/gywygw w wywgwyw y wgwywg", 558},
		{"druasd l lyatele o otplsy/ggywgw w gyygwyy y wgwwwg", 557},
		{"sewenr u weonnex e pbhear/gwywgy w yywgwyw g wgwywg", 556},
		{"tebtyi t smlasbo m udapiy/gwgwgy y wyygyww w ygwwyg", 555},
		{"sllhtl g earplal u heetue/gyyygw w wwygyww w wgwywg", 554},
		{"ipertn e reoifrk s oruunn/gywwgg w wywgyyw w ygywwg", 553},
		{"aareet d srelabt i lmruih/gwywgw g wyygyyw y wgywyg", 552},
		{"seegni o phuteoo w uednre/gwywgw y wywggyy w ygwwyg", 551},
		{"cneoca l alabemy i oprmlr/gywygy y yyygyyw w wgwwwg", 550},
		{"syikfn e wlvolai f elrryr/gwgygw w wywgwwy w ggwyyg", 549},
		{"bziler r taratsa o alpeir/gywygw y wyygyyw w wgwywg", 548},
		{"bneyha c municre t ettsha/gwywgy y wwggwwy y ygywyg", 547},
		{"rroeti o abinmpr o gtcoae/gwywgy g ywygyww w wgwgwg", 546},
		{"daeinr t dcnepme c ihigee/gywggw w wywgywy w ygwyyg", 545},
		{"zsyeyr o kesatip e nalbbs/gyyygy g yywgwyw w wgwywg", 544},
		{"paueoi v gwnolhn l aeknnt/gywwgw w ywygyyg w wgwwgg", 543},
		{"rdnnne e neianii s angety/gwywgy w ywygywy w ygwgwg", 542},
		{"dratmu i drtiitt e roeepa/gwyygw w yyygwwg w wgwgyg", 541},
		{"mnlodl s erainhu g orbwuy/gywygw g wywgwyw y wgywyg", 540},
		{"agrvee l krfovfl p oteaay/ggwwgy g wywgywy w wgwywg", 539},
		{"crdegd t lieirir e irnoly/gwywgy y yyygyyw w wgwywg", 538},
		{"nouddy c ioastla m oriral/ggywgw w wywgwyy y ygwwgg", 537},
		{"dttini a eeotmhc r ehdeoe/gwwwgw w wywgwyw y wgyyyg", 536},
		{"wihnfh s rzetino r rtfere/gyywgw w gwwgyyw y wgwyyg", 535},
		{"aselto e cmradrl g awnvlt/gwgwgy y yywgwyy w ygwywg", 534},
		{"asalwp e wropjea c helagt/gwyygw y wywgwgy w ygwwwg", 533},
		{"jreeto e sdwieed n eltruy/gwywgw w wyygyyw y wgwywg", 532},
		{"caceki e lltsaoe i nmhlnl/gwywgy y yywgyww w ygwwwg", 531},
		{"belbtr u aatmorn e utewrl/gwywgy y yywgwyw g wgwwwg", 530},
		{"cecuke r oocdorn t aaamha/gwyygw y wyygwgw w ygwwwg", 529},
		{"wikwnt n nnoonhr j ehceaa/gwgwgy g yywgwyw w wgwywg", 528},
		{"sctrpm e tatobio u sftauy/ggywgw w ywygwww g wgyywg", 527},
		{"adluyk i iclnbio a eelery/gwgwgw g wywgwyw w wgyyyg", 526},
		{"frbnhl o ooemeat o dmiiuy/gwywgw y gywggyw w wgwywg", 525},
		{"ctcrms n drnoaal h rmaihy/gwwwgw y wyygyyw w wgwywg", 524},
		{"snilna r eueiblc a lmwigh/gwyygy w gwwggyw y wgwwwg", 523},
		{"tgheea i osreemt a htigrn/gwywgw w wwgggww y wgyyyg", 522},
		{"damnyg g ererame n ikuled/gwgwgw y yyygyww w ygwwyg", 521},
		{"airdne v irtneeo m icvnae/gyyygw w wyygyyw y wgywyg", 520},
		{"puiweu o toeieeq t hrsaor/gyywgw w yywgwyw y wgwwgg", 519},
		{"mulnsa i ilrboue o tfuily/gyyygw w wywgwyw y wgwywg", 518},
		{"feeonn w rewdinl g edolly/gwyygw w wgygwyw w ygwywg", 517},
		{"cetrlu i haalcrt n epncih/gyyygw y wwwgwww y wgwywg", 516},
		{"smkupm w nknbeui e rgaiey/gyywgw w ywygwyw y wgwyyg", 515},
		{"bexkrd c aemarih k geasat/gywygy w ywygyww g wgwywg", 514},
		{"tuaehl r emrdcnu i heoeby/gwgygw w wyygwww g wgywyg", 513},
		{"fiaclt u eatolhu t ltarae/gwywgw y wyygyyw y wgwywg", 512},
		{"weoakm e hocberm i nprerl/gwwygw y wgwgyyw w ygyywg", 511},
		{"bnueet p moeioaz o sluwer/gwywgw g wywgwyw y wgywyg", 510},
		{"eqebla m eupirel n mdlwal/ggywgw w ywygwww g wgyywg", 509},
		{"sokrrn e nnaigpl u aehgee/gwywgw y wyygyyw w wgwgwg", 508},
		{"reuori e lodewcd e areiro/ggywgw w wyygyww g ygywwg", 507},
		{"cmhett e aagaalp r rtaret/gwywgw y wywgwyw w wgyyyg", 506},
		{"paldes g eenugtu t noroer/gwwwgy w yyygyyw w wgwygg", 505},
		{"baarkk t lcaerxn p redola/gwgwgy y ywygyww g wgwwwg", 504},
		{"wleerr o sapogts e arapry/gwwggw y yywgyww w ygwgwg", 503},
		{"cmnsar i haochwi y etthot/gwywgw y wygggyw y wgwywg", 502},
		{"seostl l seohatn e upeecr/gyywgy w wywgyww y ygywwg", 501},
		{"daieoe e couvowr g glnene/gwywgy g yywgwyw w wgwywg", 500},
		{"sielkr t uuadoin n epcwml/gwywgw w wyygyyw y wgwwwg", 420},
		{"hamele e uoamivt o olyetr/gwwggy w wywgwwy y wgwyyg", 419},
		{"crnsgi e emsarui s nplgns/gwywgy y yywgwyw w wgwwwg", 418},
		{"gosadr a tenisra v gelnue/ggywgg w wwygwwy w wgyywg", 417},
		{"bbrtma o froimuw s pmoeoo/gwwwgw w wyygyyw y wgwywg", 416},
		{"vahmro c isiapan l arlery/gyywgy y wgwgwyw w wgwywg", 415},
		{"snaeem r etletio v henrvr/gwwwgy y yywgwyw y wgwwwg", 414},
		{"peroeo o akndnnl o baenry/gwyygw y wywgwyw w ygwgwg", 413},
		{"pemtob n acvorgn e vhiyaa/gwwwgw w wyygyyw y wgwywg", 412},
		{"woeadt l reabulh r efmwll/ggywgw w wyygwyy y ygwwwg", 411},
		{"hplntt i ganisea r geeear/gwwwgy w yyygyyw w wgywyg", 410},
		{"ssnotl r acataer n atcoxt/gwywgw y gyyggww w wgwywg", 409},
		{"idrare d gkliaot u elnnee/gwywgy y yywgwyw w wgwwwg", 408},
		{"pealnr n ltosaro l aeonyy/gwwygw w wyygyyw y ygwwwg", 407},
		{"ttolbo i uiraego n spharo/gwywgw w wyygyyw w wgwywg", 406},
		{"txrite i iiramct a realpl/gwyggw w yyygwwy w ygwwyg", 405},
		{"thfrmr u ulcocat o etnteh/gwwwgg w gywgwyw w wgyyyg", 404},
		{"tpalti g msourea i rriynd/gwwwgy y wwwgwgy w yggywg", 403},
		{"mmubyc t eecnoui i dreuda/gwywgw w wyygyyw g wgwywg", 402},
		{"miidcr g aognone m alirie/ggywgw w wyygwyy y wgwgwg", 401},
		{"bsmcye s eaiarsl l adeeks/gwywgy y yywgwyw y wgwgwg", 400},
		{"agdrml r ueianeu l oeibrr/gwywgw y wwygywg w ggwywg", 300},
		{"lieirs w riipese n rcouye/gwywgw y wygggyw w wgwywg", 200},
		{"mmkoye u iaomerr o pahcln/gwywgw y wyggwyw w wgwwwg", 100},
		{"satemu r aeagbby c abugre/gyyygy y gwwggyw w wgwyyg", 40},
		{"cmroac s opmiorc o heepdt/gwwggw w ygygyyw w wgywwg", 39},
		// {"stcalu w tlldear i aetelh/gwwwgw w yywgwyw w wgyywg", 38}, <-- Ambiguous
		{"qtguly u gduioet n eneily/gwwygw w ywwgwyw w wgwywg", 37},
		{"frlels o eidinoi g ltmuar/ggywgy y ywwgyww w wgyywg", 36},
		{"crhamt t bunokll r ietray/gyyygy y wwwgwgw w ygyywg", 35},
		{"fslahm e gllikan l uevgay/gyyygy y wywgwyw g ygwywg", 34},
		{"prhnya e dweanos e kkurhl/gwywgy y yywgwyw w wgwwwg", 33},
		{"datrmu y tsaacro e revovr/gggwgw w wwwgwyy w wgwwwg", 32},
		{"diloyn k tladreu e iynaws/gggwgy w ywwgyyw g wgyywg", 31},
		{"luuhet e apgraut e olsmny/gwyggw w wgygwyw w wgywwg", 30},
		{"aeaayo a lelrvls c rretet/gwywgw g wywgwww w wggwwg", 29},
		{"mldaht o rnotial e utnuty/gwwwgw w yywgwyw w wgwywg", 28},
		{"fainlu e oseorve e esrdty/gyyygy y wwygyyy w ygwywg", 27},
		{"elcutr i ahxourl m atcghe/gwwwgw w ywwgwyw w wgwywg", 26},
		{"setnei o hoetteb r ndrlay/gwwggw w wgygyww y wgywwg", 25},
		{"frjodn r teihhic i udttey/gyyygy y wgygwww w ygyywg", 24},
		{"idsvrc n nlmoria r aegfee/gwwwgw w ywwggyy g wgwywg", 23},
		{"pxamen t urrbrlo w odiabl/gwwwgw w wwwgwww w ggwgwg", 22},
		{"vuemtr r eeetgio e reovrt/gwyygy y wyggwww g ygwwyg", 21},
		{"kaayka s elultat n idctey/ggyygy y wwwgyyw w ygwywg", 20},
		{"eothss y olrluoo x ctgsit/gyyygy y wwwgwgw w ygwywg", 19},
		{"clrial i aurpewe h ayzrcd/gwgygw w ywwgwyw w wgwwwg", 18},
		{"boedts n ahoagya e rwalel/ggywgw w ywwgwwg w wgwwgg", 17},
		{"sdrder i oipdepf l igiray/gwywgy w wywgwww w ygggwg", 16},
		{"yiyltu e vharakl c isocpt/gwwwgw y wwwgwwy y ggwwgg", 15},
		{"siatno y odhdver h etrmoe/gyyygy y wwwgwyw w wgwywg", 14},
		{"dggoam e gioinwn h tllelt/gwgygw w gwwgwww w wgwwwg", 13},
		// {"arcxet b ezmonje n aezodt/gwwwgw w yywggyw g wgwywg", 12}, <-- Ambiguous
		{"cldara h dxopeul i vreeey/gwwwgw w ywwgwyw w wgwywg", 11},
		{"bnebkn e avlidll n vemalt/gyywgw w wyggwyy w wgwywg", 10},
		{"tjmilo r agoailn l kneaia/gwwggg w ywggwyw w wgwyyg", 9},
		{"slropa m orodcht r ontreh/gwywgw w gyygwww w wgygwg", 8},
		{"onwigr o nrrtcra v rtxpee/gyyygy y wywgwyw w wgwywg", 7},
		{"socopr t ceatsen t amfeoy/gyyggy y wywgwww w ygwywg", 6},
		{"crmvpe r glaivye n belouy/ggwwgw w yywgwyw w wgwywg", 5},
		{"ndeeye e ltraeck a idnsks/gygygy y wwwgwyy w ygwywg", 4},
		{"speeda t ptocirn e mpeiey/gggwgw w yywgwyw g wgyywg", 3},
		{"scgoln n dindeer i uffare/gwwwgg w yyggwyw y wgwyyg", 2},
		{"fboueg i ulsoomg e loemna/gwwggw w wgygyyw y wgyywg", 1},
		{"bcstda a ieuotln e dnrloy/gwywgg y gwygyyw y wgwwwg", 0}, // Waffle's webpage review image
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

	fmt.Printf("\nDaily Waffles\n")
	TestSolve(dailyWaffles)

	fmt.Printf("\nDeluxe Waffles\n")
	TestSolve(deluxeWaffles)
}
