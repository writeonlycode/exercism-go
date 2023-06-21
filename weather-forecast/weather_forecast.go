// Package weather provides functionality related to weather conditions and
// forecasts.
package weather

// CurrentCondition represents the current weather condition.
var CurrentCondition string

// CurrentLocation represents the current location.
var CurrentLocation string

// Forecast constructs and returns a string that combines the CurrentLocation
// and CurrentCondition values to represent the current weather forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
