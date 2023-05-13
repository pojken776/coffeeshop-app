package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var databaseHost = os.Getenv("DATABASE_HOST")
var databaseUsername = os.Getenv("DATABASE_USERNAME")
var databasePassword = os.Getenv("DATABASE_PASSWORD")
var DatabaseName = os.Getenv("DATABASE_NAME")
var databasePort = os.Getenv("DATABASE_PORT")

func ConnectDB(handler *ShopHandler) {
	dbPort, err := strconv.Atoi(databasePort)
	if err != nil {
		log.Fatal("Unable to convert database port to string")
	}

	address := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", databaseUsername, databasePassword, databaseHost, dbPort, DatabaseName)
	handler.db, err = sql.Open("pgx", address)
	if err != nil {
		log.Printf("Failed to connect to database with the following connection string: %s", address)
		log.Fatal(err)
	}
	if err := handler.db.PingContext(context.Background()); err != nil {
		log.Fatal(err)
	}

}
