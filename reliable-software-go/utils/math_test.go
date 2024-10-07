package utils

import (
	"fmt"
	"testing"
)

// TestAdd tests the Add function
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// TestIsPrime tests the IsPrime function
func TestIsPrime(t *testing.T) {
	if !IsPrime(7) {
		t.Errorf("Expected true, but got false for prime number 7")
	}
	if IsPrime(4) {
		t.Errorf("Expected false, but got true for non-prime number 4")
	}
}

// TestIsPrimeTableDriven runs table-driven tests for the IsPrime function
func TestIsPrimeTableDriven(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{10, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("IsPrime(%d)", tt.input), func(t *testing.T) {
			result := IsPrime(tt.input)
			if result != tt.expected {
				t.Errorf("IsPrime(%d) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// BenchmarkAdd benchmarks the Add function
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(2, 3)
	}
}

// BenchmarkIsPrime benchmarks the IsPrime function
func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(101)
	}
}

// ExampleAdd provides an example for the Add function, used for both documentation and testing
func ExampleAdd() {
	fmt.Println(Add(2, 3))
	// Output: 5
}
