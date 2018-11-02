package rts

import (
	"net/http"
	"net/http/httputil"

	"github.com/gorilla/mux"
)

func NXIndexHandler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "/"
		p.ServeHTTP(w, r)
	}
}

func NXCoinsHandler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "/coins/"
		p.ServeHTTP(w, r)
	}
}
func NXCoinHandler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "/coin/" + mux.Vars(r)["coin"]
		p.ServeHTTP(w, r)
	}
}
