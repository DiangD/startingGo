package main

import "sync"

//dead lock 循环调用

func main() {
	var onceA, onceB sync.Once
	var initB func()
	initA := func() { onceB.Do(initB) }
	initB = func() {
		onceA.Do(initA)
	}
	onceA.Do(initA)
}
