package utils

import (
	"fmt"
)

func DisplayResponse(destination Destination, responseHelper ResponseHelper) {
	fmt.Println("")
	fmt.Println("===========================================")
	fmt.Printf("%s %s (%s)", destination.Flag, destination.Name, destination.AirportCode)
	fmt.Println("")

	for index, travelOption := range responseHelper.Data {
		firstDeparture := travelOption.Flights[0].DepartureAirport
		lastArrival := travelOption.Flights[len(travelOption.Flights) - 1].ArrivalAirport

		durationMinutes := travelOption.Duration % 60
		durationHours := (travelOption.Duration - durationMinutes) / 60

		fmt.Printf("\nDeparting from %s at %s", firstDeparture.Name, firstDeparture.Time)
		fmt.Printf("\nArriving to %s at %s", lastArrival.Name, lastArrival.Time)
		fmt.Printf("\nTotal Travel Time: %d hours, %d minutes", durationHours, durationMinutes)

		fmt.Println("")
		fmt.Println("")

		if index < len(travelOption.Flights) - 1 {
			fmt.Println("-------------------------------------------")
		}
	}
}