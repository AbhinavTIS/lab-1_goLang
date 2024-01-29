package isPalindrome;

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
