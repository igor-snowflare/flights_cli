package utils

import (
	"encoding/json"
)

type ResponseHelper struct {
	Data []FormatedResponse `json:"best_flights"`
}

type FormatedResponse struct {
	Flights []Flight `json:"flights"`
	Layovers []Layover `json:"layovers"`
	Duration int `json:"total_duration"`
	Price int `json:"price"`
}

type Flight struct {
	DepartureAirport Airport `json:"departure_airport"`
	ArrivalAirport Airport `json:"arrival_airport"`
	Duration int `json:"duration"`
	Airplane string `json:"airplane"`
	Airline string `json:"airline"`
}

type Layover struct {
	Duration int `json:"duration"`
	Name string `json:"name"`
	AirportCode string `json:"id"`
}

type Airport struct {
	Name string `json:"name"`
	AirportCode string `json:"id"`
	Time string `json:"time"`
}

func ParseResponse(data []byte) (ResponseHelper, error) {
	var parsedData ResponseHelper
	err := json.Unmarshal(data, &parsedData)

	return parsedData, err
}