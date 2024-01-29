package isPrime;

func isPrime(n int) string {
	if n <= 1 {
		return "The number is not prime"
	}

	for i := 2; i <= n/2; i++ {

		if n%i == 0 {
			return "The number is not prime"
		}
	}

	return "The number is prime"
}
