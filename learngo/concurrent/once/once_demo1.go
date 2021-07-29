package main

import (
	"fmt"
	"sync"
)

/**
sync.Once确保了即使在不同的goroutine上，调用Do传入的函数只执行一次
为sync.Once只计算Do被调用的次数，而不是调用传入Do的唯一 函数的次数
*/
func main() {
	var count int
	increment := func() {
		count++
	}

	var once sync.Once
	var wg sync.WaitGroup

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			once.Do(increment)
		}()
	}
	wg.Wait()
	fmt.Printf("Count is %d\n", count)
}
