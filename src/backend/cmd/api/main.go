package main

import (
	"fmt"
	"net/http"
	"regexp"
)

var listShopsRe = regexp.MustCompile(`^\/shops[\/]*$`)
var getShopRe = regexp.MustCompile(`^\/shops\/[\d+]*$`)
var createShopRe = regexp.MustCompile(`^\/shops[\/]*$`)
var editShopRe = regexp.MustCompile(`^\/shops[\d+]*$`)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not found"))
}

type shopHandler struct{}

func (h *shopHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch {
	case r.Method == http.MethodGet && listShopsRe.MatchString(r.URL.Path):
		fmt.Println("Run the list query")
		return
	case r.Method == http.MethodGet && getShopRe.MatchString(r.URL.Path):
		fmt.Println("Run the get query")
		return
	case r.Method == http.MethodPost && createShopRe.MatchString(r.URL.Path):
		fmt.Println("Run the create command")
		return
	case r.Method == http.MethodPut && editShopRe.MatchString(r.URL.Path):
		fmt.Println("Run the edit query")
		return
	default:
		notFound(w, r)
		return
	}

}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/shop", &shopHandler{})

	http.ListenAndServe("localhost:8080", mux)

}
