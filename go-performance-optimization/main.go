package main

import (
	"fmt"
	"go-performance-optimization/profiler"
	"go-performance-optimization/services"
	"go-performance-optimization/utils"
	"log"
	"time"
)

func main() {
	// Start the pprof profiler
	profiler.StartProfiler()

	// Simulate CPU-intensive task: Generate prime numbers
	start := time.Now()
	limit := 100000
	log.Printf("Generating prime numbers up to %d\n", limit)
	primes := services.GeneratePrimes(limit)
	duration := time.Since(start)
	log.Printf("Generated %d primes in %v\n", len(primes), duration)

	// Simulate memory-intensive task
	log.Println("Creating a large slice")
	largeSlice := utils.CreateLargeSlice(10000000)
	log.Printf("Created a large slice of length %d\n", len(largeSlice))

	// Prevent the program from exiting immediately
	fmt.Println("Press enter to exit...")
	fmt.Scanln()
}
