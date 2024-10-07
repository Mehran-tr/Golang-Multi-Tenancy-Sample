package utils

// Add returns the sum of two integers
func Add(a, b int) int {
	return a + b
}

// IsPrime checks if a given number is prime
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
