package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/igor-snowflare/flights_cli/internal/utils"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Could not open the .env file!")
	}

	apiKey := os.Getenv("API_KEY")
	currency := os.Getenv("CURRENCY")

	destinationData, err := os.ReadFile("destinations.json")

	if err != nil {
		log.Fatal("Parsing of the destinations.json failed!")
	}

	parsedDestinations, err := utils.ParseDestinations(destinationData)

	if err != nil {
		log.Fatal("Parsing destinations into JSON failed")
	}

	departureDate := time.Now().AddDate(0, 0, 14).Format(time.DateOnly)

	fmt.Println("🔍 Looking for flights departing on", departureDate)

	for _, parsedDestination := range parsedDestinations {
		response := utils.SendRequest(apiKey, os.Getenv("ORIGIN_CODE"), parsedDestination, departureDate, currency)

		formatted_response, err := utils.ParseResponse(response)

		if err != nil {
			log.Fatal("Failed to parse the upstream JSON response")
		}

		utils.DisplayResponse(parsedDestination, formatted_response, currency)
	}
}