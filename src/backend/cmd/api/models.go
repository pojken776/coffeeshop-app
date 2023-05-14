package main

import "time"

type HoursOfOperation struct {
	SundayOpen     time.Time
	SundayClose    time.Time
	MondayOpen     time.Time
	MondayClose    time.Time
	TuesdayOpen    time.Time
	TuesdayClose   time.Time
	WednesdayOpen  time.Time
	WednesdayClose time.Time
	ThursdayOpen   time.Time
	ThursdayClose  time.Time
	FridayOpen     time.Time
	FridayClose    time.Time
	SaturdayOpen   time.Time
	SaturdayClose  time.Time
}

type Amenities struct {
	Toilet int8
}

type Rating struct {
	Drinks       int8
	Food         int8
	Wifi         int8
	PowerOutlets int8
	Seating      int8
	Service      int8
}

type Address struct {
	Number string
	Street string
	City   string
	Zip    string
}

type Shop struct {
	Id               int
	Name             string
	Address          Address
	Rating           Rating
	Amenities        Amenities
	HoursOfOperation HoursOfOperation
}
