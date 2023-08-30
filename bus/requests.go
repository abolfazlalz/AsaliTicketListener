package bus

import (
	"bus_listener/locations"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type BaseResponse struct {
	Error struct {
		Message string `json:"message"`
		Code    string `json:"errorCode"`
	} `json:"error"`
	Status bool `json:"success"`
}

type BusAvailableResponse struct {
	BaseResponse
	Result struct {
		List []BusAvailable `json:"availableList"`
	}
}

type BusAvailableSeatsResponse struct {
	BaseResponse
	Result []Seats
}

const API_PREFIX = "https://ws.alibaba.ir/api/v"

func GetSeats(id string) ([]Seats, error) {
	url := fmt.Sprintf("%s1/bus/available/%s/seats", API_PREFIX, id)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var result BusAvailableSeatsResponse
	json.Unmarshal(body, &result)

	if !result.Status {
		if result.Error.Message == "" {
			return nil, errors.New("unknown error")
		}
		return nil, errors.New(result.Error.Message)
	}

	return result.Result, err
}

func GetAvailables(start locations.City, end locations.City, time time.Time) ([]BusAvailable, error) {
	date := fmt.Sprintf("%d-%d-%d", time.Year(), time.Month(), time.Day())

	url := fmt.Sprintf("%s2/bus/available?orginCityCode=%s&destinationCityCode=%s&requestDate=%s&passengerCount=1", API_PREFIX, start, end, date)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var result BusAvailableResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if !result.Status {
		if result.Error.Message == "" {
			return nil, errors.New("unknown error")
		}
		return nil, errors.New(result.Error.Message)
	}

	return result.Result.List, nil
}
