package main

import (
	"comhttpus/rts"
	"fmt"

	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/robfig/cron"
)

const (
	ComServer = "http://127.0.0.1:9998"
)

func main() {

	cr := cron.New()
	cr.AddFunc("@every 17655s", func() {
		fmt.Println("Radi kron")
		//ser.GetData(ComServer)
	})
	cr.Start()
	jsonHandler := http.FileServer(http.Dir("./JDB/"))
	jsonHandler = http.StripPrefix("/json/", jsonHandler)
	r := mux.NewRouter()
	r.PathPrefix("/json/").Handler(jsonHandler)
	r.Host("com-http.us").Path("/").HandlerFunc(rts.IndexHandler).Name("index")
	r.Host("{subdomain}.com-http.us").Path("/").HandlerFunc(rts.CoinHandler).Name("coin")
	r.Host("{subdomain}.com-http.us").Path("/api/news").HandlerFunc(rts.CoinNewsHandler).Name("coin")
	log.Fatal(http.ListenAndServe(":8985", handlers.CORS()(r)))
}
