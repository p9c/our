package rts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/parallelcointeam/our/amp"
	"github.com/parallelcointeam/our/mod"

	"github.com/gorilla/mux"
)

// func respondWithJSON(w http.ResponseWriter, code int, block interface{}) {
// 	response, _ := json.Marshal(block)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(code)
// 	w.Write(response)
// }

func ExplorerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["subdomain"]
	gCoin := getCoin(coin)
	data := mod.Explorer{
		Coin: gCoin,
		AMP:  amp.AMPB(),
	}
	renderTemplate(w, "explorerindex", "explorerbase", data)
}

func ViewBlockHeight(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["subdomain"]
	id := vars["id"]
	gCoin := getCoin(coin)
	data := mod.BlVw{
		ID:    id,
		Coin:  gCoin,
		Block: mod.Block{},
		AMP:   amp.AMPB(),
	}
	//fmt.Println("datadatadatadatadata", data)

	renderTemplate(w, "blockheight", "explorerbase", data)
}
func ViewBlockHash(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["subdomain"]
	id := vars["id"]
	gCoin := getCoin(coin)
	data := mod.BlVw{
		ID:    id,
		Coin:  gCoin,
		Block: mod.Block{},
		AMP:   amp.AMPB(),
	}
	//fmt.Println("datadatadatadatadata", data)

	renderTemplate(w, "blockhash", "explorerbase", data)
}
func ViewTx(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["subdomain"]
	id := vars["id"]
	gCoin := getCoin(coin)
	data := mod.TxVw{
		ID:   id,
		Coin: gCoin,
		Tx:   mod.Tx{},
		AMP:  amp.AMPB(),
	}
	renderTemplate(w, "tx", "explorerbase", data)
}
func ViewAddr(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["subdomain"]
	id := vars["id"]
	gCoin := getCoin(coin)
	data := mod.AdVw{
		ID:   id,
		Coin: gCoin,
		Addr: mod.Addr{},
		AMP:  amp.AMPB(),
	}
	renderTemplate(w, "addr", "explorerbase", data)
}

func ApiData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["subdomain"]
	id := vars["id"]
	tp := vars["type"]
	url := ComServer + "a/e/" + coin + "/" + tp + "/" + id
	data, _ := getData(url)
	w.Write([]byte(data))
}
func ApiLast(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["subdomain"]
	url := ComServer + "a/e/" + coin + "/last"
	data, _ := getData(url)
	//fmt.Println("blkblkblkblkblkblk", data)
	w.Write([]byte(data))
}
func ApiInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["subdomain"]
	url := ComServer + "a/e/" + coin + "/info"
	data, _ := getData(url)
	//fmt.Println("blkblkblkblkblkblk", data)
	w.Write([]byte(data))
}
func ApiMiningInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["subdomain"]
	url := ComServer + "a/e/" + coin + "/gmi"
	data, _ := getData(url)
	//fmt.Println("blkblkblkblkblkblk", data)
	w.Write([]byte(data))
}
func ApiRawPool(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["subdomain"]
	url := ComServer + "a/e/" + coin + "/rmp"
	data, _ := getData(url)
	//fmt.Println("blkblkblkblkblkblk", data)
	w.Write([]byte(data))
}

func DoSearch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["subdomain"]
	r.ParseForm()
	search := r.FormValue("src")
	fmt.Println("searchsearchsearchsearchsearchsearchsearchsearch", search)

	tps := []string{"block", "blockhash", "tx", "addr"}
	var tpt = "noresults"
	for _, tp := range tps {
		url := ComServer + "a/e/" + coin + "/" + tp + "/" + search
		fmt.Println("urlurlurlurlurlurlurlurlurlurlurl", url)
		resp, err := getData(url)
		var search map[string]interface{}
		json.Unmarshal(resp, &search)
		if err != nil {
			fmt.Println("Read error", err)
		}
		if search["d"] != nil {
			tpt = tp

		}

	}

	http.Redirect(w, r, fmt.Sprintf("/"+tpt+"/"+search), http.StatusPermanentRedirect)
}
