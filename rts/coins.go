package rts

import (
	"encoding/base64"
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

func ImgHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["coin"]
	size := vars["size"]
	gdb, err := jdb.OpenDB()
	if err != nil {
	}
	vCoin := mod.VCoin{}
	if err := gdb.Read("coins", coin, &vCoin); err != nil {
		fmt.Println("Error", err)
	}
	var img string
	switch size {
	case "16":
		img = vCoin.Coin.Imgs.Img16
	case "32":
		img = vCoin.Coin.Imgs.Img32
	case "64":
		img = vCoin.Coin.Imgs.Img64
	case "128":
		img = vCoin.Coin.Imgs.Img128
	case "256":
		img = vCoin.Coin.Imgs.Img256
	}
	encoded, _ := base64.StdEncoding.DecodeString(img)
	w.Write(encoded)
}

func CertHandler(w http.ResponseWriter, r *http.Request) {
	ser.GetNames()

}
