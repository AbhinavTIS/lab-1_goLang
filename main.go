package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("Kindly enter the input value")
	var input string
	fmt.Scanln(&input)
	// Convert the input to an integer
	number, err := strconv.Atoi(input)
	if err != nil || number <= 0 {
		fmt.Println("Invalid input. Please enter a positive integer.")
		return
	}

	fibSeries := createFibonacci(number)
	fmt.Printf("Fibonacci Series up to %d terms: %v\n", number, fibSeries)

}

// Function to generate Fibonacci series up to n terms
func createfibonacci(n int) []int {

	if n <= 0 {
		fmt.Print("Inavlid Input")
		main()
	}
	fibSeries := make([]int, n)
	fibSeries[0], fibSeries[1] = 0, 1

	for i := 2; i < n; i++ {
		fibSeries[i] = fibSeries[i-1] + fibSeries[i-2]
	}

	return fibSeries
}
