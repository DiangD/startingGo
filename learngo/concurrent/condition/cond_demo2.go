package main

import (
	"fmt"
	"sync"
)

type Button struct {
	Clicked *sync.Cond
}

func main() {
	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}
	subscribe := func(c *sync.Cond, fn func()) {
		var tempWg sync.WaitGroup
		tempWg.Add(1)
		go func() {
			tempWg.Done()
			c.L.Lock()
			defer c.L.Unlock()
			//等待阻塞
			c.Wait()
			fn()
		}()
		tempWg.Wait()
	}

	var wg sync.WaitGroup
	wg.Add(3)
	subscribe(button.Clicked, func() {
		fmt.Println("Maximizing window.")
		wg.Done()
	})

	subscribe(button.Clicked, func() {
		fmt.Println("Displaying annoying dialog box!")
		wg.Done()
	})

	subscribe(button.Clicked, func() {
		fmt.Println("Mouse clicked.")
		wg.Done()
	})
	//唤醒所有等待协程
	button.Clicked.Broadcast()
	wg.Wait()
}
