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

	templates["index"] = template.Must(template.ParseFiles("tpl/index.gohtml", "tpl/hlp/call.gohtml", "tpl/hlp/spectre.gohtml", "tpl/hlp/base.gohtml", "tpl/hlp/head.gohtml", "tpl/hlp/body.gohtml", "tpl/hlp/style.gohtml"))
	templates["coins"] = template.Must(template.ParseFiles("tpl/coins.gohtml", "tpl/hlp/spectre.gohtml", "tpl/hlp/base.gohtml", "tpl/hlp/head.gohtml", "tpl/hlp/body.gohtml", "tpl/hlp/style.gohtml"))
	templates["cmdns"] = template.Must(template.ParseFiles("tpl/cmdns.gohtml", "tpl/hlp/spectre.gohtml", "tpl/hlp/base.gohtml", "tpl/hlp/head.gohtml", "tpl/hlp/body.gohtml", "tpl/hlp/style.gohtml"))
	templates["home"] = template.Must(template.ParseFiles("tpl/home.gohtml", "tpl/hlp/spectre.gohtml", "tpl/hlp/base.gohtml", "tpl/hlp/head.gohtml", "tpl/hlp/body.gohtml", "tpl/hlp/style.gohtml"))
	templates["coin"] = template.Must(template.ParseFiles("tpl/coin.gohtml", "tpl/hlp/news.gohtml", "tpl/hlp/spectre.gohtml", "tpl/hlp/base.gohtml", "tpl/hlp/head.gohtml", "tpl/hlp/body.gohtml", "tpl/hlp/style.gohtml"))
	templates["explorerindex"] = template.Must(template.ParseFiles("tpl/explorerindex.gohtml", "tpl/hlp/spectre.gohtml", "tpl/hlp/base.gohtml", "tpl/hlp/head.gohtml", "tpl/hlp/body.gohtml", "tpl/hlp/style.gohtml"))
	templates["blockheight"] = template.Must(template.ParseFiles("tpl/block.gohtml", "tpl/block_height.gohtml", "tpl/hlp/spectre.gohtml", "tpl/hlp/base.gohtml", "tpl/hlp/head.gohtml", "tpl/hlp/body.gohtml", "tpl/hlp/style.gohtml"))
	templates["blockhash"] = template.Must(template.ParseFiles("tpl/block.gohtml", "tpl/block_hash.gohtml", "tpl/hlp/spectre.gohtml", "tpl/hlp/base.gohtml", "tpl/hlp/head.gohtml", "tpl/hlp/body.gohtml", "tpl/hlp/style.gohtml"))
	templates["tx"] = template.Must(template.ParseFiles("tpl/tx.gohtml", "tpl/hlp/spectre.gohtml", "tpl/hlp/base.gohtml", "tpl/hlp/head.gohtml", "tpl/hlp/body.gohtml", "tpl/hlp/style.gohtml"))
	templates["addr"] = template.Must(template.ParseFiles("tpl/addr.gohtml", "tpl/hlp/spectre.gohtml", "tpl/hlp/base.gohtml", "tpl/hlp/head.gohtml", "tpl/hlp/body.gohtml", "tpl/hlp/style.gohtml"))

	templates["404"] = template.Must(template.ParseFiles("tpl/404.gohtml", "tpl/hlp/search.gohtml", "tpl/hlp/spectre.gohtml", "tpl/hlp/base.gohtml", "tpl/hlp/head.gohtml", "tpl/hlp/body.gohtml", "tpl/hlp/style.gohtml"))

	templates["cmc"] = template.Must(template.ParseFiles("tpl/frame/cmc.gohtml", "tpl/hlp/spectre.gohtml", "tpl/hlp/base.gohtml", "tpl/hlp/head.gohtml", "tpl/hlp/body.gohtml", "tpl/hlp/style.gohtml"))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	coins := ser.GetCoins()
	data := mod.HCL{
		Coins: coins,
		AMP:   amp.AmP(),
	}
	renderTemplate(w, "index", "base", data)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := "COINS"
	renderTemplate(w, "home", "base", data)
}
