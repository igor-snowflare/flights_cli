package utils

import (
	"encoding/json"
)

type Destination struct {
	Name string `json:"name"`
	Flag string `json:"flag"`
	AirportCode string `json:"airport_code"`
}

func ParseDestinations(rawData []byte) ([]Destination, error) {
	var parsedDestinations []Destination
	err := json.Unmarshal(rawData, &parsedDestinations)
	return parsedDestinations, err
}