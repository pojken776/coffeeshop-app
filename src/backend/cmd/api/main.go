package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

var listShopsRe = regexp.MustCompile(`^\/shops[\/]*$`)
var getShopRe = regexp.MustCompile(`^\/shops\/[\d+]*$`)
var createShopRe = regexp.MustCompile(`^\/shops[\/]*$`)
var editShopRe = regexp.MustCompile(`^\/shops[\d+]*$`)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

type ShopHandler struct {
	db *sql.DB
}

func (h *ShopHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet && listShopsRe.MatchString(r.URL.Path):
		h.List(w, r)
		return
	case r.Method == http.MethodGet && getShopRe.MatchString(r.URL.Path):
		h.Get(w, r)
		return
	case r.Method == http.MethodPost && createShopRe.MatchString(r.URL.Path):
		h.Create(w, r)
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
	shopHandler := ShopHandler{}
	ConnectDB(&shopHandler)
	log.Println("Connected to database!")

	mux := http.NewServeMux()
	mux.Handle("/shops", &shopHandler)
	mux.Handle("/shops/", &shopHandler)

	log.Println("Listening on :8080")
	http.ListenAndServe("localhost:8080", mux)

	defer shopHandler.db.Close()

}
