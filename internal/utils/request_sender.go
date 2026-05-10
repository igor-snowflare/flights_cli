package utils

import (
	"io"
	"log"
	"net/http"
	"time"
)

func SendRequest(apiKey string, departure string, destination Destination, departureDate string, currency string) []byte {
	request, _ := http.NewRequest("GET", "https://serpapi.com/search", nil)
	
	params := request.URL.Query()
	params.Add("engine", "google_flights")
	params.Add("departure_id", departure)
	params.Add("arrival_id", destination.AirportCode)
	params.Add("currency", currency)
	params.Add("type", "2")
	params.Add("outbound_date", departureDate)
	params.Add("api_key", apiKey)

	request.URL.RawQuery = params.Encode()

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	response, err := client.Do(request)

	if err != nil {
		log.Fatal("Upstream server response failed")
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	return body
}