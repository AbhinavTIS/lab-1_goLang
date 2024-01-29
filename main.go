package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Function to generate Fibonacci series up to n terms
func createfibonacci(n int) []int { // By Jevica

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

func isPerfectSquare(num int) bool { //ByAshbir 500228410
	if num < 0 {
		return false
	}

	// Calculate the square root of the number
	sqrt := int(math.Sqrt(float64(num)))

	// Check if the square of the square root is equal to the original number
	return sqrt*sqrt == num
}

func isEven(num int) bool {
	return num%2 == 0
}

func isPrime(n int) {
	if n <= 1 {
		fmt.Println("it is not a prime number")
	}

	for i := 2; i <= n/2; i++ {

		if n%i == 0 {
			fmt.Println("it is a prime number")
		} else {
			fmt.Println("it is not a prime number")
		}
	}

}

func isPalindrome(str string) {
	str = strings.ToLower(strings.ReplaceAll(str, " ", ""))
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			fmt.Println("it is a palindrome")
		}
	}
	fmt.Println("it is not a palindrome")
}

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
func factorial(number int) int {
	result := 1

	for i := 2; i <= number; i++ {
		result = result * i
	}
	return result
}
func armstrongNumber(value int) bool {

	//Declaring variables
	var sum int = 0
	var save int
	var rem int
	var cube int

	//Save the inputNumber
	save = value

	//Multiplying each digit of the inputInt 3 times
	//And Adding each digit cube of inputInt to variable sum
	for i := value; i > 0; i = i / 10 {
		rem = i % 10
		cube = rem * rem * rem
		sum = sum + cube
	}

	// //Check the condition if the sum is equal to the save number
	if sum == save {
		fmt.Println("%d,%s", save, "Is a armstrong number")
		return true

	}
	return false

}

func isLeapYear(year int) {

	if year%4 == 0 {

		if year%100 != 0 || year%100 == 0 && year%400 == 0 {

			fmt.Println("It is a leap year")
		}

	} else {
		fmt.Println("It is a not leap year")
	}
}

func main() {

	fmt.Println("Kindly enter the input value to create fibonacci series")
	var input string
	fmt.Scanln(&input)
	// Convert the input to an integer
	number, err := strconv.Atoi(input)
	if err != nil || number <= 0 {
		fmt.Println("Invalid input. Please enter a positive integer.")
		return
	}

	fibSeries := createfibonacci(number)
	fmt.Printf("Fibonacci Series up to %d terms: %v\n", number, fibSeries)

	//Perfect Square
	var perfectnumber int

	fmt.Print("Enter a number to check a perfect square: ")

	// Get input from user
	fmt.Scan(&perfectnumber)

	// Check if the input is a perfect square
	result := isPerfectSquare(perfectnumber)

	// Print the result
	fmt.Printf("%d is a perfect square: %t\n", perfectnumber, result)

	// Prompt the user to enter a number
	fmt.Print("Enter a number to check if number is prime: ")
	var checkPrimeNumber int
	fmt.Scan(&checkPrimeNumber)

	isPrime(checkPrimeNumber)

	var isaPalindromeNumber string
	fmt.Print("Enter a string to check if string is palindrome: ")
	fmt.Scanln(&isaPalindromeNumber)

	isPalindrome(isaPalindromeNumber)

	// Read temperature in Celsius from user
	fmt.Print("Enter temperature in Celsius: ")
	var temperature string
	fmt.Scanln(&temperature)
	temperatureConverter(temperature)

	fmt.Println("Kindly enter the input value to check if number is armstrong number")
	var isArmstrongNumber string
	fmt.Scanln(&isArmstrongNumber)
	inputInt, _ := strconv.Atoi(isArmstrongNumber)
	armstrongNumber(inputInt)

	fmt.Println("Kindly enter the input value to check year is leap year")
	var isLeapYearNumber string
	fmt.Scanln(&isLeapYearNumber)

	isLeapYearNumberInt, _ := strconv.Atoi(isLeapYearNumber)
	isLeapYear(isLeapYearNumberInt)
}
