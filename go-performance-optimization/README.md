# Tuning Go Code for Maximum Performance

This project demonstrates how to profile and optimize Go applications for maximum performance. It covers memory and CPU optimization, using Go's concurrency model efficiently, and leveraging `pprof` for performance profiling.

## Key Features

- **Profiling with pprof**: Measure CPU and memory usage to find performance bottlenecks.
- **Memory Optimization**: Reduce memory allocations to minimize garbage collection overhead.
- **Concurrency Optimization**: Use goroutines and synchronization primitives to handle CPU-bound tasks efficiently.


## Project Structure

 ```

    go-performance-optimization/
    ├── go.mod
    ├── main.go
    ├── profiler/
    │   └── profiler.go
    ├── services/
    │   └── compute_service.go
    ├── utils/
    │   └── memory_utils.go
    ├── README.md
          
```

## How to Run

- **in cmd**:

    ```
        go run main.go
    ```

- **Access the profiler:**:
    ```
        http://localhost:6060/debug/pprof/
    ```

- **Example Output**:
    ```
        Generating prime numbers up to 100000
        Generated 9592 primes in 1.231486142s
        Creating a large slice
        Created a large slice of length 10000000
        Press enter to exit...

    ```

