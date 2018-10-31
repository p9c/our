package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gorilla/mux"
)

//Target url: https://httpbin.org/headers
//Url through proxy:  http://localhost:3002/forward/headers

func main() {
	target := "http://127.0.0.1:7227/"
	remote, err := url.Parse(target)
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	r := mux.NewRouter()
	r.HandleFunc("/{coin}", handler(proxy))
	http.Handle("/", r)
	http.ListenAndServe(":3002", r)
}

func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = mux.Vars(r)["coin"]
		p.ServeHTTP(w, r)
	}
}
