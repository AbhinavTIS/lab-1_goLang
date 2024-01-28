package main

func factorial(number int) int {
	result := 1

	for i := 2; i <= number; i++ {
		result = result * i
	}
	return result
}
