// Package weather tells you the current weather conditions at a specified location.
package main

/**
learn about string formatting
*/

var (
	// CurrentCondition represents the current weather conditions as a string.
	CurrentCondition string
	// CurrentLocation represents the current user location as a string.
	CurrentLocation string
)

// Forecast returns a string value describing the weather condition at a location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

func main() {
	println(Forecast("Lagos", "rainy"))
}
