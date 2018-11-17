package ser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/parallelcointeam/our/conf"
	"github.com/parallelcointeam/our/jdb"
	"github.com/parallelcointeam/our/mod"
)

var cf = conf.CsYsConf()
var ComServer = cf.ComServer

func init() {
//	getData()
}
func GetCoins() (coins []mod.VCoin) {
	gdb, err := jdb.OpenDB()
	if err != nil {
	}
	gcoins, err := gdb.ReadAll("coins")
	if err != nil {
		fmt.Println("Read Coins Error", err)
		getData()
	}
	for _, gc := range gcoins {
		coin := mod.VCoin{}
		if err := json.Unmarshal([]byte(gc), &coin); err != nil {
			fmt.Println("Read Coin range Error", err)
		}
		coins = append(coins, coin)
	}
	return coins
}

// func GetHomeCoins() (coins []mod.HC) {
// 	gdb, err := jdb.OpenDB()
// 	if err != nil {
// 	}
// 	gcoins, err := gdb.ReadAll("coins")
// 	if err != nil {
// 		fmt.Println("Read Coins Error", err)
// 		getData()
// 	}
// 	for _, gc := range gcoins {
// 		gcoin := mod.Coin{}
// 		coin := mod.HC{}
// 		if err := json.Unmarshal([]byte(gc), &gcoin); err != nil {
// 			fmt.Println("Read Coin range Error", err)
// 		}
// 		n := gcoin.Name
// 		s := gcoin.Symbol
// 		g := gcoin.Slug
// 		// a := gcoin.Algo.A
// 		// z := gcoin.Algo.Z
// 		i := gcoin.Imgs.Img32
// 		coin = mod.HC{n, s, g, i}
// 		coins = append(coins, coin)
// 	}
// 	return coins
// }

func getData() {
	gdb, err := jdb.OpenDB()
	if err != nil {
	}
	gamp, err := http.Get(ComServer + "a/c/a")
	if err != nil {
		fmt.Println("AMP gampgampgampgamp", gamp)
	}

	fmt.Println("Read error", err)
	defer gamp.Body.Close()
	mapCoins, err := ioutil.ReadAll(gamp.Body)
	var gcoins []mod.Coin
	// var coins []mod.CoinAmp
	// var bcoins []mod.BCoinAmp
	json.Unmarshal(mapCoins, &gcoins)
	if err != nil {
		fmt.Println("Read error", err)
	}
	//	algos := make(map[string]mod.Algo)
	for _, coin := range gcoins {
		gcoin := mod.Coin{}
		if err := gdb.Read("coins", coin.Slug, &gcoin); err != nil {
			fmt.Println("Error", err)
		}
		if coin.Slug != gcoin.Slug {
			// var acoin mod.CoinAmp
			// var bcoin mod.BCoinAmp
			gimg, err := http.Get(ComServer + "a/img/" + coin.Slug)
			if err != nil {
				fmt.Println("Img get error", err)
			}
			defer gimg.Body.Close()
			mapImgs, err := ioutil.ReadAll(gimg.Body)
			var imgs mod.Imgs
			json.Unmarshal(mapImgs, &imgs)
			coin = mod.Coin{
				Name:     coin.Name,
				Symbol:   coin.Symbol,
				Slug:     coin.Slug,
				Algo:     coin.Algo,
				BitNode: coin.BitNode,
				CData:    coin.CData,
				Imgs:     imgs,
			}
			// acoin = mod.CoinAmp{
			// 	Name:   coin.Name,
			// 	Symbol: coin.Symbol,
			// 	Slug:   coin.Slug,
			// 	Algo:   coin.Algo,
			// }
			cO := map[string]interface{}{"coin": coin}
			gdb.Write("coins", coin.Slug, cO)
			fmt.Println("Inserteded coin:", coin.Slug)
			// coins = append(coins, acoin)

			// if coin.Explorer != false {
			// 	bcoin = mod.BCoinAmp{
			// 		Name:   coin.Name,
			// 		Symbol: coin.Symbol,
			// 		Slug:   coin.Slug,
			// 		Algo:   coin.Algo,
			// 	}
			// 	bcoins = append(bcoins, bcoin)
			// }

		}
		//cEs := map[string]interface{}{"coins": bcoins}
		//cOs := map[string]interface{}{"coins": coins}
		//gdb.Write("index", "bitnodes", cEs)
		//gdb.Write("index", "coins", cOs)
	}
}

// func GetAlgos(server string) []mod.Algo {
// 	gdb, err := db.OpenDB()
// 	if err != nil {
// 	}
// 	records, err := gdb.ReadAll("algo")
// 	if err != nil {
// 		fmt.Println("ErrorR1", err)
// 		GetData(server)
// 	}
// 	algos := []mod.Algo{}
// 	for _, f := range records {
// 		algo := mod.Algo{}
// 		if err := json.Unmarshal([]byte(f), &algo); err != nil {
// 			fmt.Println("ErrorR2", err)
// 		}
// 		//algos = append(algos, algo)
// 	}
// 	return algos
// }

// func GetCoinsByAlgo(algo, server string) interface{} {
// 	gdb, err := db.OpenDB()
// 	if err != nil {
// 	}
// 	records, err := gdb.ReadAll("coin")
// 	if err != nil {
// 		fmt.Println("ErrorR1", err)
// 		GetData(server)
// 	}

// 	coins := []mod.Coin{}
// 	for _, f := range records {
// 		coin := mod.Coin{}
// 		if err := json.Unmarshal([]byte(f), &coin); err != nil {
// 			fmt.Println("ErrorR2", err)
// 		}
// 		calgo := coin.Algo.Z
// 		if algo == calgo {

// 			coins = append(coins, coin)
// 		}

// 	}

// 	return coins

// }
