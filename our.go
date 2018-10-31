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
	jsonHandler := http.FileServer(http.Dir("./JDB/"))
	jsonHandler = http.StripPrefix("/json/", jsonHandler)
	r := mux.NewRouter()
	r.PathPrefix("/json/").Handler(jsonHandler)

	//	r.Host("com-http.us").Path("/").HandlerFunc(rts.NXCoinsHandler).Name("nxcoins")
	//r.Host("com-http.us").Path("/").HandlerFunc(rts.IndexHandler).Name("index")

	// r.Host("com-http.us").Path("/").Handler(httputil.NewSingleHostReverseProxy(&url.URL{
	// 	Scheme: "http",
	// 	Host:   "127.0.0.1:7227",
	// })).Name("index")

	target := "http://127.0.0.1:3553/"
	remote, err := url.Parse(target)
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)

	r.Host("com-http.us").Path("/").HandlerFunc(rts.NXIndexHandler(proxy)).Name("nxindex")
	r.Host("{coin}.com-http.us").Path("/").HandlerFunc(rts.NXCoinHandler(proxy)).Name("nxcoin")

	//	r.Host("com-http.us").Path("/home").HandlerFunc(rts.HomeHandler).Name("home")
	//	r.Host("{coin}.com-http.us").Path("/").HandlerFunc(rts.CoinHandler).Name("coin")

	// r.Host("{coin}.com-http.us").Path("/explorer").HandlerFunc(rts.ExplorerHandler).Name("explorer")
	// r.Host("{coin}.com-http.us").Path("/block/{id}").HandlerFunc(rts.ViewBlockHeight).Name("block")
	// r.Host("{coin}.com-http.us").Path("/blockhash/{id}").HandlerFunc(rts.ViewBlockHash).Name("blockhash")
	// r.Host("{coin}.com-http.us").Path("/tx/{id}").HandlerFunc(rts.ViewTx).Name("tx")
	// r.Host("{coin}.com-http.us").Path("/addr/{id}").HandlerFunc(rts.ViewAddr).Name("addr")

	r.Host("{coin}.com-http.us").Path("/explorer").HandlerFunc(rts.NXExplorerHandler(proxy)).Name("nxexplorer")
	r.Host("{coin}.com-http.us").Path("/explorer/block/{id}").HandlerFunc(rts.NXBlockHandler(proxy)).Name("nxblock")
	r.Host("{coin}.com-http.us").Path("/explorer/hash/{id}").HandlerFunc(rts.NXHashHandler(proxy)).Name("nxhash")
	r.Host("{coin}.com-http.us").Path("/explorer/tx/{id}").HandlerFunc(rts.NXTxHandler(proxy)).Name("nxtx")
	r.Host("{coin}.com-http.us").Path("/explorer/addr/{id}").HandlerFunc(rts.NXAddrsHandler(proxy)).Name("nxaddr")

	r.Host("{coin}.com-http.us").Path("/explorer/search").HandlerFunc(rts.DoSearch).Name("search")
	//api
	r.Host("{coin}.com-http.us").Path("/a/last").HandlerFunc(rts.ApiLast).Name("last")
	r.Host("{coin}.com-http.us").Path("/a/info").HandlerFunc(rts.ApiInfo).Name("info")
	r.Host("{coin}.com-http.us").Path("/a/peer").HandlerFunc(rts.ApiPeer).Name("peer")
	r.Host("{coin}.com-http.us").Path("/a/mining").HandlerFunc(rts.ApiMiningInfo).Name("mining")
	r.Host("{coin}.com-http.us").Path("/a/rawpool").HandlerFunc(rts.ApiRawPool).Name("rawpool")

	r.Host("{coin}.com-http.us").Path("/a/{type}/{id}").HandlerFunc(rts.ApiData).Name("coin")

	r.Host("{coin}.com-http.us").Path("/a/news").HandlerFunc(rts.CoinNewsHandler).Name("news")
	r.Host("{coin}.com-http.us").Path("/f/cmc").HandlerFunc(rts.CMCHandler).Name("cmc")

	r.Host("{coin}.com-http.us").Path("/favicon.ico").HandlerFunc(rts.IcoHandler).Name("ico")

	r.Host("i.com-http.us").Path("/{coin}/{size}").HandlerFunc(rts.ImgHandler).Name("img")

	r.Host("c.com-http.us").Path("/c").HandlerFunc(rts.CoinsHandler).Name("coins")
	r.Host("c.com-http.us").Path("/madness").HandlerFunc(rts.CoinsMadnessHandler).Name("cmdns")

	r.Host("com-http.us").Path("/cert").HandlerFunc(rts.CertHandler).Name("cert")

	r.NotFoundHandler = http.HandlerFunc(rts.FOFHandler)

	go log.Fatal(http.ListenAndServe(":80", handlers.CORS()(handlers.CompressHandler(r))))

	// err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	// if err != nil {
	// 	go log.Fatal("ListenAndServe: ", err)
	// }

}
