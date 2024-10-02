# Memory Management and Pointers in Go

This project demonstrates memory management techniques and pointer usage in Go. The program highlights the difference between passing variables by value vs. by pointer and showcases the use of `new` and `make` for memory allocation.
## Explanation:


- **Working with Structs and Pointers**:  
        We define a Person struct and show the difference between passing a struct by value (which doesn't affect the original variable) and by pointer (which does affect the original).

- **Using new and make for Memory Allocation**: 
        We demonstrate how the new function allocates memory for a single variable and returns a pointer.
        We also demonstrate make, which is used to initialize slices, maps, or channels.

- **createPerson**: 
        The createPerson function creates a new Person and returns a pointer to it, demonstrating how Go handles struct creation and pointer returns.

## Understanding Memory Allocation with new and make

    new(T) allocates memory for a single item of type T and returns a pointer to it. It's often used for basic types and structs.
    make(T, len) is used for slices, maps, and channels. It initializes and allocates the underlying data structure, returning a value (not a pointer).

## Sample Output:

    Before modification: {Name:John Age:30}
    After modifyAgeByValue: {Name:John Age:30} (no change)
    After modifyAgeByPointer: {Name:John Age:31} (changed)
    Address of num: 0xc0000b4000, Value of num: 0
    Updated Value of num: 100
    Initial slice nums: [0 0 0 0 0]
    Updated slice nums: [10 0 0 0 0]

## Key Features
- **Pointer vs. Value Passing**: Shows how changes to variables are handled differently when passed by value vs. by pointer.
- **Memory Allocation**: Demonstrates how to use `new` for pointer allocation and `make` for initializing slices.
- **Efficient Memory Usage**: The program is designed to showcase Go's efficient memory management and pointer handling.

## How to Run

```bash
go run main.go
```