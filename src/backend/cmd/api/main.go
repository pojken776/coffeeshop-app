package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/jackc/pgx"
)

var listShopsRe = regexp.MustCompile(`^\/shops[\/]*$`)
var getShopRe = regexp.MustCompile(`^\/shops\/[\d+]*$`)
var createShopRe = regexp.MustCompile(`^\/shops[\/]*$`)
var editShopRe = regexp.MustCompile(`^\/shops[\d+]*$`)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

var databaseHost = "127.0.0.1"
var databaseUsername = os.Getenv("POSTGRES_USERNAME")
var databasePassword = os.Getenv("POSTGRES_PASSWORD")
var databasePort uint16 = 5432

type databaseConnection struct {
	host     string
	port     uint16
	username string
	password string
}

func connectDB(connectionDetails databaseConnection) {
	pgxConfig := pgx.ConnConfig{Host: connectionDetails.host, Port: connectionDetails.port, Password: connectionDetails.password, User: connectionDetails.username}
	conn, err := pgx.Connect(pgxConfig)
	if err != nil {
		log.Printf("Failed to connect to database at %s:%d", connectionDetails.host, connectionDetails.port)
		log.Fatal(err)
	}

	defer conn.Close()
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
	log.Println("Connected to database!")
	connectDB(databaseConnection{host: databaseHost, port: databasePort, username: databaseUsername, password: databasePassword})

	mux := http.NewServeMux()
	mux.Handle("/shops", &shopHandler{})

	log.Println("Listening on :8080")
	http.ListenAndServe("localhost:8080", mux)

}
