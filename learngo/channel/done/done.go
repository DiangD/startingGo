package main

import (
	"fmt"
)

//channel的发送接收是阻塞的，一次写入的操作灯带一次读取（阻塞式io）
func chanDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}
	for i, w := range workers {
		w.in <- 'a' + i
	}

	for i, w := range workers {
		w.in <- 'A' + i
	}

	for _, w := range workers {
		<-w.done
	}
}

func doWorker(id int, w worker) {
	go func() {
		for n := range w.in {
			fmt.Printf("channel %d revecived %c\n", id, n)

			go func() {
				w.done <- true
			}()
		}
	}()
}

type worker struct {
	in   chan int
	done chan bool
}

func createWorker(id int) worker {

	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWorker(id, w)
	return w
}

func main() {
	chanDemo()
}
