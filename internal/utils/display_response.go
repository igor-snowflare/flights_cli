package utils

import (
	"fmt"
)

func ShowLayoverDuration(airportCode string, layovers []Layover) {
	for _, layover := range layovers {
		if layover.AirportCode == airportCode {
			layoverMinutes := layover.Duration % 60
			layoverHours := (layover.Duration - layoverMinutes) / 60
			fmt.Printf("\n        💤 Layover on Arrival: %d hours %d minutes", layoverHours, layoverMinutes)
		}
	}
}

func DisplayResponse(destination Destination, responseHelper ResponseHelper, currency string) {
	fmt.Println("")
	fmt.Println("===========================================")
	fmt.Printf("%s %s (%s)", destination.Flag, destination.Name, destination.AirportCode)
	fmt.Println("")

	for optionIndex, travelOption := range responseHelper.Data {
		firstDeparture := travelOption.Flights[0].DepartureAirport
		lastArrival := travelOption.Flights[len(travelOption.Flights) - 1].ArrivalAirport

		durationMinutes := travelOption.Duration % 60
		durationHours := (travelOption.Duration - durationMinutes) / 60

		fmt.Printf("\n🛫 Departing from %s at %s", firstDeparture.Name, firstDeparture.Time)
		fmt.Printf("\n🛬 Arriving to %s at %s", lastArrival.Name, lastArrival.Time)
		fmt.Printf("\n⏱️ Total Travel Time: %d hours %d minutes", durationHours, durationMinutes)
		fmt.Printf("\n🛩️ Total Flights: %d", len(travelOption.Flights))
		fmt.Printf("\n💰 Cost: %d %s", travelOption.Price, currency)

		fmt.Println("")
		fmt.Println("")

		fmt.Println("🛄 Transfers")

		for _, flight := range travelOption.Flights {
			flightDurationMinutes := flight.Duration % 60
			flightDurationHours := (flight.Duration - flightDurationMinutes) / 60

			fmt.Printf("\n    %s ➡️ %s", flight.DepartureAirport.Name, flight.ArrivalAirport.Name)
			fmt.Printf("\n        ⏳ Flight Duration: %d hours %d minutes", flightDurationHours, flightDurationMinutes)
			fmt.Printf("\n        👩‍✈️ Airline: %s", flight.Airline)
			fmt.Printf("\n        ✈️ Airplane: %s", flight.Airplane)
			ShowLayoverDuration(flight.ArrivalAirport.AirportCode, travelOption.Layovers)
		}

		if optionIndex < len(responseHelper.Data) - 1 {
			fmt.Println("")
			fmt.Println("")
			fmt.Println("-------------------------------------------")
		}
	}
}