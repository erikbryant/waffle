package main

import (
	"flag"
	"fmt"
	"strings"

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
		{"dleeiaep u a aerpenoex e e xtrntryma g l nssegdda/ywgwgwyy w w ygwgggwgw w y wgwgggwgy w w wwwgwgyy", 112},
		{"ohfneera d s lchuaisgp n t ktmdtosei p e aeireers/wygwgyyw w w wgwgygygw g g wgwgwgwgy w w wywgwgwy", 111},
		{"lsfaetad l s tbrrclpre s i etdalpeek m e onuemego/ywgwgywy w y ygwgwgwgy g g wgwgygwgw w w ywwgwgww", 110},
		{"egcaaedo e i srpmxneee p e dhracskti e a ncosdabm/yygygwyw w y wgwgwgwgy g g wgwgwgwgw w w wywgygyw", 109},
		{"rrttadel i u mpiatoanc o o sottraleg p v rdcrcoii/ywgygyyy w w wgwgggwgw y w ygwgggwgw w y wywgwgyy", 108},
		{"cdirpmaf l e etsmbouae e f ohrrvelts p d dialotoe/yygwgwwy w w ygwgggwgy w w wgwgggwgw w w yywgwgwy", 107},
		{"ngdnefbc d n iansratns s m eeeitidet o s toxrmeaa/ywgwgwww w y wgwgwgygy g g wgwgwgygy y w yywgwgwy", 106},
		{"maceolde s u rciaotnce a s lahmleesx n h adedilmi/ywgygyyy w w wgygggwgw w w wgwgggwgy w w yyygwgyw", 105},
		{"yracieou l o itnaitrre s o bcldenned s a nstneeen/wwgwgyyw w w wgwgggygy w w wgwgggwgy w w yywgygww", 104},
		{"maoeercm w a lalauehrr t r ibtotnsey t r oesyrlsl/ywgwgwyw w y ygwgygwgy g g wgwgwgwgw w w wywgygwy", 103},
		{"ectehuea d e ulandtrco k e atrakfrli s h iaadelut/yygwgwyw w w wgygwgwgy g g ygygwgwgw w w wwwgygwy", 102},
		{"ulneicyp r h eaitacede e i nhlavisrg c a tueelrnl/ywgygwyy w y ygygggwgw w w wgwgggwgy y w wwwgygwy", 101},
		{"bearhbse n r eopserees d c vsdecicsa n e oubsdees/yygwgwyy w w ygwgggwgw w w wgwgggwgy w y wywgygyw", 100},
		{"raanigeo v r earrostln c l egoaniler l t hadyylrt/ywgygyww w w wgwgggwgw w y wgwgggwgy y y wywgygyy", 99},
		{"enotadss r g gfttpensi d r eimoaanet c n lluroevi/ywgwgyww w y wgwgwgygy g g ygwgwgwgy w w wywgwgww", 98},
		{"cordvddn l d evwnegert a i rsuomarha u a aiinctae/wwgwgwyw w w ygwgggwgw w w wgwgggwgy y y wywgwgyy", 97},
		{"einnoiup r l isoaniile i y edgiftmdr e l snegtdre/yygwgwyw y w ygwgggwgw w w wgwgggwgy y w yywgwgyw", 96},
		{"gofbaiei t f dbniocvei t n yshawerde u a reaertrl/wygygwwy y w ygwgggwgw w y wgwgggwgw w w yyygygww", 95},
		{"ihgntuic d e suhatirma z c utiihklrr e a rnegeten/yygygwww w y wgwgwgygw g g ygwgygwgy w w wwwgwgwy", 94},
		{"lhllgini a l csnatlotl d l yaainaeeo l g ttrgttre/ywgygywy w w wgwgwgwgy g g ygwgwgwgw w w yywgwgww", 93},
		{"atoregrd e i earemosec l p ssnabieds p r nrneemrt/wwgygwyy w w ygwgggwgw w w ygwgggygw y w yywgygyw", 92},
		{"lveiadae p r rchasscce i e edtsseetl l n uotdxeie/yygwgwwy y y wgwgggwgy w w ygwgggwgw y w wywgwgwy", 91},
		{"motnosli m l kcwuntuyr r e rasamadto r n npaiotcs/ywgwgwwy y y ygwgggwgw w w ygwgggygw w w ywwgygwy", 90},
		{"ctuniuaa e a elltiilee e t ptdraignh o o dlsdtlon/wwgygwyw y y ygwgwgwgw g g wgwgwgwgy w w yywgwgww", 89},
		{"nnoaepot e u iontgrrwa h r raschioep i r coeligvn/ywgwgyyw w w wgwgggwgw w w ygwgggwgy y w wyygwgwy", 88},
		{"anvlleoh r c dmdneaiae i t edalehane r y oolatrip/yygwgyyw w w ygygwgwgw g g ygwgwgwgw w w wywgwgww", 87},
		{"insufnoa i t naiadrsme n b ecunoeeta f e lcgdnrie/ywgygwww w w wgwgwgwgy g g ygwgygwgw w w ywwgygwy", 86},
		{"gebuorid e c iceuttlrt s p uaectisna y l lroruhno/wwgygwyy y y wgygggygw w w wgwgggwgw w w wywgwgww", 85},
		{"nhruitvv l m orlcccyec r l seiuniset v e lseteeen/wwgygwyw w w wgygwgygy g g wgwgwgygw w w yywgwgwy", 84},
		{"euadmlnn m u lamtsnimi e d utenrfdlt e n pioaelye/ywgygyww w y wgwgwgwgw g g wgygwgwgy w w yywgygwy", 83},
		{"naiekoee h u sitspedtc h f rtiidertn c e nssgetnr/wwgygwyy w w wgwgggwgw w w ygwgggygw w y wywgygwy", 82},
		{"crsaeeby t t epnttesnn a r rexlarneo g e yurrnwsa/wygygyww y w wgwgggwgw w w ygwgggygw w w ywwgygyy", 81},
		{"uaittnps h n eirfnrrol l r nrnaeitgr e h tedeteur/ywgygwwy w y wgwgygwgw g g wgygygwgy w w ywwgygwy", 80},
		{"hpnenela p c asaittten a l giiherute u e cderelnb/yygygyyw w w wgwgggwgy w w wgygggwgw w w yyygwgwy", 79},
		{"iavslrid c d upertnpra o e slnngeetd a e giaaetri/ywgwgwyy w w wgygggwgy w y wgwgggwgy y w wywgwgyw", 78},
		{"eaowttne s m eentrekec d d uptayfdll c u xwltehli/ywgwgyyw y w ygwgggwgw w w wgwgggwgw y w wwwgygwy", 77},
		{"nrrteiid n r hrcgciigt a t wuiteitto o e siearncg/wygygwyw w w ygwgwgygy g g wgwgwgwgw w w wyygygww", 76},
		{"aelennki c c nrlostsro s n eioitaaev m a torgtete/yygwgwww w w ygwgggwgy w w ygwgggwgy w w ywwgygwy", 75},
		{"tnsiatod d i ecentaxnb r i eehitiagn t u ihdlngto/yygygyyy y w wgwgggwgw w w wgwgggwgw w w yyygwgyw", 74},
		{"aafuzvea a u tnnrnaaay e l nmemdomhi e t dienyslt/ywgygwyw y y ygygwgwgw g g wgwgwgygw w w ywwgygwy", 73},
		{"dmauevab h l etxreigey e s rteactndi t s reeebdol/wwgwgywy w w ygwgygwgw g g wgygwgwgw w w wywgygwy", 72},
		{"gnsdesdt n l csraitenn t r auaupulli i p sneastss/wygygyyw w w wgygwgwgw g g wgygwgwgy w w wyygygww", 71},
		{"notogeen f c ssiroanmo n c cecoeibnu l c aoefsrar/yygygwyw y w wgwgwgwgy g g wgwgwgwgw w w wyygwgww", 70},
		{"opealsae o m maeptirts e s cplrluret e e tagrlenr/wwgygyyy y w wgwgwgwgy g g wgwgwgwgw w y yyygygwy", 69},
		{"nsasante c m sueherddr a r etrrranet s l eseyiepf/ywgwgwyy w w ygygggwgy y y ygwgggwgw y w wywgwgww", 68},
		{"cacgihsa i h ncealkrei p w eantcersd c e rlrrhlcr/yygygywy w w wgwgygygy g g ygwgygygw w w ywwgwgyw", 67},
		{"artdamap n e iawprohee s e ntcothldv e k oerttroe/ywgwgwyw w y wgwgggwgy y w ygwgggwgw w w wyygwgwy", 66},
		{"bnoretpn y a artlyihgl t s eimquentc g u nrlentii/wwgygyyy w w wgwgggwgw w w wgwgggygw w w wyygwgwy", 65},
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
		{"maoxeyvu o p sertvattv d s rerooicea r t yrnrtnbn/wwgwgwyw y w wgygwgwgw g g ygwgwgwgw w w yyygygww", 15},
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
		{"naaaro v eeaucta b clncsh/gwywgw y wwygywg w ggwywg", 910},
		{"sanoaf e ahdfoul u ltaony/ggywgw y wywgyww y ygyywg", 909},
		{"aaitdg l wlebalr m oniaul/gwywgg w gyygyyw w wgwywg", 908},
		{"dbrani r dhwogua t vliuae/gwyygw w ywwgwyy y wgwyyg", 907},
		{"inrreo u trnieol r ptoaer/gwywgw w wywgwyw w wgyyyg", 906},
		{"tvopls r milapri o yclirt/gwgwgw y ywygywg w wgwwyg", 905},
		{"cnsuno e olimaut o lhtkdy/gwwwgw g wyygyyw g wgwywg", 904},
		{"adtpni e lognais a oecrmh/ggwwgg w wwwgywy y ygywwg", 903},
		{"laetdc a cetllio k ihoiph/gwwwgy y yyygyyw g wgwwwg", 902},
		{"untcde l ewpiher s mrreoy/ggywgy w ywygwww g wgyywg", 901},
		{"gahcno e miagoca n trrrih/gwwwgw w wywgwyy y ygwywg", 900},
		{"tlmngt a wehogum i nadnje/gwyggy y wgwgwyw w wgwywg", 899},
		{"hcceer a ouewrta l nherey/gywygw w wyygyyw w wgwgwg", 898},
		{"clickb u umhiunr a hcpaet/gwwggy y yywgwww w wgwyyg", 897},
		{"sliatm g fteaeir p hfarld/gywygw g wywgwyw g wgwywg", 896},
		{"cawodn e rhnteoy s rsrefr/gwyygw y gyggyww w wgwwwg", 895},
		{"lerinu s eyvsito n armehe/gywygw w wwygyww y wgywyg", 894},
		{"poragr f eoneala t dlmegy/gyywgw w wwwggyg w ygwywg", 893},
		{"puslek l eoeditv m atnroe/gywygw w wywgwyy w ygwgwg", 892},
		{"mvaore n ispiiuk m wcelhy/gwyggw y wgygwww w wgywyg", 891},
		{"cvrgbe p muiprue r aeuide/gwywgw g wyygyyw w wgwywg", 890},
		{"geehnu e urrgaer n etrvae/gygwgw y wywgyyw w wgywwg", 889},
		{"ideirr u ennbnim a cteeur/gwwwgy w yywgwyw y wgywyg", 888},
		{"riised e epaelee f itwxll/ggywgw w yywgwyg y wgwwwg", 887},
		{"stamea a glmoeni w yprlne/gywygy w yywgwyw w wgwwwg", 886},
		{"kdepal f ueaodbh l kiaont/gwwwgw g wyyggyy y wgwwyg", 885},
		{"llrmha n rmobnau d aaaegy/gwywgy y yywgwyw w wgwwwg", 884},
		{"hnroyt a esarwuc m dymaay/gwywgy y wgygwyw w wgwywg", 883},
		{"aluieh o oldapnr g lailrw/gwywgy w yyygyyw y wgwgwg", 882},
		{"pgroeo a llisnba o ttnasy/gwywgy y wgwgwyw g wgyywg", 881},
		{"muooet a bgyuerh n aleagr/gwywgw w wywgwyw g wgyyyg", 880},
		{"sdaakl e hrcileb a npkllt/gwgwgw y wywgygw w ygwywg", 879},
		{"otearw d craafie f tttrry/gwywgw y wywgwyg y ggwywg", 878},
		{"bacnhi e cevroei l awvrue/gwywgw g wyygwyw w ygwywg", 877},
		{"meroya l ureiter d aroodl/gwywgw w wyygyyw y wgwwwg", 876},
		{"stciki o hatglit a idltny/gwyygw y wyygyyw w wgwgwg", 875},
		{"cnbrke i vulilnr l adolnl/gyyygw y wywgwyw y wgwywg", 874},
		{"vedngl i ysoliia v utrohm/gwyggw y wgygwww w wgywyg", 873},
		{"ceuuls m rsiagit s eewnae/gwwwgw g wywgwyy g ygwywg", 872},
		{"vudiat c lnsarua e uegivt/gwwygw g wwwggyy y wgwwyg", 871},
		{"saetda w tttanok t spoiry/gyyygy g yywgwyw w wgwywg", 870},
		{"sttcbe u oirvoie r ctrurh/gwwygw w wyygwyw y wgywwg", 869},
		{"mriacp a haoiscz i ewrmhe/gwywgw g wywgwyw w wgyyyg", 868},
		{"crxehs l ourbnam o etioen/ggywgw w ywygwww g ygyywg", 867},
		{"epnler g nutiebc g wroley/gyyygw y wwygyww w wgwwwg", 866},
		{"serrda l speirae i rghagd/gwwggw y wyygyyw w ygwywg", 865},
		{"sfntft e dptieew i utnecd/gyyygw w wyygyyw w wgwgwg", 864},
		{"vuerma t anoduce l neaaih/gwywgg w wywgwyw w yggywg", 863},
		{"rzteeo c lroaxgs e uegrut/gwywgg w gywgwyw y wgwywg", 862},
		{"ypoeht i lolsieu t adeeex/gwywgw w ywygwyw y wggywg", 861},
		{"rtceht a ietoflo u uterar/gyyygw w wyygyyw g wgwywg", 860},
		{"clathc p mhnuhpo c daroay/gwyygy y wyygwwg w ggwwwg", 859},
		{"snfufp o hautdee d nrecey/gwywgw y wywgwyg y ggwwwg", 858},
		{"teoaos r hewaedr y wlstar/gwywgw y ywggwwy y wgwwyg", 857},
		{"wrrehs e thikcer o dfwaay/gwywgw y wywgwyg w ggwywg", 856},
		{"bphoet r maaastp i msatiy/gwywgw y gywgyyw y wgyywg", 855},
		{"raoyon a teccaoo e drioen/gwywgw w wyygyyw y wgwywg", 854},
		{"peicnm e reabiry g iducee/ggwygw y wywgyww w ygwywg", 853},
		{"jsoutc e oarcvmr l bonhut/gyyygw y wywgwyw g wgwwwg", 852},
		{"cniels m uhaomde a opcgre/gwwggw y gyygwyw w wgywwg", 851},
		{"dsrloe n pedkaeg e ieddar/gwwwgw y wyygyyw g wgwgwg", 850},
		{"basptd e lnntryl j recuit/gwwwgy y wwwgwgy w ygyywg", 849},
		{"alfatu e heodugi r nloair/gwywgg y gywgwyw w wgwywg", 848},
		{"brhudm b vylbsee a etnere/ggwwgw g ywygwyw y wgwywg", 847},
		{"imweep e oiraepa x rlnnny/gwywgy y yywgwyw w wgwwwg", 846},
		{"ravgli a ieaoenr e etlpet/gywygg y wyygwww w ygwwgg", 845},
		{"dkrsdi o rntnuic i ayeiey/gwywgy g yywgwyw g wgwywg", 844},
		{"thmgbo i ortaeei l nreaan/ggywgw w ywygwwy g wgyywg", 843},
		{"mimofe o halutgc d aaoybr/gywygw g wywgwyw w wgywyg", 842},
		{"cuskms n mraismr e veahvy/gwywgw g wyyggyw y wgwwyg", 841},
		{"ssaofr i roitcar n fpette/gwwwgy y yywgwyw y wgwwwg", 840},
		{"tirnnt g eaihthi e chsiee/gyywgy w wwygyyw g ygwywg", 839},
		{"cyrsde l vkaortr i asrmre/gwywgw w wwygywy y ygwwwg", 838},
		{"fluine t rnevron e ahoroy/gywwgw w wyygyww g ygwgwg", 837},
		{"dnauki n nrxcrea t reucma/gywygw g wywgwyy y ygwwwg", 836},
		{"ainhnh g sbiiniu l cegsgt/gyywgw y wywggyy w wgywwg", 835},
		{"daaeaz m zlbarpm e otacwk/gwgwgw g wywgwyw y wgywyg", 834},
		{"vworlo a yeracep m areail/gyywgw w wwggyyw w ygwywg", 833},
		{"qlreal p aueeudt o atriey/gwywgw g wywgwyw y wgywyg", 832},
		{"fnoana w xroivoa r lkmned/gwgwgy w wyygwyw w wgwyyg", 831},
		{"fxupse l sominer p ltdpot/gwywgw w wgwgwgy y ygwywg", 830},
		{"flsihi o ecanero z oontge/gwwwgy y wwwgggy w ygyywg", 829},
		{"mtstee a iatutno d togerr/gwwwgw g wwygywy y ygwywg", 828},
		{"dpgdeo c eeosldr t gaoeit/gwyygw y yyggyww w ygywwg", 827},
		{"drkema u iieuroo e gtltcn/gwywgw w wyygyyw w wgwywg", 826},
		{"taepym k iaarura s oraetn/ggywgw w wyggwyw y wgywwg", 825},
		{"orlyne d alcipte o acailn/gwywgw y wywgwyw w wgywyg", 824},
		{"traopm e wnedeiu i rlidyc/ggwwgw y wywggyw w wgyyyg", 823},
		{"csrccm n lehbhoo r eesoit/gwwwgw y wywgwyy w ygwywg", 822},
		{"tagutr p ihnuals l itmsay/ggwygw w gywgwww y wgwyyg", 821},
		{"tteake t eoelunl i nmnrul/gwwwgy y yyygyyw g wgwwwg", 820},
		{"soxltl n rienpuu e hepcwe/gwwwgg g ywygwyw y ygwwwg", 819},
		{"fdjedv a airocho o itbrey/gwywgw w wyygyyw y wgwwwg", 818},
		{"mgguli e larioda r ocpdve/gwwwgy y wywgwgw y ygyywg", 817},
		{"bglddu t oneuaol r ttnxny/gwywgg y gywgwyw w wgwywg", 816},
		{"edgael o intetla f enavxl/gyyygg w wwygwyw w wggwwg", 815},
		{"stornd r fdefnoi e etpany/gwgwgw y wwygywy g ygwywg", 814},
		{"ecoile n txlleoo g osshoe/gyywgw y wwygygg w ygwwwg", 813},
		{"surkep h rrotnik o vtdove/gwwwgw w wyygyyy y ygwgwg", 812},
		{"blkoea e tutstoe c nhrmny/gwyygw y wyyggww w ygwwwg", 811},
		{"vrioor e oamedec e klaaan/gwywgw w wywgwyy w ygwywg", 810},
		{"dctuhr e arrgoan l ulecnh/gwywgg y wwggyyw w ygwywg", 809},
		{"tmbhbn i siysoue t uaeelr/gywygw w wyygyyw y wgyyyg", 808},
		{"aoromt a pralomu i enbldh/gwywgw g yyygyyw w wgwwwg", 807},
		{"reiltn d maoidne h ooolbr/gwwwgy y ywygyww w wgwywg", 806},
		{"kfpnla o nbeieul s lbtesy/gwyygy y wgwgwyw w wgwgwg", 805},
		{"bispna r ygmapwu a slcanh/gwwwgg w gwygyww y wgywyg", 804},
		{"gieare r gvnrefp s otmhae/ggywgw w ywygwww g wgyywg", 803},
		{"mtniri d ielarew s oaeayc/gwwwgy w yyygyyw g wgwwwg", 802},
		{"lmchhe r eeytmer u oredne/gwywgg y wywgyyw w ygwwwg", 801},
		{"deoipe i amrbtlr f rtoarn/gwgwgw y wywgwyy w ygwywg", 800},
		{"lcixtu i iaantmc h aeiiat/gwywgw w yywgyww w ygwwyg", 799},
		{"boerdf o uletool s nhinly/gyyygw y wywgwyw w wgwgwg", 798},
		{"tecakh h naaeerh r ywdank/gwwygg w wwwgyyw y ygwgwg", 797},
		{"cjefea l smlohte t itrnle/gwywgw g wywgwyy w ygwywg", 796},
		{"reeboy i ecestuo t rdacrr/ggywgw w wywgwyy g wgwwyg", 795},
		{"rralth l gnteigi c iteave/gwgwgy w yywgwyy w ygwwwg", 794},
		{"hevphd o efdoity a adrety/ggywgw w gywgwyy y wgwwwg", 793},
		{"faanlo e vluvgii o rsecae/gyyygw y wwwgwwy w ygwwwg", 792},
		{"vanttn c akeioei u ilraln/ggywgw w gywgwyy y wgwwwg", 791},
		{"rcueya o cllbeac i upkala/gwwwgy w ywwgwwy w ygyyyg", 790},
		{"wirdrl s tslooki d ohriey/ggyygw w wywgwyy y wgwwwg", 789},
		{"shuoen t uampral i wmeitl/gwywgw y wgwgwgw y wgywyg", 788},
		{"bcouht t reattlr r pnaaiy/gwgwgw y ywygwyy w ggwwwg", 787},
		{"woeedu e rthguca l opichh/gwgwgw y wyygyyw w wgwywg", 786},
		{"hihaye r rdlaseo p wdwopl/ggywgw w wywgwyy y wggwwg", 785},
		{"esenli e ttpsoee a kgnsxe/gwywgy w yywgwyw w wgwywg", 784},
		{"botcno e icrmenr t shuaah/gyywgw g wwwggyy w wgwywg", 783},
		{"vimade e mllaapu p ersiry/gywygy g yywgwyw y wgwwwg", 782},
		{"baesye e owdnsni d enialy/ggywgw g wywgyyy y wgwwwg", 781},
		{"bnvehe e slcvsar n atleie/gywygw g wywgwyw y wgwywg", 780},
		{"celrhe s oasgeca v nnuuue/gwywgy w wywgyyw w yggywg", 779},
		{"cecekt t ettainr v iporcy/gwywgw y wywgwyy y ygwwwg", 778},
		{"ceegkr m rlorare e npeill/gygwgw w wywgyww w yggywg", 777},
		{"goulme a mtmiukr o uysdry/gyyygw w wywgwyw y wgwywg", 776},
		{"crenke s veuumnp a ithoed/gwwggy g wywgyyw w wgwwyg", 775},
		{"sekmlo c atpeiml i gpstnh/gwywgy w yywgwyw y wgwgwg", 774},
		{"baanee v uenisno r psavgt/ggywgw w wyygwyy y wgwwwg", 773},
		{"sitgla w aopresu d anctle/gwywgy w yywgwyw y wgwwwg", 772},
		{"blefez e raeagir t mtewxt/ggywgw w ywygwww g wgyywg", 771},
		{"strooe n exuiyse l hfvear/gwgwgy y yywgwyw w wgwywg", 770},
		{"clveno e ceetuwu t dhgroe/ggwwgg y wywgywy w wgywwg", 769},
		{"ptlelu r eaebuii t aytlad/gyyygw w wyygyyw y wgwywg", 768},
		{"swnipm n tativfd a ofcaoy/ggywgw w ywygwww g ygyywg", 767},
		{"sednkt c cuduell l aeeclt/gwwwgw g wyygyyw w wgyyyg", 766},
		{"qulaew o lteboeu e ombnur/ggywgw y wywggyy y wgwwwg", 765},
		{"bemeos e theuwpa e strnvd/gwywgy y yywgwyw w wgwwwg", 764},
		{"wofihi t aeggetn c irhceh/gwwygw y wywgwgw w ygwyyg", 763},
		{"mkoila r oidpdoc a ortuuy/gwywgw y wywgwyw y wgwywg", 762},
		{"bapgeh a utsarel m slsane/ggywgw w ywygwww g ygyywg", 761},
		{"senaad l ahcigxa u lerlby/gwywgy w yywgwyy y ygwywg", 760},
		{"bslobr o rtzhteu r teeeah/gwywgw w yywgwyw y ggwyyg", 759},
		{"darlsr m aeobulc e erpesl/gwywgw g wywgwyy g ygwywg", 758},
		{"skprlt e eneseah u ediulr/ggywgw y wywggyy y wgwwwg", 757},
		{"fubrti r naibrca o umgwae/gywygw g wyygyyw w wgywyg", 756},
		{"bglrne l dinagld n aeaeur/gywwgw g ywygyww g ygwywg", 755},
		{"gneymo o rledreo h ttsrrh/gwywgy y ywygyww w wgwwwg", 754},
		{"sepuee o ranpler l pdeaey/gyywgw w ywwgwyw y wggywg", 753},
		{"cyiiov g lrlivnl l emahet/gwwwgw y wwgggwy y ygwywg", 752},
		{"ceihma n adpouae s rtllrt/gywwgy w wywggyw w wgwyyg", 751},
		{"benids i nwlaiat l tdgiao/gwywgy y ywgggww w wgwywg", 750},
		{"tnhona r nbciior c elaune/gwyygy w wgwgwww w yggwyg", 749},
		{"lgteee i lvpseii o ddaere/gywygw w wwygywg w ggwywg", 748},
		{"snaoer e nirieev r vlsrty/gwwygw y wyggwgw w ygwywg", 747},
		{"loesde l oiaiekr t wrlrex/gwywgy y ywwgwww w wgyyyg", 746},
		{"srvdej i lgeohil b stycrt/gwyggw y wgygwww w wgywyg", 745},
		{"syurpa i neelkuu y epmare/gwwwgw y wywgwyg y ggwywg", 744},
		{"aerntu b uiliwhu o feghgt/gwwwgw w wwygwyg y ggyywg", 743},
		{"cnnafe i nllnlir l uglfiy/gwywgw y wywgwyw w wgywyg", 742},
		{"sasshv s itsaaal t iewlmy/gyyggw y gywgyww w wgwwwg", 741},
		{"salcla e edyieaw e tlftvr/gwywgy y yyygyyw w wgwgwg", 740},
		{"pature l rytcamn h upteey/ggywgw w ywygwww g wgyywg", 739},
		{"riceai i eugnifl a urgbmn/gwywgw y wyygyyw y wgwwwg", 738},
		{"bpcoha j eaansnr i asudiy/gwywgg y wwggyww y ygwwwg", 737},
		{"seoalu c cetfrat o lbihfh/gwgwgy w yyygyyw w wgwwwg", 736},
		{"baartu o lmsdera i heisue/gwyygw w wywgwwy w ygywwg", 735},
		{"vdiedt l aisltav i proiiy/gwywgw y wyygyyw w wgwwwg", 734},
		{"pkaanb k lnwarew n atcerd/gwgwgw y wywgwyw w ygwgyg", 733},
		{"hisehe e eetmtaa r nlpoar/gwywgy y yyygyyw y wgwwwg", 732},
		{"toebht c okeacea r uolicr/ggywgw y wywggyy y wgwwwg", 731},
		{"ssgalu u iosinlo p phmiun/gwwwgy y ywgggww y wgywyg", 730},
		{"satcnn e ribotnl e ittavh/gyywgy y wwwgwgy w ygyywg", 729},
		{"frewee l uugnaoo g rteout/gwywgw w wywgwyy y ygwywg", 728},
		{"scdapr e mrsdest o apiuoe/ggywgw y wywggyy y wgwwwg", 727},
		{"lhardc o iitceuu a scyoum/gwywgw g wyygyyw w wgwywg", 726},
		{"cuhecn l erebmni i ifooed/gwwwgw y ywygywy w wgwyyg", 725},
		{"caoasi e osottra r mcneut/gwwwgy y ywgggww y wgwywg", 724},
		{"edgatc r piiicsl n belare/ggywgw w yywgwww g wgyywg", 723},
		{"snlado l cpaeocd c vepare/gwywgy w yyygyyw w wgwgwg", 722},
		{"trrrln h etcgach o iwialh/ggywgw w wywgwyy y wggywg", 721},
		{"fespas e tllsoll o thnray/gwwwgy y yywgwyw y wgwwwg", 720},
		{"quailn r loktepu e alnoly/ggywgw w wywgwyy g wgywwg", 719},
		{"ieotoo t getatso o lrenpr/gwywgw y wyygyyw w wgwwwg", 718},
		{"llikne l enseuai g ttshay/gwyygw y wwygyyy w ggwwyg", 717},
		{"brnhdu n eerdruu u dtanom/gwywgy w yywgwyw y wgwwwg", 716},
		{"hoabre a gvdoiir l atseie/ggywgw w ywygwyw g wgyywg", 715},
		{"snagwc o henirlr h hpehle/gwwwgg y gywgwyw y wgwywg", 714},
		{"coosrh c ndekaaw e ohieoy/gwgwgw w wwygyww g ygwyyg", 713},
		{"tteulr u oaegrwe l cmncih/gwywgw w wwygywg y ggwywg", 712},
		{"fuutrl i oeuraat l udsily/gwywgw w wyygwwy y ygwwgg", 711},
		{"dskuye c eiweruu p olkews/gyyygy w ywygyww y wgwywg", 710},
		{"gruese u siancnu o wtntsk/ggwwgy g wywgwyy y wgwwwg", 709},
		{"llamrg e caeerag n pleyah/gwywgy w yyygyyw w wgwywg", 708},
		{"tnnoar m zkaaigw e ngaiey/gwwwgw g wwyggyy y ygwwyg", 707},
		{"aroame t rcuoral e nrdatl/gywygw w wyygyyw w wgwywg", 706},
		{"mioynu t wlsacor r esauin/gwywgg g wywgyyw y wgwwyg", 705},
		{"brcodd o iuornta g ahpdey/gywygw y wywgwyw w wgwywg", 704},
		{"lmihti r uiyrvms e aritee/gyywgg y wywgyww y ygwwwg", 703},
		{"kkehis u asaiebl s ibglne/gyyygw w wywgwyw w wgyyyg", 702},
		{"ggsipm d sreiaas s asrtse/gwywgw y yywgyyg w wgwwwg", 701},
		{"smeite h dtoiiae l ttlewr/gwgwgw w wyygyyw y wgwgwg", 700},
		{"eonits l ytrenah x rcaiat/gwywgw w wywgggy w ygywwg", 699},
		{"fremtp e ceaatsl t aemany/gwywgw g wwygywy w ygywyg", 698},
		{"ardatp e ntiomhl e btjrre/gwyygw y wwygywy w ggwwgg", 697},
		{"pcthsa r teeasri x emserr/gwywgw w wywgwyy y ygwywg", 696},
		{"ccrykc s uepalgr u aecmht/gywwgy g ywygwww y ygywwg", 695},
		{"pnuael t mlzpoir h teaxil/gwywgg y gywgwyw w wgwywg", 694},
		{"drlall r gidedfh a atluee/ggywgw w ywygwww g wgyywg", 693},
		{"nrbnlv a oatsail i eyaien/gwwwgw y wywgwyg y ggywyg", 692},
		{"sblelx a rprteaa e edhkly/gwywgw y wyygygw w ygwwwg", 691},
		{"ttoepo x aritaap c rretoy/gwgwgw w wwygywy y ygwgwg", 690},
		{"cahtnt w iyughon r uctity/gwywgw w wywggwy w ygywgg", 689},
		{"leleeh g ungidag y aeentr/gwwwgw y wyygyyw w wgyyyg", 688},
		{"hienyi u omtsric k udatuy/gyywgy w wgwgwyw w ygwgwg", 687},
		{"ldtano s aealrva t uosese/gywygw w wwygywy g ygwywg", 686},
		{"rodati e eaetfsv o ntgroy/ggwwgg y wwygwwy y wgywwg", 685},
		{"daexlr e mwliroi d wflpuy/gwgwgy y yywgwyw g wgwwwg", 684},
		{"fsnnkp a rdoixax l eseeld/gwwwgw w wyygyyw g ygwgwg", 683},
		{"nlowyr u tpsamsh a alefmy/gywygw w wyygyyw w wgwgwg", 682},
		{"rsnpeo e tseuqra i urrsay/gywwgw w ygwgyyw w wgwygg", 681},
		{"fctnen v rmxilvu n eeaeet/gywygw y wywgwyy y ygwgwg", 680},
		{"cgatla h aoaoioa a ynnbsl/gwywgy w ywygwyw w ygwywg", 679},
		{"roeite e seepksc s wyaoet/gwywgw g wywgwyy y ygywyg", 678},
		{"ttmuep u rpmicmo d mtariy/gwywgw y wyygyyw y wgwwwg", 677},
		{"uionni l hnnltih c iteiae/gyyygy w yywgwyw w wgwywg", 676},
		{"arioyp d ddeipea l utirel/gwwwgw y ygygwyw w wggywg", 675},
		{"ksuikt m raotucp o tyicia/gywygy w ywygyww y wgwgwg", 674},
		{"srrumn u etrgaar s ehwcuy/gyywgw g gyygyww w wgwwyg", 673},
		{"pbcrea o eliiudd d mazlbt/gywygy y yywgwyw y wgwwwg", 672},
		{"betamp s celeonc l krotah/gwwwgw y yyygwyw g wgwgwg", 671},
		{"smttdl a anuulic l aptsey/gwywgg y gwygyww w wgywyg", 670},
		{"roefhg i kunirle n gegluy/ggywgw w ywygwww g wgyywg", 669},
		{"dreeru i oeovoel n emonen/gwywgw w wywgwyw y wgywyg", 668},
		{"meatrn v eoouire t aftdcy/gyywgw y gywgwyw y wgywwg", 667},
		{"senogv u vredigd e lniune/gwywgy y ywygyww w wgwywg", 666},
		{"adleto e loruhel m htssve/gwyygy y wywgwgy w ygwwwg", 665},
		{"mdrmod h lobuoie a rldbty/gwywgy g yywgwyw w wgwgwg", 664},
		{"liuose t asccamo a trhure/gwyygg y ywwgyyw y ygwwwg", 663},
		{"smdiyo w gnalwld e enarpy/gyyygw w wwygyww g wgwgwg", 662},
		{"eoiaya l nchnioe y lsevtd/gyywgw y wyygwww g ggwywg", 661},
		{"wirato a oeeiigs s nrdtco/gywygw w wywgwyy w ygwgwg", 660},
		{"lrnsns o auessra i eriohe/gyywgy y wywgywy w wggwwg", 659},
		{"odrlne u dtrieei w tdecey/gyyygy w ygwgwgw w wgwwwg", 658},
		{"breihn e hntboro m utrtna/ggywgy w wywgyyw y wgwywg", 657},
		{"sckeep p eurcaxg l oekuoy/gwywgw w wywgwyy w ygwywg", 656},
		{"sornfp m gceigke r nlaley/gwywgg y ywygyyy w wgwwyg", 655},
		{"rietpd t eenmdao o aweacy/gwywgw y wyygyyw y wgwwwg", 654},
		{"akbotn a exnilab t stiegn/gwwwgw w ywygyyy y wggywg", 653},
		{"cekuei e eohtapa c hcbdjt/gwywgy y ywygyww w wgwgwg", 652},
		{"sklare t kisydfb l sfloiy/ggywgw w ywygwww g ygyywg", 651},
		{"tuarpm x ninbroi n hpueme/gywygy y ywygyww y wgwwwg", 650},
		{"slarll g tmtreoa e vloare/gygwgy w wywgyyw w wgwgwg", 649},
		{"slaehi e acaorsk w invwme/gwgwgy y ywygyww y wgwwwg", 648},
		{"pogphc a godiwel n eruoee/ggywgw w ywygwyw g wgyywg", 647},
		{"svrblc e tsthiua o nenaah/gwywgw w wywgwyy g ygwywg", 646},
		{"tvwuts a rurflin t csloie/gwywgy y wgwgwyw w wgyywg", 645},
		{"lnnyge s ustaime s reavny/gyyygg w gywgwyw w wgwywg", 644},
		{"awdetp e iebrhht t bcoiuy/gwyygw g wwygywy w ggwywg", 643},
		{"psaohc r boebsar d feeiry/gwywgw y wyygyyw w wgwgwg", 642},
		{"flrohs r durdnsl e etehoe/ggywgy w ywygwww g wgyywg", 641},
		{"calose e lhlpmih x ameesr/gwywgw y wwygyww y wgwywg", 640},
		{"praiel i rirbste e oldaer/ggwwgw y wyygwwy w wgwygg", 639},
		{"sloane n eetmunt o ntchoh/gyyygw y wyygyyw g wgwwwg", 638},
		{"acdgnr n nonovsr i itdeay/ggywgw w ywygwww g wgyywg", 637},
		{"csalke l keeival l npraod/gwywgw g wywgwyy w ygwywg", 636},
		{"gswoni e errdaor p ehrvde/gwyygw y gywgyww w wgywwg", 635},
		{"mlakll m eobaule o mraraa/gwywgw y wywgwyg w ggwywg", 634},
		{"coatrr f eyafeaa w uraoel/ggywgw w wyggwyy y wgywwg", 633},
		{"sayalo c etcaerv t rpwate/gwywgw w wyygyyw w wgwywg", 632},
		{"sehdtn e uaaooeb l tknprl/gyywgy y wgwgwww w yggwwg", 631},
		{"cuelna d oieveiv n aroair/gwywgw w wywgwyy y ygwywg", 630},
		{"oewear g rrlaced m nrlety/gywwgg w wyygwyy w wgwyyg", 629},
		{"sfvdta r pieaehe c rtghnt/gwwwgw g wwygywy g ygwywg", 628},
		{"sorsgg n nwoteii u ngonre/gwywgw g wyygwyy w wggwyg", 627},
		{"rnisrr h mnsoeaa c gtoeid/gwwwgw g wywgwyy y ygwgwg", 626},
		{"bespeo j ureniin i rseale/ggwwgg w wwwgywy y ygywwg", 625},
		{"heisyi y cosnraw h dtoidt/gwywgy y yywgwyw w wgwwwg", 624},
		{"eebohe t mecsotn o areepr/gwwygy y wywgwgw y ygwywg", 623},
		{"feighg c aonoliw s tloile/gwywgw w wyygyyw g wgwywg", 622},
		{"vealeu a teiaefl y edsarr/gwyygw w yywgwwy y ggwwyg", 621},
		{"tnoens a ceiilob o rhnhoy/gwgwgw w wyygyyw g wgywyg", 620},
		{"soalth a ngaegcn a irugue/gywwgw w ygygwww g ygwywg", 619},
		{"bnathp c etotree r imoeay/gwywgw w wyygyyw w wgwywg", 618},
		{"grfief m euanuia m aaasel/gwgwgw w ywwggyy w ygywwg", 617},
		{"lgnsha o iiemrar a tnuech/gwywgw w wyygyyw y wgwywg", 616},
		{"sontew o tnreveo t impnrr/gwywgw w ywygyyy y wgwwwg", 615},
		{"terogn n niolnyh i ocigue/gwwwgy w yyygyyw y wgwywg", 614},
		{"trwlht p ittailu a epeosy/ggywgw w wywgwyy y wggwwg", 613},
		{"wccagt u iaeslpn t yhnnrh/gwwwgw y wywgwyy y ygwywg", 612},
		{"aicade i ubrneoo v wterrd/gywwgw w wywgyyy w wgygwg", 611},
		{"snmlgs r oembene i oculie/gyyygw w wywgwyw y wgwwwg", 610},
		{"ciiaci h enmoflv a npdllt/ggywgw w yywgwyw y wgwywg", 609},
		{"nieehc l nolvaar d eroety/gwywgy y yywgwyw w wgwywg", 608},
		{"clnupu e ielcmrs g udrire/ggwygw y wywgwgw w ygyywg", 607},
		{"vdedta a iitppiu l irteiy/gwywgw w wyygyyw y wgwwwg", 606},
		{"shuigr r agiduew e rkontl/ggywgy y wwygyww y ygwywg", 605},
		{"finenl r ydeitop o ntmpwh/gwywgg w gywgwyw y wgwywg", 604},
		{"uprcre m ubeizen p wtloed/ggywgw w wywgwyy y wgwwgg", 603},
		{"ccraea i sgsimln n hesume/gwywgy y ywygyww g wgwwwg", 602},
		{"allare w atlnosn a uhloey/ggywgw w ywygwww g wgyywg", 601},
		{"rnnudr t myeivoi h lrldee/gyyygw y wwwgwwy w ygwwwg", 600},
		{"vbsmtt e itiyhea a ralaod/gwwygy y wywgwgw y ygwywg", 599},
		{"sohcnk t eottute e rlaxey/gwywgw w wgwgwgw y wgyyyg", 598},
		{"mdluci o piamiye r ocnmby/gwwwgy y wwwgwgy w ygyywg", 597},
		{"kemaln e ragakil r ibfbey/gwywgw y wywgwyw w wgywyg", 596},
		{"lbgeha e aurdrur e adaddy/gwywgw w gyygywy g wgwywg", 595},
		{"boqndr e apsutat p astiiy/gwywgw y wwygyww y wgwywg", 594},
		{"prneer u aisnetl i etnlce/ggywgw w wywgyyy y wgwwwg", 593},
		{"svhoep p leetrsn f utoode/gyyygg w gwwgwww y wgwywg", 592},
		{"exlglo i etmorhl l otoaun/ggywgw w ywygwww g wgyywg", 591},
		{"coelhd a htlixol e asxdiy/gyyygw w wwygyww w wgwywg", 590},
		{"gterma s adcarae a lpwxel/gwgwgy w yywggyw w wgwyyg", 589},
		{"ptuiee m lmutenr r nblemy/gwgwgw w wyygyyw y wgwywg", 588},
		{"ianreg t innaiag r gttony/gyyygy w wywgyww w wgwwyg", 587},
		{"raebts r uesilts l ltewey/gwywgw y wyygyyw y wgwywg", 586},
		{"drvaee t feeatrr o wtbxvt/ggywgw y wywggyy y wgwwwg", 585},
		{"wxtaea i ecigamc a hnniol/gwywgy g yyygyyw w wgwwwg", 584},
		{"lrheel r siaagts h stiaiy/gwywgw y gywgyww w ygwgwg", 583},
		{"qcttee u eiehrue u atsigt/gwywgy w ywygyww w wgwywg", 582},
		{"chiota r reeomsz j nsrhhe/ggywgw w yywgwww g wgyywg", 581},
		{"tlepoi i stagamo r erihec/gwywgw w wywgwyw y wgywyg", 580},
		{"hoivrn t osesori c etecoh/ggyygw w wyygwwy w ygwyyg", 579},
		{"tnoahe g eioioet u hrcnvl/gwywgy w yywgwyw y wgwgwg", 578},
		{"sasgka u hrtaeae y cwardk/gwywgy y wgwgyyw w wgyywg", 577},
		{"daobrm o uuonimg r dteiey/gwywgw w wyygyyw y wgwywg", 576},
		{"dorayi a liylvia e erarnn/gwyygg y wywgyww w yggwwg", 575},
		{"tperde c itlxiio i ecmiia/gwywgw w wwygywy y ygwywg", 574},
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
		// ...
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
		// ...
		{"agdrml r ueianeu l oeibrr/gwywgw y wwygywg w ggwywg", 300},
		// ...
		{"lieirs w riipese n rcouye/gwywgw y wygggyw w wgwywg", 200},
		// ...
		{"mmkoye u iaomerr o pahcln/gwywgw y wyggwyw w wgwwwg", 100},
		// ...
		{"ldraly s oaoreat w atsloy/gwwggw y ggygwww w wgywwg", 50},
		{"dsathp i ebiseet n otecuh/gwwggw w wgygwww w wggwwg", 49},
		{"mhsioo i eruusbt t rlacre/gwyygw w ygwgwyw y wgwywg", 48},
		{"mrhyac c gniaiai i tcetfc/gwywgw y wwwgwwg g ygwwwg", 47},
		{"srduni e riocaer e ncivey/gwwwgy w yywgwyw w wgwywg", 46},
		{"fnrhta r rlagoos o urtegy/gwywgw w wyygyww y wgyywg", 45},
		{"sncemn c neligin o uprrae/gwywgw w ywwgwyw w wgwywg", 44},
		{"tsperk f aexohae y itkoan/gwwggw w ggygwww w wgywyg", 43},
		{"klados a tehstcy e aikzol/gwywgw w yywgyyw y ggwywg", 42},
		{"meeleo r aeatech a troenr/ggyygy y wwwgwyw y ygwywg", 41},
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
		{"arcxet b ezmonje n aezodt/gwwwgw w yywggyw g wgwywg", 12},
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

type Stats struct {
	minPuzzle int
	maxPuzzle int
	dict      map[string]bool
	newWords  map[int]int
}

var (
	stats = Stats{
		0,
		0,
		map[string]bool{},
		map[int]int{},
	}
)

func recordPuzzle(index int, words []string, solved bool) {
	stats.minPuzzle = min(stats.minPuzzle, index)
	stats.maxPuzzle = max(stats.maxPuzzle, index)
	stats.newWords[index] = 0

	if !solved {
		stats.newWords[index] = 0
		return
	}
	for _, word := range words {
		if !stats.dict[word] {
			stats.newWords[index]++
		}
		stats.dict[word] = true
	}
}

func printStats() {
	fmt.Println()
	fmt.Println("Index,NewWords")
	for i := stats.minPuzzle; i <= stats.maxPuzzle; i++ {
		newWords, ok := stats.newWords[i]
		if !ok {
			newWords = -1
		}
		fmt.Printf("%d, %d\n", i, newWords)
	}
}

func dailyStats() {
	testCases := dailyWaffles
	for i := len(testCases) - 1; i >= 0; i-- {
		waffle := board.Parse(testCases[i].serial)
		s := solver.New(waffle)
		s.Solve()
		recordPuzzle(testCases[i].index, s.Words(), s.Solved())
	}
}

func TestSolve(testCases []TestCase) {
	total := 0
	count := 0
	for _, testCase := range testCases {
		waffle := board.Parse(testCase.serial)
		s := solver.New(waffle)
		if !s.Solve() {
			fmt.Println("Unable to solve:", testCase.serial)
			s.Print()
			fmt.Println()
			continue
		}
		path := pathfinder.New(s)
		path.Find()
		count++
		total += path.PathLen()
		fmt.Printf("Game: %3d Steps: %3d Average: %3.2f\n", testCase.index, path.PathLen(), float64(total)/float64(count))
	}
}

func TestParseSolution(testCases []TestCase) {
	for _, testCase := range testCases {
		waffle := board.Parse(testCase.serial)
		s := solver.New(waffle)
		if !s.Solve() {
			fmt.Println("Unable to solve:", testCase.serial)
			s.Print()
			fmt.Println()
			continue
		}
		puzzle := strings.Split(testCase.serial, "/")[0]
		solution := s.Serialize()
		signature := solver.ParseSolution(puzzle, solution)
		if signature != testCase.serial {
			s.Print()
			fmt.Printf("Failed to generate expected signature\n  Expected: %s\n  Got:      %s\n", testCase.serial, signature)
		}
		fmt.Printf("Game: %3d passes\n", testCase.index)
	}
}

func main() {
	statsOnly := flag.Bool("stats-only", false, "Only print word count stats")
	flag.Parse()

	if *statsOnly {
		dailyStats()
		printStats()
		return
	}

	fmt.Printf("Welcome to waffle regression tests!\n")

	fmt.Printf("\n---------- Quick sanity check ----------\n")

	fmt.Printf("\nDaily Waffles - ParseSolution\n")
	TestParseSolution(dailyWaffles[:2])

	fmt.Printf("\nDeluxe Waffles - ParseSolution\n")
	TestParseSolution(deluxeWaffles[:2])

	fmt.Printf("\nDaily Waffles\n")
	TestSolve(dailyWaffles[:2])

	fmt.Printf("\nDeluxe Waffles\n")
	TestSolve(deluxeWaffles[:2])

	fmt.Printf("\n---------- Exhaustive run of all tests ----------\n")

	fmt.Printf("\nDaily Waffles\n")
	TestSolve(dailyWaffles)

	fmt.Printf("\nDeluxe Waffles\n")
	TestSolve(deluxeWaffles)

	fmt.Printf("\nDaily Waffles - ParseSolution\n")
	TestParseSolution(dailyWaffles)

	fmt.Printf("\nDeluxe Waffles - ParseSolution\n")
	TestParseSolution(deluxeWaffles)
}
