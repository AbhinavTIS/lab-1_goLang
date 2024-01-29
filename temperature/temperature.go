package Temperature

import (
	"fmt" //fmt for formatted I/O,

	"strconv" //strconv for string conversions.
)

func temperatureConverter(temperatureInput string) {

	// Convert input to float
	temperatureCelsius, err := strconv.ParseFloat(temperatureInput[:len(temperatureInput)-1], 64)
	//This line converts the input string to a floating-point number using strconv.ParseFloat
	// It removes the newline character at the end using [:len(temperatureInput)-1]
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid temperature.")
		return
	}

	// Convert Celsius to Fahrenheit using the formula
	temperatureFahrenheit := (9 * temperatureCelsius / 5) + 32

	// Print the result with 2 decimal places
	fmt.Printf("Temperature in Fahrenheit: %.2f\n", temperatureFahrenheit)
}
