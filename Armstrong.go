// package main

// import (
	
// 	"fmt"
// 	"strconv"
// )

// func main() {

	
// 	fmt.Println("Kindly enter the input value")
// 	var input string
// 	fmt.Scanln(&input)
// 	inputInt, _ := strconv.Atoi(input)
// 	armstrongNumber(inputInt)
// }

// func armstrongNumber(value int) bool {

// 	//Declaring variables
// 	var sum int = 0
// 	var save int
// 	var rem int
// 	var cube int

// 	//Save the inputNumber
// 	save = value

// 	//Multiplying each digit of the inputInt 3 times
// 	//And Adding each digit cube of inputInt to variable sum
// 	for i := value; i > 0; i = i / 10 {
// 		rem = i % 10
// 		cube = rem * rem * rem
// 		sum = sum + cube
// 	}

// 	// //Check the condition if the sum is equal to the save number
// 	if sum == save {
// 		fmt.Printf("%d,%s", save, "Is a armstrong number")
// 		return true

// 	}
// 	fmt.Printf("%d,%s", save, "Is not a armstrong number")
// 	return false

// }
