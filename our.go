/*************************************
    __    __    __    __    __    __
   /  \  /  \  /  \  /  \  /  \  /  \
  /  __\/  __\/  __\/  __\/  __\/  __\
 /  /__/  /__/  /__/  /__/  /__/  /__/
  \   / \   / \   / \   / \   / \   /
   \_/   \_/   \_/   \ /   \_/   \_/

App : our
Name: Ouroboros

*************************************/
package main

import (
	"fmt"
	"time"

	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/parallelcointeam/our/rts"
	"github.com/robfig/cron"
)

func main() {
	cr := cron.New()
	cr.AddFunc("@every 121215s", func() {
		fmt.Println("Radi kron")
	})
	cr.Start()

	r := mux.NewRouter()

	//	r.Host("com-http.us").Path("/").HandlerFunc(rts.NXCoinsHandler).Name("nxcoins")
	//r.Host("com-http.us").Path("/").HandlerFunc(rts.IndexHandler).Name("index")

	// r.Host("com-http.us").Path("/").Handler(httputil.NewSingleHostReverseProxy(&url.URL{
	// 	Scheme: "http",
	// 	Host:   "127.0.0.1:7227",
	// })).Name("index")

	atarget := "http://127.0.0.1:3553/"
	aremote, err := url.Parse(atarget)
	if err != nil {
		panic(err)
	}
	amprx := httputil.NewSingleHostReverseProxy(aremote)

	// ftarget := "http://127.0.0.1:3557/"
	// fremote, err := url.Parse(ftarget)
	// if err != nil {
	// 	panic(err)
	// }
	// ifsprx := httputil.NewSingleHostReverseProxy(fremote)

	r.Host("com-http.us").Path("/").HandlerFunc(rts.NXIndexHandler(amprx)).Name("nxindex")
	r.Host("com-http.us").Path("/coins").HandlerFunc(rts.NXCoinsHandler(amprx)).Name("nxcoins")
	r.Host("com-http.us").Path("/words").HandlerFunc(rts.NXWordsHandler(amprx)).Name("nxwords")
	r.Host("{coin}.com-http.us").Path("/").HandlerFunc(rts.NXCoinHandler(amprx)).Name("nxcoin")

	//	r.Host("com-http.us").Path("/home").HandlerFunc(rts.HomeHandler).Name("home")
	//	r.Host("{coin}.com-http.us").Path("/").HandlerFunc(rts.CoinHandler).Name("coin")

	// r.Host("{coin}.com-http.us").Path("/explorer").HandlerFunc(rts.ExplorerHandler).Name("explorer")
	// r.Host("{coin}.com-http.us").Path("/block/{id}").HandlerFunc(rts.ViewBlockHeight).Name("block")
	// r.Host("{coin}.com-http.us").Path("/blockhash/{id}").HandlerFunc(rts.ViewBlockHash).Name("blockhash")
	// r.Host("{coin}.com-http.us").Path("/tx/{id}").HandlerFunc(rts.ViewTx).Name("tx")
	// r.Host("{coin}.com-http.us").Path("/addr/{id}").HandlerFunc(rts.ViewAddr).Name("addr")

	r.Host("com-http.us").Path("/a/coins").HandlerFunc(rts.CoinsAMP).Name("coinsamp")
	r.Host("com-http.us").Path("/a/coinsimg").HandlerFunc(rts.CoinsAMPimg).Name("coinsampimg")
	r.Host("com-http.us").Path("/a/bitnodes").HandlerFunc(rts.CoinsBNAMP).Name("coinsampbitnodes")

	r.Host("{coin}.com-http.us").Path("/explorer").HandlerFunc(rts.NXExplorerHandler(amprx)).Name("nxexplorer")
	r.Host("{coin}.com-http.us").Path("/explorer/block/{id}").HandlerFunc(rts.NXBlockHandler(amprx)).Name("nxblock")
	r.Host("{coin}.com-http.us").Path("/explorer/hash/{id}").HandlerFunc(rts.NXHashHandler(amprx)).Name("nxhash")
	r.Host("{coin}.com-http.us").Path("/explorer/tx/{id}").HandlerFunc(rts.NXTxHandler(amprx)).Name("nxtx")
	r.Host("{coin}.com-http.us").Path("/explorer/addr/{id}").HandlerFunc(rts.NXAddrsHandler(amprx)).Name("nxaddr")

	r.Host("{coin}.com-http.us").Path("/explorer/search").HandlerFunc(rts.DoSearch).Name("search")

	r.Host("{coin}.com-http.us").Path("/network").HandlerFunc(rts.NXNetworkHandler(amprx)).Name("nxnetwork")

	r.Host("{coin}.com-http.us").Path("/price").HandlerFunc(rts.NXPriceHandler(amprx)).Name("nxprice")

	r.Host("{coin}.com-http.us").Path("/ecosystem").HandlerFunc(rts.NXEcoHandler(amprx)).Name("nxeco")

	r.Host("f.com-http.us").Path("/{frame}/{file}").HandlerFunc(rts.Frames(amprx)).Name("nxframes")

	//api
	//r.Host("{coin}.com-http.us").Path("/a/l").HandlerFunc(rts.ApiLast).Name("last")

	r.Host("{coin}.com-http.us").Path("/a/b").HandlerFunc(rts.ApiLastBlock).Name("b")
	r.Host("{coin}.com-http.us").Path("/a/bta/{id}").HandlerFunc(rts.ApiBlockTxAddr).Name("bta")

	r.Host("{coin}.com-http.us").Path("/a/i").HandlerFunc(rts.ApiInfo).Name("info")
	r.Host("{coin}.com-http.us").Path("/a/p").HandlerFunc(rts.ApiPeer).Name("peer")
	r.Host("{coin}.com-http.us").Path("/a/m").HandlerFunc(rts.ApiMiningInfo).Name("mining")
	r.Host("{coin}.com-http.us").Path("/a/r").HandlerFunc(rts.ApiRawMemPool).Name("rawmempool")

	r.Host("{coin}.com-http.us").Path("/a/n").HandlerFunc(rts.NodesHandler).Name("nodes")

	r.Host("{coin}.com-http.us").Path("/a/{type}/{id}").HandlerFunc(rts.ApiData).Name("coin")

	r.Host("{coin}.com-http.us").Path("/a/n").HandlerFunc(rts.CoinNewsHandler).Name("news")
	r.Host("{coin}.com-http.us").Path("/f/cmc").HandlerFunc(rts.CMCHandler).Name("cmc")

	r.Host("{coin}.com-http.us").Path("/favicon.ico").HandlerFunc(rts.IcoHandler).Name("ico")

	r.Host("i.com-http.us").Path("/{coin}/{size}").HandlerFunc(rts.ImgHandler).Name("img")

	r.Host("{coin}.com-http.us").Path("/frames/{frame}").HandlerFunc(rts.FrameHandler).Name("frame")

	// r.Host("c.com-http.us").Path("/c").HandlerFunc(rts.CoinsHandler).Name("coins")
	// r.Host("c.com-http.us").Path("/madness").HandlerFunc(rts.CoinsMadnessHandler).Name("cmdns")

	// r.Host("{coin}.com-http.us").Path("/__webpack_hmr").HandlerFunc(rts.NXWebPackHmr(ifsprx)).Name("nxwebpackhmr")

	// r.Host("{coin}.com-http.us").Path("/_nuxt/{nfile}").HandlerFunc(rts.NXNuxt(ifsprx)).Name("nxnuxt")
	// r.Host("{coin}.com-http.us").Path("/_nuxt/{nsub}/{nfile}").HandlerFunc(rts.NXNuxtSub(ifsprx)).Name("nxnuxtsub")
	// r.Host("{coin}.com-http.us").Path("/_nuxt/{nsub}/{nssub}/{nfile}").HandlerFunc(rts.NXNuxtSSub(ifsprx)).Name("nxnuxtssub")
	// r.Host("{coin}.com-http.us").Path("/_nuxt/{nsub}/{nssub}/{nsssub}/{nfile}").HandlerFunc(rts.NXNuxtSSSub(ifsprx)).Name("nxnuxtsssub")
	// r.Host("{coin}.com-http.us").Path("/_nuxt/{nsub}/{nssub}/{nsssub}/{nssssub}/{nfile}").HandlerFunc(rts.NXNuxtSSSSub(ifsprx)).Name("nxnuxtssssub")
	// r.Host("{coin}.com-http.us").Path("/_nuxt/{nsub}/{nssub}/{nsssub}/{nssssub}/{nsssssub}/{nfile}").HandlerFunc(rts.NXNuxtSSSSSub(ifsprx)).Name("nxnuxtsssssub")

	// r.Host("com-http.us").Path("/frame/nodes").HandlerFunc(rts.NXFrame(ifsprx)).Name("nxiframe")

	// r.Host("com-http.us").Path("/cert").HandlerFunc(rts.CertHandler).Name("cert")

	// r.PathPrefix("/_nuxt/").HandlerFunc(rts.NXNuxt(ifsprx))

	//r.Host("{coin}.com-http.us").Path("/frames/{frame}").HandlerFunc(rts.NXFrame(ifsprx)).Name("nxiframe")

	//	r.Host("libs.com-http.us").Path("/{libs}/{lib}").HandlerFunc(rts.NXLib(ifsprx)).Name("nxiframe")

	r.PathPrefix("/json/").Handler(http.StripPrefix("/json/", http.FileServer(http.Dir("./JDB/"))))

	r.Host("l.com-http.us").PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./static/libs/"))))

	r.NotFoundHandler = http.HandlerFunc(rts.FOFHandler)
	r.Schemes("https")

	//go log.Fatal(http.ListenAndServe(":80", handlers.CORS()(handlers.CompressHandler(r))))

	srv := &http.Server{
		Handler: handlers.CORS()(handlers.CompressHandler(r)),
		//Addr:    "com-http.us:443",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	go log.Fatal(srv.ListenAndServeTLS("cert.crt", "key.key"))
}
