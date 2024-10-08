package services

import (
	"sync"
)

// IsPrime checks if a number is prime (CPU-intensive operation)
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

// GeneratePrimes generates prime numbers up to a certain limit
// This version uses basic iteration and appending primes to the slice.
func GeneratePrimes(limit int) []int {
	// Preallocate memory for the slice to avoid repeated allocations
	primes := make([]int, 0, limit/2) // Rough estimate for number of primes
	for i := 2; i <= limit; i++ {
		if IsPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

// GeneratePrimesConcurrent generates prime numbers using multiple goroutines
// This version splits the range into segments and processes them concurrently.
func GeneratePrimesConcurrent(limit int, workers int) []int {
	primes := make([]int, 0, limit/2) // Preallocate for better performance
	var mu sync.Mutex
	var wg sync.WaitGroup

	step := limit / workers

	// Split the range into multiple chunks for each goroutine to process.
	for i := 0; i < workers; i++ {
		start := i*step + 2
		end := (i + 1) * step
		if i == workers-1 {
			end = limit
		}

		// Launch a goroutine for each range
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			localPrimes := []int{}
			for j := start; j <= end; j++ {
				if IsPrime(j) {
					localPrimes = append(localPrimes, j)
				}
			}
			// Lock the shared primes slice when appending to prevent race conditions
			mu.Lock()
			primes = append(primes, localPrimes...)
			mu.Unlock()
		}(start, end)
	}

	// Wait for all goroutines to complete
	wg.Wait()
	return primes
}
