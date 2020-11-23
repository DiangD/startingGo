package main

import (
	"fmt"
	"sync"
)

type Lock struct {
	ch chan struct{}
}

func NewLock() *Lock {
	var lock Lock
	lock.ch = make(chan struct{}, 1)
	lock.ch <- struct{}{}
	return &lock
}

func (l *Lock) Lock() bool {
	result := false
	select {
	case <-l.ch:
		result = true
	}
	return result
}

func (l *Lock) Unlock() {
	l.ch <- struct{}{}
}

var counter int

func main() {
	var lock = NewLock()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !lock.Lock() {
				// log error
				fmt.Println("lock_demo failed")
				return
			}
			counter++
			fmt.Println("current counter", counter)
			lock.Unlock()
		}()
	}
	wg.Wait()
}
