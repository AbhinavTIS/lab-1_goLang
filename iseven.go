package main

import "fmt"

func main() {
	// Prompt the user to enter a number
	fmt.Print("Enter a number: ")
	var number int
	fmt.Scan(&number)

}

// Function to check if a number is even
func isEven(num int) bool {
	return num%2 == 0
}
