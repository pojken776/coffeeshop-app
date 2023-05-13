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

	switch {
	case r.Method == http.MethodGet && listShopsRe.MatchString(r.URL.Path):
		fmt.Println("Run the list query")
		return
	case r.Method == http.MethodGet && getShopRe.MatchString(r.URL.Path):
		h.Get(w, r)
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

func (handler *ShopHandler) Get(w http.ResponseWriter, r *http.Request) {
	rows, err := handler.db.Query("SELECT id, name FROM coffeeshops WHERE id = $1", 1)
	if err != nil {
		log.Println("Error when attempting to list rows")
		log.Fatal(err)
	}
	defer rows.Close()
	shop := Shop{}
	for rows.Next() {
		err := rows.Scan(&shop.Id, &shop.Name)
		if err != nil {
			log.Println("Error when reading in line")
			log.Fatal(err)
		}
		log.Print(shop)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
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
