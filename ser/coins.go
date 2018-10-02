package ser

import (
	"comhttpus/jdb"
	"comhttpus/mod"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetCoins(server string) (coins []mod.Coin) {
	gdb, err := jdb.OpenDB()
	if err != nil {
	}
	gcoins, err := gdb.ReadAll("coins")
	if err != nil {
		fmt.Println("Read Coins Error", err)
		getData(server)
	}
	for _, gc := range gcoins {
		coin := mod.Coin{}
		if err := json.Unmarshal([]byte(gc), &coin); err != nil {
			fmt.Println("Read Coin range Error", err)
		}
		coins = append(coins, coin)
	}
	return coins
}

func GetHomeCoins(server, static string) (coins []mod.HC) {
	gdb, err := jdb.OpenDB()
	if err != nil {
	}
	gcoins, err := gdb.ReadAll("coins")
	if err != nil {
		fmt.Println("Read Coins Error", err)
		getData(server)
	}
	for _, gc := range gcoins {
		gcoin := mod.Coin{}
		coin := mod.HC{}
		if err := json.Unmarshal([]byte(gc), &gcoin); err != nil {
			fmt.Println("Read Coin range Error", err)
		}
		n := gcoin.Name
		s := gcoin.Symbol
		g := gcoin.Slug
		// a := gcoin.Algo.A
		// z := gcoin.Algo.Z
		i := static + "/" + g + "/" + g + "32.png"
		coin = mod.HC{n, s, g, i}
		coins = append(coins, coin)
	}
	return coins
}

func getData(server string) {
	gdb, err := jdb.OpenDB()
	if err != nil {
	}
	gamp, err := http.Get(server + "/api/amp")
	if err != nil {
		fmt.Println("AMP error", err)
	}
	defer gamp.Body.Close()
	mapCoins, err := ioutil.ReadAll(gamp.Body)
	var gcoins []mod.Coin
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
			gimg, err := http.Get(server + "/api/img/" + coin.Slug)
			if err != nil {
				fmt.Println("Img get error", err)
			}
			defer gimg.Body.Close()
			mapImgs, err := ioutil.ReadAll(gimg.Body)
			var imgs mod.Imgs
			json.Unmarshal(mapImgs, &imgs)
			coin = mod.Coin{
				Name:   coin.Name,
				Symbol: coin.Symbol,
				Slug:   coin.Slug,
				Algo:   coin.Algo,
				CData:  coin.CData,
				Imgs:   imgs,
			}
			cO := map[string]interface{}{"coin": coin}
			gdb.Write("coins", coin.Slug, cO)
			fmt.Println("Inserteded coin:", coin.Slug)
		}
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
