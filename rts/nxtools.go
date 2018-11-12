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

func NXFrame(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "/coin/" + mux.Vars(r)["coin"] + "/frame/" + mux.Vars(r)["frame"]
		p.ServeHTTP(w, r)
	}
}
func NXNuxt(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		nfile := vars["nfile"]
		r.URL.Path = "/_nuxt/" + nfile
		p.ServeHTTP(w, r)
	}
}
