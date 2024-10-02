package main

import (
	"fmt"
)

// Person is a struct with name and age fields
type Person struct {
	Name string
	Age  int
}

// modifyAgeByValue modifies the age field of a Person by value
func modifyAgeByValue(p Person) {
	p.Age += 1
}

// modifyAgeByPointer modifies the age field of a Person by pointer
func modifyAgeByPointer(p *Person) {
	p.Age += 1
}

// demonstrateMakeAndNew shows how new and make allocate memory
func demonstrateMakeAndNew() {
	// Using new to allocate memory for an int
	num := new(int) // new returns a pointer to a newly allocated memory
	fmt.Printf("Address of num: %p, Value of num: %d\n", num, *num)
	*num = 100
	fmt.Printf("Updated Value of num: %d\n", *num)

	// Using make to allocate memory for a slice
	nums := make([]int, 5)
	fmt.Printf("Initial slice nums: %v\n", nums)
	nums[0] = 10
	fmt.Printf("Updated slice nums: %v\n", nums)
}

// createPerson returns a pointer to a new Person struct
func createPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

func main() {
	// Creating a new Person
	p1 := createPerson("John", 30)
	fmt.Printf("Before modification: %+v\n", *p1)

	// Modifying age by value (will not change original)
	modifyAgeByValue(*p1)
	fmt.Printf("After modifyAgeByValue: %+v (no change)\n", *p1)

	// Modifying age by pointer (will change original)
	modifyAgeByPointer(p1)
	fmt.Printf("After modifyAgeByPointer: %+v (changed)\n", *p1)

	// Demonstrating memory management with new and make
	demonstrateMakeAndNew()
}
