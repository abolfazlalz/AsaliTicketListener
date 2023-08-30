package bus

import (
	"time"
)

type ServiceInterface interface {
	Busses(startCity string, destinationCity string, time time.Time) ([]Ticket, error)
	BussesWithChannel(startCity string, destinationCity string, time time.Time, ch chan []Ticket)
	TicketSeats(Ticket) (TicketSeats, error)
	Shahroud() string
	Tehran() string
}

type Helper struct {
	service ServiceInterface
}

func NewHelper(service ServiceInterface) *Helper {
	return &Helper{service: service}
}

func (s Helper) BussesByReceivingTime(start, dest string, date time.Time) {
	yesterday := date.AddDate(0, 0, -1)

	ch := make(chan []Ticket)
	go s.service.BussesWithChannel(start, dest, date, ch)
	go s.service.BussesWithChannel(start, dest, yesterday, ch)

	val1, val2 := <-ch, <-ch

}
