package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Custom error type to provide more context
type FileError struct {
	FileName string
	Problem  string
}

func (e *FileError) Error() string {
	return fmt.Sprintf("error with file %s: %s", e.FileName, e.Problem)
}

// readFile reads a file and returns an error if something goes wrong
func readFile(fileName string) (string, error) {
	// Check if the file exists
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return "", &FileError{FileName: fileName, Problem: "file does not exist"}
	}

	// Attempt to read the file
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", &FileError{FileName: fileName, Problem: "failed to read file"}
	}

	return string(data), nil
}

func main() {
	fileName := "test.txt"

	// Example of idiomatic error handling
	content, err := readFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File Content:", content)
	}

	// Demonstrating panic and recover
	demonstratePanicRecovery()
}

func divide(a, b int) int {
	if b == 0 {
		panic("cannot divide by zero")
	}
	return a / b
}

// demonstratePanicRecovery demonstrates how to recover from a panic
func demonstratePanicRecovery() {
	defer func() {
		// recover() captures the panic and prevents the program from crashing
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// Simulating a panic by dividing by zero
	fmt.Println("Attempting division by zero...")
	result := divide(10, 0)        // This will cause a panic
	fmt.Println("Result:", result) // This will not be executed
}
