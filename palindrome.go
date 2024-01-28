package main

import (
	"strings"
)

func isPalindrome(str string) bool {
	str = strings.ToLower(strings.ReplaceAll(str, " ", ""))
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

// func main() {
// 	var input string
// 	fmt.Print("Enter a string: ")
// 	fmt.Scanln(&input)

// 	if isPalindrome(input) {
// 		fmt.Println("The entered string is a palindrome.")
// 	} else {
// 		fmt.Println("The entered string is not a palindrome.")
// 	}
// }
