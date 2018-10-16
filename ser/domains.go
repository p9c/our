package ser

import "fmt"

func GetNames() {
	coins := GetCoins()

	//var cert models.OneCert
	//var ocert []string
	//mico := 0
	for ico, coin := range coins {
		url := coin.Coin.Slug + ".com-http.us"
		fmt.Println("Numero : ", ico)
		fmt.Println("Ristreto : ", url)

	}

}
