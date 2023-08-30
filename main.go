package main

import (
	"bus_listener/bus"
	"time"
)

func main() {
	alibaba := bus.NewAlibaba()
	helper := bus.NewHelper(alibaba)
	helper.BussesByReceivingTime(alibaba.Tehran(), alibaba.Shahroud(), time.Now().AddDate(0, 0, 2))
	//result, err := alibaba.Busses(alibaba.Tehran(), alibaba.Shahroud(), time.Now().AddDate(0, 0, 1))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for _, ticket := range result {
	//	if ticket.Seats > 0 {
	//		log.Println(ticket.Name, ticket.Date, ticket.Time)
	//	}
	//}
}
