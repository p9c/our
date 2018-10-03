package rts

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/parallelcointeam/our/amp"
	"github.com/parallelcointeam/our/conf"
	"github.com/parallelcointeam/our/jdb"
	"github.com/parallelcointeam/our/mod"
	"github.com/parallelcointeam/our/ser"
)

var cf = conf.CsYsConf()
var ComServer = cf.ComServer
var templates = make(map[string]*template.Template)

var last string

func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	templates["index"] = template.Must(template.ParseFiles("tpl/index.gohtml", "tpl/spectre.gohtml", "tpl/style.gohtml", "tpl/base.gohtml"))
	templates["coins"] = template.Must(template.ParseFiles("tpl/coins.gohtml", "tpl/spectre.gohtml", "tpl/style.gohtml", "tpl/emptybase.gohtml"))
	templates["home"] = template.Must(template.ParseFiles("tpl/home.gohtml", "tpl/spectre.gohtml", "tpl/style.gohtml", "tpl/emptybase.gohtml"))
	templates["coin"] = template.Must(template.ParseFiles("tpl/coin.gohtml", "tpl/spectre.gohtml", "tpl/style.gohtml", "tpl/base.gohtml"))
	templates["explorerindex"] = template.Must(template.ParseFiles("tpl/explorerindex.gohtml", "tpl/spectre.gohtml", "tpl/style.gohtml", "tpl/explorerbase.gohtml"))
	templates["block"] = template.Must(template.ParseFiles("tpl/block.gohtml", "tpl/spectre.gohtml", "tpl/style.gohtml", "tpl/explorerbase.gohtml"))
	templates["tx"] = template.Must(template.ParseFiles("tpl/tx.gohtml", "tpl/spectre.gohtml", "tpl/style.gohtml", "tpl/explorerbase.gohtml"))
	templates["addr"] = template.Must(template.ParseFiles("tpl/addr.gohtml", "tpl/spectre.gohtml", "tpl/style.gohtml", "tpl/explorerbase.gohtml"))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	coins := ser.GetCoins()
	data := mod.HCL{
		Coins:  coins,
		AMPimg: amp.AMPI(),
	}
	renderTemplate(w, "index", "base", data)
}
func CoinsHandler(w http.ResponseWriter, r *http.Request) {
	coins := ser.GetCoins()
	data := mod.HCL{
		Coins:  coins,
		AMPimg: amp.AMPI(),
	}
	renderTemplate(w, "coins", "emptybase", data)
}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := "COINS"
	renderTemplate(w, "home", "emptybase", data)
}
func CoinHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["subdomain"]
	gdb, err := jdb.OpenDB()
	if err != nil {
	}
	vCoin := mod.VCoin{}
	if err := gdb.Read("coins", coin, &vCoin); err != nil {
		fmt.Println("Error", err)
	}
	gCoin := vCoin.Coin
	data := mod.CoinVw{
		Coin:   gCoin,
		AMPimg: amp.AMPI(),
	}
	renderTemplate(w, "coin", "base", data)
}

func CoinNewsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["subdomain"]
	url := "https://query.yahooapis.com/v1/public/yql?q=select%20title%2Clink%2C%20pubDate%2Cdescription%20from%20rss%20where%20url%20%3D%20'https%3A%2F%2Fnews.google.com%2Fnews%2Frss%2Fsearch%2Fsection%2Fq%2F" + coin + "%3Fned%3Dus%26gl%3DUS%26hl%3Den'&format=json&env=store%3A%2F%2Fdatatables.org%2Falltableswithkeys"
	data, _ := getData(url)
	w.Write([]byte(data))
}

func getData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
	}
	defer resp.Body.Close()
	mapBody, err := ioutil.ReadAll(resp.Body)
	return mapBody, err
}

func renderTemplate(w http.ResponseWriter, name string, template string, viewModel interface{}) {
	tmpl, ok := templates[name]
	if !ok {
		http.Error(w, "The template does not exist.", http.StatusInternalServerError)
	}
	err := tmpl.ExecuteTemplate(w, template, viewModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
