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

type ShopRow struct {
	Id             int           // db field: id
	Name           string        // db field: name
	Number         string        // db field: address_number
	Street         string        // db field: address_street
	City           string        // db field: address_city
	Zip            string        // db field: address_zip
	Drinks         sql.NullInt16 // db field: rating_drinks
	Food           sql.NullInt16 // db field: rating_food
	Wifi           sql.NullInt16 // db field: rating_wifi
	PowerOutlets   sql.NullInt16 // db field: rating_poweroutlets
	Seating        sql.NullInt16 // db field: rating_seating
	Service        sql.NullInt16 // db field: rating_service
	Toilet         sql.NullInt16 // db field: amenities_toilet
	SundayOpen     sql.NullTime  // db field: hours_of_operation_sunday_open
	SundayClose    sql.NullTime  // db field: hours_of_operation_sunday_close
	MondayOpen     sql.NullTime  // db field: hours_of_operation_monday_open
	MondayClose    sql.NullTime  // db field: hours_of_operation_monday_close
	TuesdayOpen    sql.NullTime  // db field: hours_of_operation_tuesday_open
	TuesdayClose   sql.NullTime  // db field: hours_of_operation_tuesday_close
	WednesdayOpen  sql.NullTime  // db field: hours_of_operation_wednesday_open
	WednesdayClose sql.NullTime  // db field: hours_of_operation_wednesday_close
	ThursdayOpen   sql.NullTime  // db field: hours_of_operation_thursday_open
	ThursdayClose  sql.NullTime  // db field: hours_of_operation_thursday_close
	FridayOpen     sql.NullTime  // db field: hours_of_operation_friday_open
	FridayClose    sql.NullTime  // db field: hours_of_operation_friday_close
	SaturdayOpen   sql.NullTime  // db field: hours_of_operation_saturday_open
	SaturdayClose  sql.NullTime  // db field: hours_of_operation_saturday_close
}

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

func ListShops(db *sql.DB) (*[]ShopRow, error) {

	shopRows := []ShopRow{}

	rows, err := db.Query("SELECT * FROM coffeeshops")
	if err != nil {
		log.Println("Error querying all rows")
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		shopRow := ShopRow{}

		if err := rows.Scan(&shopRow.Id, &shopRow.Name, &shopRow.Number, &shopRow.Street, &shopRow.City, &shopRow.Zip, &shopRow.Drinks, &shopRow.Food, &shopRow.Wifi, &shopRow.PowerOutlets, &shopRow.Seating, &shopRow.Service, &shopRow.Toilet, &shopRow.SundayOpen, &shopRow.SundayClose, &shopRow.MondayOpen, &shopRow.MondayClose, &shopRow.TuesdayOpen, &shopRow.TuesdayClose, &shopRow.WednesdayOpen, &shopRow.WednesdayClose, &shopRow.ThursdayOpen, &shopRow.ThursdayClose, &shopRow.FridayOpen, &shopRow.FridayClose, &shopRow.SaturdayOpen, &shopRow.SaturdayClose); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}

		shopRows = append(shopRows, shopRow)
	}

	return &shopRows, err
}

func GetShop(id int, db *sql.DB) (*ShopRow, error) {
	shopRow := ShopRow{}
	err := db.QueryRow("SELECT * FROM coffeeshops WHERE id = $1;", id).Scan(&shopRow.Id, &shopRow.Name, &shopRow.Number, &shopRow.Street, &shopRow.City, &shopRow.Zip, &shopRow.Drinks, &shopRow.Food, &shopRow.Wifi, &shopRow.PowerOutlets, &shopRow.Seating, &shopRow.Service, &shopRow.Toilet, &shopRow.SundayOpen, &shopRow.SundayClose, &shopRow.MondayOpen, &shopRow.MondayClose, &shopRow.TuesdayOpen, &shopRow.TuesdayClose, &shopRow.WednesdayOpen, &shopRow.WednesdayClose, &shopRow.ThursdayOpen, &shopRow.ThursdayClose, &shopRow.FridayOpen, &shopRow.FridayClose, &shopRow.SaturdayOpen, &shopRow.SaturdayClose)

	return &shopRow, err
}
