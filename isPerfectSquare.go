package main

import (
	"fmt"
	"math"
)

func isPerfectSquare(num int) bool { //ByAshbir500228410
	if num < 0 {
		return false
	}

	// Calculate the square root of the number
	sqrt := int(math.Sqrt(float64(num)))

	// Check if the square of the square root is equal to the original number
	return sqrt*sqrt == num
}

func main() {
	var number int

	fmt.Print("Enter a number: ")

	// Get input from user
	fmt.Scan(&number)

	// Check if the input is a perfect square
	result := isPerfectSquare(number)

	// Print the result
	fmt.Printf("%d is a perfect square: %t\n", number, result)
}
