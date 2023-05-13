package main

import (
	"time"
)

type HoursOfOperation struct {
	Day   time.Weekday
	Open  time.Time
	Close time.Time
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
	HoursOfOperation []HoursOfOperation
}
