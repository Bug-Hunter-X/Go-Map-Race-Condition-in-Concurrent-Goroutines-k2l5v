# Go Map Race Condition
This repository demonstrates a race condition that can occur when accessing a Go map concurrently from multiple goroutines without proper synchronization. The bug.go file contains the buggy code, while bugSolution.go provides a solution using a mutex to protect the map.

## Problem
The main function creates 1000 goroutines, each incrementing a counter within a map. Without synchronization, multiple goroutines may try to modify the map simultaneously, leading to data corruption and unexpected behavior (inconsistent map sizes, missed updates, etc.).

## Solution
The solution utilizes a sync.Mutex to protect the map from concurrent access. The mutex ensures that only one goroutine can access the map at a time, preventing data races and ensuring data consistency. 

## Running the code
1. Clone the repository.
2. Navigate to the repository's directory in your terminal.
3. Run the buggy code using `go run bug.go` and observe the output. The map's length will likely be less than 1000.
4. Then run the solution using `go run bugSolution.go` and observe the correct output with the map length being 1000. 

This example highlights the importance of using synchronization mechanisms when working with shared resources in concurrent Go programs. 