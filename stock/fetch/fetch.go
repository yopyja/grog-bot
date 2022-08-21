package fetch

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func Item(sku string) {
	file, err := os.Create("./json/temp.json")

	url := "https://webapps2.abc.utah.gov/ProdApps/ProductLocatorCore/Products/LoadProductTable"
	method := "POST"

	payload := strings.NewReader("draw=13&columns%5B0%5D%5Bdata%5D=name&columns%5B0%5D%5Bname%5D=&columns%5B0%5D%5Bsearchable%5D=true&columns%5B0%5D%5Borderable%5D=true&columns%5B0%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B0%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B1%5D%5Bdata%5D=sku&columns%5B1%5D%5Bname%5D=&columns%5B1%5D%5Bsearchable%5D=true&columns%5B1%5D%5Borderable%5D=false&columns%5B1%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B1%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B2%5D%5Bdata%5D=displayGroup&columns%5B2%5D%5Bname%5D=&columns%5B2%5D%5Bsearchable%5D=true&columns%5B2%5D%5Borderable%5D=true&columns%5B2%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B2%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B3%5D%5Bdata%5D=status&columns%5B3%5D%5Bname%5D=&columns%5B3%5D%5Bsearchable%5D=true&columns%5B3%5D%5Borderable%5D=true&columns%5B3%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B3%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B4%5D%5Bdata%5D=warehouseQty&columns%5B4%5D%5Bname%5D=&columns%5B4%5D%5Bsearchable%5D=true&columns%5B4%5D%5Borderable%5D=true&columns%5B4%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B4%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B5%5D%5Bdata%5D=storeQty&columns%5B5%5D%5Bname%5D=&columns%5B5%5D%5Bsearchable%5D=true&columns%5B5%5D%5Borderable%5D=true&columns%5B5%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B5%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B6%5D%5Bdata%5D=onOrderQty&columns%5B6%5D%5Bname%5D=&columns%5B6%5D%5Bsearchable%5D=true&columns%5B6%5D%5Borderable%5D=true&columns%5B6%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B6%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B7%5D%5Bdata%5D=currentPrice&columns%5B7%5D%5Bname%5D=&columns%5B7%5D%5Bsearchable%5D=true&columns%5B7%5D%5Borderable%5D=true&columns%5B7%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B7%5D%5Bsearch%5D%5Bregex%5D=false&order%5B0%5D%5Bcolumn%5D=0&order%5B0%5D%5Bdir%5D=asc&start=0&length=50&search%5Bvalue%5D=&search%5Bregex%5D=false&item_code=" + sku + "&item_name=&category=&sub_category=&price_min=&price_max=&on_spa=false&status=")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept-Language", "en-US,en;q=0.9")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Length", "1932")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Host", "webapps2.abc.utah.gov")
	req.Header.Add("Origin", "https://webapps2.abc.utah.gov")
	req.Header.Add("Referer", "https://webapps2.abc.utah.gov/ProdApps/ProductLocatorCore")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("Sec-GPC", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.81 Safari/537.36")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Cookie", "AuthToken=145002252188171174114090038063080226241224035089178225017239118164014043135185096234243017147236023036173089169009243218019104186154126135216059190166088172237236055088006018056017168117205028132242044241137245029136020251115025212088067047035116228250119076090228216209147038096173197098091022038131194108068105076080058026021157189016135090147068084111170089119103154132107195188106082088215008173025125145124195098166162010206003187144245171014082156245008102084151041000032122171036114138244003215254025048174137205040114214137103159248175153083185016199251088019244131006217121078051185124167106212146248148165180238208139089017137075126115208040048071054212191227015018118073002087193101064202214193254039046172171023051059073136238245001038209139249168168190199219080090003027169208187223169195099241239050144215190101232236032070086026199075200235088067230021100193242160011058216092056231226039255069073069110068153218106088221242153190113214066185197044101093038043197145218204018114104170140120219146137208125008001055158221125110149110217102209201161236188133153090189142132072248003167146053229122222163196045180184003013081041170137121187065095180092086168150214168092187247136090213253064117027102089112058201012228000041185039200125118252160177186108144180146091113038018184096062110223157141065175061096067125248204030231130138235078154206001175171194187109009203207189062064219068084173122182213037117202159064068235135226003182099006175140221014032132090100072232203172177124112073252143212179088119144144200185222127228185147023246147037038014014158142188126198048114116039144127115021169083231137060208102038132149229243010138208112122024093244171218180136214251208068019227046151056131201247119191031055110218124144031107246189017222224079107158044220120052180044016076135015010216023119154034250095236098189073030161121077148226024141089108205098167214114119041046196233248187024050061214048120170121169227206109009049227192125219011199143130068105165149154192161032133213014230042128248000194067093059189212156016236117169161205130055158156163250146099083096047073150241146025157216136052096123160149023165068029080246048054230057141087109150206100220215212091048076036025058184180192076235077213116040069174029068197148231230216095069179041247141049030002010228100110089236199099; filter=145161011038246168217042088014084086236240085157201087017229014040225188225188255055123193179020049146197023200140223137096189162074010074076057210110140252151174212081109038201072204055008097152204218032061119194200007031245040183196150118056062001084068127049137002108176104241175201039116174117191163106061201158249228124171139015254217082207164176124015153084180097248213151137133199140040154151138100205130005094084117021007227151020065242141247069102214150055002041135223078240202035160130192127116194176241171039101091220075076031240138250004119176005042107144006253202224122035190033037070222113088019032094244114126231176162213130187211249197174230035106015145245060078020226157074164091042235154114030147218132")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err != nil {
		log.Fatal(err)
	}

	mw := io.MultiWriter(os.Stdout, file)
	fmt.Fprintln(mw, string(body))
}