package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (handler *ShopHandler) List(w http.ResponseWriter, r *http.Request) {
	shops := []Shop{}

	shopRows, err := ListShops(handler.db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, shopRow := range *shopRows {
		shops = append(shops, Shop{Id: shopRow.Id, Name: shopRow.Name, Address: Address{
			Number: shopRow.Number,
			Street: shopRow.Street,
			City:   shopRow.City,
			Zip:    shopRow.Zip,
		}, Rating: Rating{
			Drinks:       int8(shopRow.Drinks.Int16),
			Food:         int8(shopRow.Food.Int16),
			Wifi:         int8(shopRow.Wifi.Int16),
			PowerOutlets: int8(shopRow.PowerOutlets.Int16),
			Seating:      int8(shopRow.Seating.Int16),
			Service:      int8(shopRow.Service.Int16),
		}, HoursOfOperation: HoursOfOperation{
			SundayOpen:     shopRow.SundayOpen.Time,
			SundayClose:    shopRow.SundayClose.Time,
			MondayOpen:     shopRow.MondayOpen.Time,
			MondayClose:    shopRow.MondayClose.Time,
			TuesdayOpen:    shopRow.TuesdayOpen.Time,
			TuesdayClose:   shopRow.TuesdayClose.Time,
			WednesdayOpen:  shopRow.WednesdayOpen.Time,
			WednesdayClose: shopRow.WednesdayClose.Time,
			ThursdayOpen:   shopRow.ThursdayOpen.Time,
			ThursdayClose:  shopRow.ThursdayClose.Time,
			FridayOpen:     shopRow.FridayOpen.Time,
			FridayClose:    shopRow.FridayClose.Time,
			SaturdayOpen:   shopRow.SaturdayOpen.Time,
			SaturdayClose:  shopRow.SaturdayClose.Time,
		}},
		)
	}
	responseBody, err := json.Marshal(shops)
	if err != nil {
		log.Printf("Error marshaling response body: %v", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}

func (handler *ShopHandler) Get(w http.ResponseWriter, r *http.Request) {
	targetId, err := strconv.Atoi(strings.Split(r.URL.Path, "/")[2])
	if err != nil {
		log.Printf("Unable to convert id to int from path: %s", r.URL.Path)
		return
	}

	shopRow, err := GetShop(targetId, handler.db)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No shops with id %d", targetId)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Shop not found"))
		return
	case err != nil:
		log.Printf("Query error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	default:

		shop := Shop{Id: shopRow.Id, Name: shopRow.Name, Address: Address{
			Number: shopRow.Number,
			Street: shopRow.Street,
			City:   shopRow.City,
			Zip:    shopRow.Zip,
		}, Rating: Rating{
			Drinks:       int8(shopRow.Drinks.Int16),
			Food:         int8(shopRow.Food.Int16),
			Wifi:         int8(shopRow.Wifi.Int16),
			PowerOutlets: int8(shopRow.PowerOutlets.Int16),
			Seating:      int8(shopRow.Seating.Int16),
			Service:      int8(shopRow.Service.Int16),
		}, HoursOfOperation: HoursOfOperation{
			SundayOpen:     shopRow.SundayOpen.Time,
			SundayClose:    shopRow.SundayClose.Time,
			MondayOpen:     shopRow.MondayOpen.Time,
			MondayClose:    shopRow.MondayClose.Time,
			TuesdayOpen:    shopRow.TuesdayOpen.Time,
			TuesdayClose:   shopRow.TuesdayClose.Time,
			WednesdayOpen:  shopRow.WednesdayOpen.Time,
			WednesdayClose: shopRow.WednesdayClose.Time,
			ThursdayOpen:   shopRow.ThursdayOpen.Time,
			ThursdayClose:  shopRow.ThursdayClose.Time,
			FridayOpen:     shopRow.FridayOpen.Time,
			FridayClose:    shopRow.FridayClose.Time,
			SaturdayOpen:   shopRow.SaturdayOpen.Time,
			SaturdayClose:  shopRow.SaturdayClose.Time,
		}}

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

func (handler *ShopHandler) Create(w http.ResponseWriter, r *http.Request) {
	shop := Shop{}
	if err := json.NewDecoder(r.Body).Decode(&shop); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	statement, err := handler.db.Prepare("INSERT INTO coffeeshops(id, name, address_number, address_street, address_city, address_zip) VALUES(DEFAULT, $1, $2, $3, $4, $5);")
	if err != nil {
		log.Printf("Error preparing insert shop statement: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer statement.Close()

	_, err = statement.Exec(shop.Name, shop.Address.Number, shop.Address.Street, shop.Address.City, shop.Address.Zip)
	if err != nil {
		log.Printf("Error inserting shop: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Inserted shop"))

}
