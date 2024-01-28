package main

import (
	"bufio"   // bufio is used for buffered I/O
	"fmt"     //fmt for formatted I/O,

	"strconv" //strconv for string conversions.
)
func main(){

	fmt.Println("Temperature Converter from Celsius to Fahrenheit:")
	
}


func temperatureConverter() {
	
	// Read temperature in Celsius from user
	fmt.Print("Enter temperature in Celsius: ")
	reader := bufio.NewReader(os.Stdin)              // taking input from the user using bufio
	temperatureInput := reader.ReadString('\n') // it reads the line until the newline is encoundtered, err is used to discard the error returned by ReadString for simplicity.
	}

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
	return temperatureFahrenheit
}
