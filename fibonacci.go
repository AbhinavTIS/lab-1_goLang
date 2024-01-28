package main

// // Function to generate Fibonacci series up to n terms
// func createfibonacci(n int) []int {

// 	if n <= 0 {
// 		fmt.Print("Inavlid Input")
// 		main()
// 	}
// 	fibSeries := make([]int, n)
// 	fibSeries[0], fibSeries[1] = 0, 1

// 	for i := 2; i < n; i++ {
// 		fibSeries[i] = fibSeries[i-1] + fibSeries[i-2]
// 	}

// 	return fibSeries
// }

// func main() {
// 	// Prompt the user to enter the number of terms
// 	fmt.Print("Enter the number of terms for the Fibonacci series: ")
// 	var input string
// 	fmt.Scanln(&input)

// 	// Convert the input to an integer
// 	n, err := strconv.Atoi(input)
// 	if err != nil || n <= 0 {
// 		fmt.Println("Invalid input. Please enter a positive integer.")
// 		return
// 	}

// 	// Generate and print the Fibonacci series
// 	fibSeries := createFibonacci(n)
// 	fmt.Printf("Fibonacci Series up to %d terms: %v\n", n, fibSeries)
// }
