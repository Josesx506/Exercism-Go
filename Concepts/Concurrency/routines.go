package main

import (
	"fmt"
	"sync"
	"time"
)

var Words = make([]string, 0, 100)
var mux sync.Mutex

func write_name_serial(word string) {
	Words = append(Words, word)
	time.Sleep(2 * time.Second)
}

func write_name_concurrent(word string, wg *sync.WaitGroup) {
	defer wg.Done()
	mux.Lock() // Lock to avoid race conditions
	Words = append(Words, word)
	mux.Unlock() // Unlock

	time.Sleep(2 * time.Second)
}

func main() {
	startNow := time.Now()
	var wg sync.WaitGroup
	fmt.Println("Concurrency")

	for _, name := range []string{
		"Joses", "Daniel", "Matthew",
		"David", "Eggie", "Tobe",
	} {
		// Concurrent processing
		wg.Add(1)
		go write_name_concurrent(name, &wg)

		// Serial processing
		// write_name_serial(name)
	}

	// Force the program to wait for results before print
	wg.Wait()

	fmt.Printf("Length of array: %d\n", len(Words))
	fmt.Printf("%v \n", Words)
	fmt.Println("This operation took: ", time.Since(startNow))
}
