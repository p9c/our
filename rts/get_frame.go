package rts

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/parallelcointeam/our/jdb"
	"github.com/parallelcointeam/our/mod"
)

func CMCHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cn := vars["subdomain"]

	gdb, err := jdb.OpenDB()
	if err != nil {
	}
	coin := mod.Coin{}
	if err := gdb.Read("coins", cn, &coin); err != nil {
		fmt.Println("Error", err)
	}
	sym := coin.Symbol

	renderTemplate(w, "cmc", "basehtml", sym)
}
