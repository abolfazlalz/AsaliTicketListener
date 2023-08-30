package bus

type BusAvailable struct {
	ID       string `json:"proposalId"`
	Name     string `json:"companyName"`
	Date     string `json:"departureDate"`
	Time     string `json:"departureTime"`
	Terminal string `json:"orginTerminal"`
}

type Seats struct {
	Index  int    `json:"index"`
	Number int    `json:"number"`
	Column int    `json:"column"`
	Row    int    `json:"row"`
	Status string `json:"status"`
}

func (ba *BusAvailable) Seats() ([]Seats, error) {
	return GetSeats(ba.ID)
}
