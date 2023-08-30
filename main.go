package main

import (
	"bus_listener/bus"
	"bus_listener/locations"
	"log"
	"time"
)

func main() {
	date := time.Now()
	availables, err := bus.GetAvailables(locations.Tehran, locations.Shahroud, date)
	if err != nil {
		log.Fatal(err)
	}

	for _, available := range availables {
		seats, _ := available.Seats()
		log.Printf("Name: %s -> Number of seats: %d", available.Name, len(seats))
	}
}
