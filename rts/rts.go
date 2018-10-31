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

	//templates["base"] = template.Must(template.ParseFiles("tpl/hlp/plgs.gohtml","tpl/css/boot.gohtml", "tpl/hlp/base.gohtml", "tpl/hlp/head.gohtml", "tpl/hlp/body.gohtml", "tpl/hlp/style.gohtml", "tpl/hlp/spectre.gohtml"))

	templates["index"] = template.Must(template.ParseFiles("tpl/hlp/plgs.gohtml", "tpl/css/boot.gohtml", "tpl/css/grid.gohtml", "tpl/css/typo.gohtml", "tpl/css/btn.gohtml", "tpl/hlp/base.gohtml", "tpl/hlp/body.gohtml", "tpl/hlp/head.gohtml", "tpl/hlp/style.gohtml", "tpl/hlp/spectre.gohtml", "tpl/css/homecss.gohtml", "tpl/index.gohtml", "tpl/hlp/call.gohtml", "tpl/pnls/coinpnls.gohtml", "tpl/pnls/fb.gohtml"))

	templates["coins"] = template.Must(template.ParseFiles("tpl/hlp/plgs.gohtml", "tpl/css/boot.gohtml", "tpl/css/grid.gohtml", "tpl/css/typo.gohtml", "tpl/css/btn.gohtml", "tpl/hlp/base.gohtml", "tpl/hlp/body.gohtml", "tpl/hlp/head.gohtml", "tpl/hlp/style.gohtml", "tpl/hlp/spectre.gohtml", "tpl/css/homecss.gohtml", "tpl/coins.gohtml"))

	templates["cmdns"] = template.Must(template.ParseFiles("tpl/cmdns.gohtml"))
	templates["home"] = template.Must(template.ParseFiles("tpl/hlp/plgs.gohtml", "tpl/css/boot.gohtml", "tpl/css/grid.gohtml", "tpl/css/typo.gohtml", "tpl/css/btn.gohtml", "tpl/hlp/base.gohtml", "tpl/hlp/body.gohtml", "tpl/hlp/head.gohtml", "tpl/hlp/style.gohtml", "tpl/hlp/spectre.gohtml", "tpl/home.gohtml"))

	templates["coin"] = template.Must(template.ParseFiles("tpl/hlp/plgs.gohtml", "tpl/css/boot.gohtml", "tpl/css/grid.gohtml", "tpl/css/typo.gohtml", "tpl/css/btn.gohtml", "tpl/hlp/base.gohtml", "tpl/hlp/body.gohtml", "tpl/hlp/head.gohtml", "tpl/hlp/style.gohtml", "tpl/hlp/spectre.gohtml", "tpl/css/coincss.gohtml", "tpl/coin.gohtml", "tpl/pnls/coinpnls.gohtml", "tpl/pnls/fb.gohtml"))

	templates["explorer"] = template.Must(template.ParseFiles("tpl/hlp/plgs.gohtml", "tpl/css/boot.gohtml", "tpl/css/grid.gohtml", "tpl/css/typo.gohtml", "tpl/css/btn.gohtml", "tpl/hlp/base.gohtml", "tpl/hlp/body.gohtml", "tpl/hlp/head.gohtml", "tpl/hlp/style.gohtml", "tpl/hlp/spectre.gohtml", "tpl/pnls/explorerpnls.gohtml", "tpl/pnls/coinpnls.gohtml", "tpl/css/explorercss.gohtml", "tpl/explorer.gohtml"))

	templates["blockheight"] = template.Must(template.ParseFiles("tpl/hlp/plgs.gohtml", "tpl/css/boot.gohtml", "tpl/css/grid.gohtml", "tpl/css/typo.gohtml", "tpl/css/btn.gohtml", "tpl/hlp/base.gohtml", "tpl/hlp/body.gohtml", "tpl/hlp/head.gohtml", "tpl/hlp/style.gohtml", "tpl/hlp/spectre.gohtml", "tpl/pnls/explorerpnls.gohtml", "tpl/pnls/coinpnls.gohtml", "tpl/css/explorercss.gohtml", "tpl/block.gohtml", "tpl/block_height.gohtml"))
	templates["blockhash"] = template.Must(template.ParseFiles("tpl/hlp/plgs.gohtml", "tpl/css/boot.gohtml", "tpl/css/grid.gohtml", "tpl/css/typo.gohtml", "tpl/css/btn.gohtml", "tpl/hlp/base.gohtml", "tpl/hlp/body.gohtml", "tpl/hlp/head.gohtml", "tpl/hlp/style.gohtml", "tpl/hlp/spectre.gohtml", "tpl/pnls/explorerpnls.gohtml", "tpl/pnls/coinpnls.gohtml", "tpl/css/explorercss.gohtml", "tpl/block.gohtml", "tpl/block_hash.gohtml"))
	templates["tx"] = template.Must(template.ParseFiles("tpl/hlp/plgs.gohtml", "tpl/css/boot.gohtml", "tpl/css/grid.gohtml", "tpl/css/typo.gohtml", "tpl/css/btn.gohtml", "tpl/hlp/base.gohtml", "tpl/hlp/body.gohtml", "tpl/hlp/head.gohtml", "tpl/hlp/style.gohtml", "tpl/hlp/spectre.gohtml", "tpl/pnls/explorerpnls.gohtml", "tpl/pnls/coinpnls.gohtml", "tpl/css/explorercss.gohtml", "tpl/tx.gohtml"))
	templates["addr"] = template.Must(template.ParseFiles("tpl/hlp/plgs.gohtml", "tpl/css/boot.gohtml", "tpl/css/grid.gohtml", "tpl/css/typo.gohtml", "tpl/css/btn.gohtml", "tpl/hlp/base.gohtml", "tpl/hlp/body.gohtml", "tpl/hlp/head.gohtml", "tpl/hlp/style.gohtml", "tpl/hlp/spectre.gohtml", "tpl/pnls/explorerpnls.gohtml", "tpl/pnls/coinpnls.gohtml", "tpl/css/explorercss.gohtml", "tpl/addr.gohtml"))

	templates["404"] = template.Must(template.ParseFiles("tpl/hlp/plgs.gohtml", "tpl/css/boot.gohtml", "tpl/css/grid.gohtml", "tpl/css/typo.gohtml", "tpl/css/btn.gohtml", "tpl/hlp/base.gohtml", "tpl/hlp/body.gohtml", "tpl/hlp/head.gohtml", "tpl/hlp/style.gohtml", "tpl/hlp/spectre.gohtml", "tpl/404.gohtml", "tpl/hlp/search.gohtml"))

	templates["cmc"] = template.Must(template.ParseFiles("tpl/frame/cmc.gohtml"))
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
