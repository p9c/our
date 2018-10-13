package rts

import (
	"html/template"
	"net/http"

	"github.com/parallelcointeam/our/amp"
	"github.com/parallelcointeam/our/conf"
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
	templates["coins"] = template.Must(template.ParseFiles("tpl/coins.gohtml", "tpl/spectre.gohtml", "tpl/style.gohtml", "tpl/base.gohtml"))
	templates["home"] = template.Must(template.ParseFiles("tpl/home.gohtml", "tpl/spectre.gohtml", "tpl/style.gohtml", "tpl/base.gohtml"))
	templates["coin"] = template.Must(template.ParseFiles("tpl/coin.gohtml", "tpl/spectre.gohtml", "tpl/style.gohtml", "tpl/base.gohtml"))
	templates["explorerindex"] = template.Must(template.ParseFiles("tpl/explorerindex.gohtml", "tpl/spectre.gohtml", "tpl/style.gohtml", "tpl/explorerbase.gohtml"))
	templates["block"] = template.Must(template.ParseFiles("tpl/block.gohtml", "tpl/spectre.gohtml", "tpl/style.gohtml", "tpl/explorerbase.gohtml"))
	templates["tx"] = template.Must(template.ParseFiles("tpl/tx.gohtml", "tpl/spectre.gohtml", "tpl/style.gohtml", "tpl/explorerbase.gohtml"))
	templates["addr"] = template.Must(template.ParseFiles("tpl/addr.gohtml", "tpl/spectre.gohtml", "tpl/style.gohtml", "tpl/explorerbase.gohtml"))

	templates["cmc"] = template.Must(template.ParseFiles("tpl/frame/cmc.gohtml", "tpl/basehtml.gohtml"))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	coins := ser.GetCoins()
	data := mod.HCL{
		Coins: coins,
		AMP:   amp.AMPC(),
	}
	renderTemplate(w, "index", "base", data)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := "COINS"
	renderTemplate(w, "home", "emptybase", data)
}
