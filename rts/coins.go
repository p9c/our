package rts

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/parallelcointeam/our/amp"
	"github.com/parallelcointeam/our/jdb"
	"github.com/parallelcointeam/our/mod"
	"github.com/parallelcointeam/our/ser"
)

func CoinsHandler(w http.ResponseWriter, r *http.Request) {
	coins := ser.GetCoins()
	data := mod.HCL{
		Coins: coins,
		AMP:   amp.AMPC(),
	}
	renderTemplate(w, "coins", "base", data)
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
		Coin: gCoin,
		AMP:  amp.AMPC(),
	}
	renderTemplate(w, "coin", "base", data)
}
