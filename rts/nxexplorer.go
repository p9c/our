package rts

import (
	"net/http"
	"net/http/httputil"

	"github.com/gorilla/mux"
)

// func respondWithJSON(w http.ResponseWriter, code int, block interface{}) {
// 	response, _ := json.Marshal(block)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(code)
// 	w.Write(response)
// }

func NXExplorerHandler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "/coin/" + mux.Vars(r)["coin"] + "/explorer"
		p.ServeHTTP(w, r)
	}
}

func NXBlockHandler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "/coin/" + mux.Vars(r)["coin"] + "/explorer/block/" + mux.Vars(r)["id"]
		p.ServeHTTP(w, r)
	}
}
func NXHashHandler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "/coin/" + mux.Vars(r)["coin"] + "/explorer/hash/" + mux.Vars(r)["id"]
		p.ServeHTTP(w, r)
	}
}
func NXTxHandler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "/coin/" + mux.Vars(r)["coin"] + "/explorer/tx/" + mux.Vars(r)["id"]
		p.ServeHTTP(w, r)
	}
}
func NXAddrsHandler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "/coin/" + mux.Vars(r)["coin"] + "/explorer/addr/" + mux.Vars(r)["id"]
		p.ServeHTTP(w, r)
	}
}
