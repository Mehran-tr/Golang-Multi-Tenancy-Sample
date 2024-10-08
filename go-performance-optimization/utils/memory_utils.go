package utils

// CreateLargeSlice creates a large slice to simulate high memory usage
func CreateLargeSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = i
	}
	return slice
}
