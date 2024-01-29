package isPerfectSquare

import (
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
