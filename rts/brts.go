package rts

import (
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

func ViewBlock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	data := mod.BlVw{
		ID:       id,
		Block:    mod.Block{},
		AMPblock: amp.AMPB(),
	}
	renderTemplate(w, "block", "explorerbase", data)
}
func ViewTx(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	data := mod.TxVw{
		ID:       id,
		Tx:       mod.Tx{},
		AMPblock: amp.AMPB(),
	}
	renderTemplate(w, "tx", "explorerbase", data)
}
func ViewAddr(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	data := mod.AdVw{
		ID:       id,
		Addr:     mod.Addr{},
		AMPblock: amp.AMPB(),
	}
	renderTemplate(w, "addr", "explorerbase", data)
}

func ApiData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["subdomain"]
	id := vars["id"]
	tp := vars["type"]
	url := ComServer + "api/explorer/" + coin + "/" + tp + "/" + id
	data, _ := getData(url)
	w.Write([]byte(data))
}
func ApiLast(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["subdomain"]
	url := ComServer + "api/explorer/" + coin + "/last"
	data, _ := getData(url)
	fmt.Println("blkblkblkblkblkblk", data)
	w.Write([]byte(data))
}
func ApiInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["subdomain"]
	url := ComServer + "api/explorer/" + coin + "/info"
	data, _ := getData(url)
	//fmt.Println("blkblkblkblkblkblk", data)
	w.Write([]byte(data))
}
func ApiMiningInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["subdomain"]
	url := ComServer + "api/explorer/" + coin + "/gmi"
	data, _ := getData(url)
	//fmt.Println("blkblkblkblkblkblk", data)
	w.Write([]byte(data))
}
func ApiRawPool(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["subdomain"]
	url := ComServer + "api/explorer/" + coin + "/rmp"
	data, _ := getData(url)
	//fmt.Println("blkblkblkblkblkblk", data)
	w.Write([]byte(data))
}

func ExplorerHandler(w http.ResponseWriter, r *http.Request) {
	//tmpl := template.Must(template.ParseFiles(".html"))
	data := mod.Index{
		Blocks:   []mod.Block{},
		AMPblock: amp.AMPB(),
	}
	renderTemplate(w, "explorerindex", "explorerbase", data)
}

func DoSearch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["subdomain"]
	r.ParseForm()
	search := r.FormValue("src")
	fmt.Println("searchsearchsearchsearchsearchsearchsearchsearch", search)

	tps := []string{"block", "blockhash", "tx", "addr"}
	var tpt string
	for _, tp := range tps {
		url := ComServer + "api/explorer/" + coin + "/" + tp + "/" + search
		fmt.Println("urlurlurlurlurlurlurlurlurlurlurl", url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("___________________________________")
			print(err.Error())
			fmt.Println("___________________________________")
		} else {
			fmt.Println("___________________________________")
			print(string(resp.StatusCode) + resp.Status)
			fmt.Println("___________________________________")
		}

	}

	http.Redirect(w, r, fmt.Sprintf("/"+tpt+"/"+search), http.StatusPermanentRedirect)
}
