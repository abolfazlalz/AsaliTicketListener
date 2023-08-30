package bus

import "time"

type Ticket struct {
	ID       string
	Name     string
	Date     time.Time
	Time     string
	Terminal string
	Seats    int
}

type TicketSeat struct {
	Index  int    `json:"index"`
	Number int    `json:"number"`
	Column int    `json:"column"`
	Row    int    `json:"row"`
	Status string `json:"status"`
}

type TicketSeats []TicketSeat
