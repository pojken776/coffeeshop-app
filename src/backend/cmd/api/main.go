package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
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

func (handler *ShopHandler) Create(w http.ResponseWriter, r *http.Request) {
	shop := Shop{}
	if err := json.NewDecoder(r.Body).Decode(&shop); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	statement, err := handler.db.Prepare("INSERT INTO coffeeshops(id, name, address_number, address_street, address_city, address_zip) VALUES(DEFAULT, $1, $2, $3, $4, $5);")
	if err != nil {
		log.Printf("Error prepapring insert shop statement: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer statement.Close()

	result, err := statement.Exec(shop.Name, shop.Address.Number, shop.Address.Street, shop.Address.City, shop.Address.Zip)
	if err != nil {
		log.Printf("Error inserting shop: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println(result)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Inserted shop"))

}

func (handler *ShopHandler) Get(w http.ResponseWriter, r *http.Request) {
	targetId, err := strconv.Atoi(strings.Split(r.URL.Path, "/")[2])
	if err != nil {
		log.Printf("Unable to convert id to int from path: %s", r.URL.Path)
		return
	}

	shop := Shop{}
	err = handler.db.QueryRow("SELECT id, name FROM coffeeshops WHERE id = $1", targetId).Scan(&shop.Id, &shop.Name)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No shops with id %d", targetId)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Shop not found"))
		return
	case err != nil:
		log.Printf("Query error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	default:
		log.Printf("Found shop with id: %d and name: %s", shop.Id, shop.Name)

		responseBody, err := json.Marshal(shop)
		if err != nil {
			log.Printf("Unable to marshal shop response")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(responseBody)
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
