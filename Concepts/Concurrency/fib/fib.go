package main

import (
	"fmt"
	"sync"
)

/**
Fib(93) causes an overflow (Negative in int64), resulting in a
reversed (negative) integer, hence the stop at 92
*/

var fibCache = make([]int, 92)
var mu sync.Mutex

func fib(n int) int {
	if n <= 1 {
		mu.Lock()
		fibCache[n] = 1
		mu.Unlock()
		return 1
	} else if fibCache[n-1] > 0 {
		val := fibCache[n-1] + fibCache[n-2]
		mu.Lock()
		fibCache[n] = val
		mu.Unlock()
		return val
	} else {
		return fib(n-1) + fib(n-2)
	}
}

// uni-directional channels
// jobs can only receive
// results can only send
func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func main() {
	// Buffered channels
	jobs := make(chan int, 92)
	results := make(chan int, 92)

	// Worker pools are fired into the background waiting for messages
	go worker(jobs, results) // worker 1
	go worker(jobs, results) // worker 2
	go worker(jobs, results) // worker 3
	go worker(jobs, results) // worker 4

	for i := 0; i < 92; i++ {
		jobs <- i // Send 92 messages into the jobs buffered channel
	}
	close(jobs)

	for j := 0; j < 92; j++ {
		fmt.Println(<-results) // receive fibonacci numbers from results channel
	}
	close(results)
}
