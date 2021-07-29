package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var mu sync.Mutex

	increment := func() {
		mu.Lock()
		defer mu.Unlock()
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}

	decrement := func() {
		mu.Lock()
		defer mu.Unlock()
		count--
		fmt.Printf("Decrementing: %d\n", count)
	}

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			decrement()
		}()
	}

	wg.Wait()
	fmt.Printf("\nArithmetic complete.\n")
}
