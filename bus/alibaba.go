package bus

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type JsonBirthDate time.Time

// UnmarshalJSON Implement Marshaler and Unmarshaler interface
func (j *JsonBirthDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02T15:04:05", s)
	if err != nil {
		return err
	}
	*j = JsonBirthDate(t)
	return nil
}

type Alibaba struct {
	prefix string
}

func (a Alibaba) BussesWithChannel(startCity string, destinationCity string, date time.Time, ch chan []Ticket) {
	r, _ := a.Busses(startCity, destinationCity, date)
	ch <- r
}

type BaseAlibabaResponse struct {
	Error struct {
		Message string `json:"message"`
		Code    string `json:"errorCode"`
	} `json:"error"`
	Status bool `json:"success"`
}

type AlibabaAvailableTicket struct {
	ID       string        `json:"proposalId"`
	Name     string        `json:"companyName"`
	Date     string        `json:"departureDate"`
	Time     string        `json:"departureTime"`
	DateTime JsonBirthDate `json:"departureDateTime"`
	Terminal string        `json:"orginTerminal"`
	Seats    int           `json:"availableSeats"`
}

type AlibabaSeatsTicket struct {
	Index  int    `json:"index"`
	Number int    `json:"number"`
	Column int    `json:"column"`
	Row    int    `json:"row"`
	Status string `json:"status"`
}

type AlibabaAvailableResponse struct {
	BaseAlibabaResponse
	Result struct {
		List []AlibabaAvailableTicket `json:"availableList"`
	} `json:"result"`
}

type AlibabaAvailableSeatsResponse struct {
	BaseAlibabaResponse
	Result []TicketSeat
}

func NewAlibaba() *Alibaba {
	return &Alibaba{prefix: "https://ws.alibaba.ir/api/v"}
}

func responseToTicket(response *AlibabaAvailableResponse) []Ticket {
	tickets := make([]Ticket, len(response.Result.List))
	for i, ticket := range response.Result.List {
		tickets[i] = Ticket{
			ID:       ticket.ID,
			Name:     ticket.Name,
			Date:     time.Time(ticket.DateTime),
			Terminal: ticket.Terminal,
			Seats:    ticket.Seats,
		}
	}
	return tickets
}

func (a Alibaba) Busses(startCity string, destinationCity string, time time.Time) ([]Ticket, error) {
	date := fmt.Sprintf("%d-%d-%d", time.Year(), time.Month(), time.Day())

	url := fmt.Sprintf("%s2/bus/available?orginCityCode=%s&destinationCityCode=%s&requestDate=%s&passengerCount=1", a.prefix, startCity, destinationCity, date)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = response.Body.Close()
	}()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var result AlibabaAvailableResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if !result.Status {
		if result.Error.Message == "" {
			return nil, errors.New("unknown error")
		}
		return nil, errors.New(result.Error.Message)
	}

	return responseToTicket(&result), nil
}

func (a Alibaba) TicketSeats(ticket Ticket) (TicketSeats, error) {
	url := fmt.Sprintf("%s1/bus/available/%s/seats", a.prefix, ticket.ID)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = response.Body.Close()
	}()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var result AlibabaAvailableSeatsResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if !result.Status {
		if result.Error.Message == "" {
			return nil, errors.New("unknown error")
		}
		return nil, errors.New(result.Error.Message)
	}

	return result.Result, err
}

func (a Alibaba) GetCity(cityId string) string {
	if cityId == "87330000" {
		return "شاهرود"
	} else if cityId == "11320000" {
		return "تهران"
	}
	return "ناآشنا"
}

func (a Alibaba) Shahroud() string {
	return "87330000"
}

func (a Alibaba) Tehran() string {
	return "11320000"
}
