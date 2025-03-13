// Package weather reports the current weather forecast.
package weather

// CurrentCondition is used to store the current weather condition.
var CurrentCondition string

// CurrentLocation is used to store the current weather location.
var CurrentLocation string

// Forecast accepts the forcasted city and condition as strings and returns the
// weather condition as a string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
