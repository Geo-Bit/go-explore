// Concurency vs Parallelism
// GO natively takes advnantage of multiple cores
// Concurrency: Design pattern. A way to write your code. Code that has the potential to run in parrell
// Parallelism: Works with more than 1 core. Allows your code to execute in parallel to take advantage of multi-core resources

// WaitGroup
// Writing concurrent code
// With go routines, the main function does not wait for the go routines to execute
// Once main is done, the program is done.
// Use sync routines to wait for routines to complete before main completes
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutine\t", runtime.NumGoroutine())

	wg.Add(1) // synchronization primitive
	go foo()
	bar()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutine\t", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done() // Tells waitgroup to wait for this before letting main end
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}