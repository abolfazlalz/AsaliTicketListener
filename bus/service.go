package bus

import (
	"bus_listener/env"
	"bus_listener/notification"
	"context"
	"fmt"
	"log/slog"
	"time"
)

type ServiceInterface interface {
	Busses(startCity string, destinationCity string, time time.Time) ([]Ticket, error)
	BussesWithChannel(startCity string, destinationCity string, time time.Time, ch chan []Ticket)
	TicketSeats(Ticket) (TicketSeats, error)
	GetCity(cityId string) string
	Shahroud() string
	Tehran() string
}

type Helper struct {
	service      ServiceInterface
	notification notification.Interface
}

func NewHelper(service ServiceInterface, notification notification.Interface) *Helper {
	return &Helper{
		service:      service,
		notification: notification,
	}
}

func (s Helper) BussesByArrivedTime(start, dest string, date time.Time) []Ticket {
	yesterday := date.AddDate(0, 0, -1)

	ch := make(chan []Ticket)
	go s.service.BussesWithChannel(start, dest, date, ch)
	go s.service.BussesWithChannel(start, dest, yesterday, ch)

	busList := append(<-ch, <-ch...)

	result := make([]Ticket, 0)

	for _, bus := range busList {
		hour := bus.Date.Hour()
		if (hour < 3 || hour > 19) && bus.Seats > 0 {
			result = append(result, bus)
		}
	}

	return result
}

func (s Helper) CheckInterval(start string, end string, quit chan struct{}) {
	ctx := context.Background()
	duration, duErr := time.ParseDuration(env.Service.Interval)

	// If program has error make duration to 10
	if duErr != nil {
		duration = time.Second * 10
	}

	ticker := time.NewTicker(duration)

	go func() {
		for {
			select {
			case <-ticker.C:
				result := s.BussesByArrivedTime(start, end, time.Now().AddDate(0, 0, 1))
				if len(result) == 0 {
					msg := fmt.Sprintf("No ticket to %s for tomorrow !!", end)
					slog.Log(ctx, slog.LevelInfo, msg)
					continue
				}
				for _, ticket := range result {
					slog.With(
						"start",
						start,
						"end",
						end,
						"name",
						ticket.Name,
						"date",
						ticket.Date,
						"seats",
						ticket.Seats,
					).Log(
						ctx,
						slog.LevelInfo,
						ticket.Name,
						ticket.Date,
						ticket.Seats,
					)

					msg := fmt.Sprintf("اتوبوس موجود شد %s به %s در ساعت %s به تعداد صندلی %d", s.service.GetCity(start), s.service.GetCity(end), ticket.Date, ticket.Seats)
					err := s.notification.Send(notification.Message{
						Title: "",
						Text:  msg,
					})
					if err != nil {
						slog.Error(err.Error())
					}
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
