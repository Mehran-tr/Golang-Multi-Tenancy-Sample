# Building Reliable Software with Go's Testing Tools

This task demonstrates how to write reliable software in Go using its testing tools. It covers unit tests, table-driven tests, benchmarks, and examples that can be used for documentation.

## Key Features

- **Unit Tests**: Demonstrates how to write tests for functions using the `testing` package.
- **Table-Driven Tests**: A concise way to run multiple test cases with different inputs.
- **Benchmarking**: Measure the performance of functions using Go’s benchmarking tools.
- **Examples and Documentation Tests**: Write examples that can be used both for documentation and for verifying output during testing.

## Project Structure

 ```
reliable-software-go/
├── go.mod
├── main.go
├── utils/
│   ├── math.go
│   └── math_test.go
├── README.md
            
```


## How to Run Tests


   Run test:
   ```bash
     go test
   ```
   You should see a response like:

    ```
    ok  	reliable-software-go/utils	0.003s

    ```

   Run benchmark:
   ```bash
    go test -bench .
   ```

   You should see a result like:

    ```
    BenchmarkAdd-8      999999999               0.343 ns/op
    BenchmarkIsPrime-8  25948958                45.9 ns/op

    ```
   Run Specific Test or Benchmark:

   ```bash
    go test -run TestIsPrime
    go test -bench BenchmarkAdd
   ```