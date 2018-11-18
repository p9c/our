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

	//	"time"

	"log"
	"net/http"

	//	"net/http/httputil"
	//	"net/url"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/parallelcointeam/our/rts"
	"github.com/robfig/cron"
)

const (
	// HTTPMethodOverrideHeader is a commonly used
	// http header to override a request method.
	HTTPMethodOverrideHeader = "X-HTTP-Method-Override"
	// HTTPMethodOverrideFormKey is a commonly used
	// HTML form key to override a request method.
	HTTPMethodOverrideFormKey = "_method"
)

func main() {
	cr := cron.New()
	cr.AddFunc("@every 121215s", func() {
		fmt.Println("Radi kron")
	})
	cr.Start()

	r := mux.NewRouter()

	r.Host("com-http.us").Path("/").HandlerFunc(rts.NXIndexHandler).Name("nxindex")
	r.Host("com-http.us").Path("/coins").HandlerFunc(rts.NXCoinsHandler).Name("nxcoins")
	r.Host("com-http.us").Path("/words").HandlerFunc(rts.NXWordsHandler).Name("nxwords")
	r.Host("{coin}.com-http.us").Path("/").HandlerFunc(rts.NXCoinHandler).Name("nxcoin")

	r.Host("com-http.us").Path("/a/coins").HandlerFunc(rts.CoinsAMP).Name("coinsamp")
	r.Host("com-http.us").Path("/a/coinsimg").HandlerFunc(rts.CoinsAMPimg).Name("coinsampimg")
	r.Host("com-http.us").Path("/a/bitnodes").HandlerFunc(rts.CoinsBNAMP).Name("coinsampbitnodes")

	r.Host("{coin}.com-http.us").Path("/explorer").HandlerFunc(rts.NXExplorerHandler).Name("nxexplorer")
	r.Host("{coin}.com-http.us").Path("/explorer/block/{id}").HandlerFunc(rts.NXBlockHandler).Name("nxblock")
	r.Host("{coin}.com-http.us").Path("/explorer/hash/{id}").HandlerFunc(rts.NXHashHandler).Name("nxhash")
	r.Host("{coin}.com-http.us").Path("/explorer/tx/{id}").HandlerFunc(rts.NXTxHandler).Name("nxtx")
	r.Host("{coin}.com-http.us").Path("/explorer/addr/{id}").HandlerFunc(rts.NXAddrsHandler).Name("nxaddr")

	r.Host("{coin}.com-http.us").Path("/explorer/search").HandlerFunc(rts.DoSearch).Name("search")

	r.Host("{coin}.com-http.us").Path("/network").HandlerFunc(rts.NXNetworkHandler).Name("nxnetwork")
	r.Host("{coin}.com-http.us").Path("/price").HandlerFunc(rts.NXPriceHandler).Name("nxprice")
	r.Host("{coin}.com-http.us").Path("/ecosystem").HandlerFunc(rts.NXEcoHandler).Name("nxeco")
	r.Host("f.com-http.us").Path("/{frame}/{file}").HandlerFunc(rts.Frames).Name("nxframes")

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

	r.PathPrefix("/json/").Handler(http.StripPrefix("/json/", http.FileServer(http.Dir("./JDB/"))))

	r.Host("l.com-http.us").PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./static/libs/"))))

	r.NotFoundHandler = http.HandlerFunc(rts.FOFHandler)
	r.Schemes("https")

	//go log.Fatal(http.ListenAndServe(":80", handlers.CORS()(r)))

	srv := &http.Server{
		Handler: handlers.CORS()(handlers.CompressHandler(r)),
		Addr:    "com-http.us:443",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	go log.Fatal(srv.ListenAndServeTLS("./comhttp.crt", "./comhttp.key"))
}
