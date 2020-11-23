package main

import (
	"fmt"
	"sync"
)

//channel的发送接收是阻塞的，一次写入的操作灯带一次读取（阻塞式io）
func chanDemo() {
	var workers [10]worker
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}
	wg.Add(len(workers) * 2)
	for i, w := range workers {
		w.in <- 'a' + i
	}

	for i, w := range workers {
		w.in <- 'A' + i
	}
	wg.Wait()
}

func doWorker(id int, w worker) {
	go func() {
		for n := range w.in {
			fmt.Printf("channel %d revecived %c\n", id, n)

			go func() {
				w.done()
			}()
		}
	}()
}

type worker struct {
	in   chan int
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {

	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWorker(id, w)
	return w
}

func main() {
	chanDemo()
}

//不要通过共享内存来通信，要通过通信来共享内存
