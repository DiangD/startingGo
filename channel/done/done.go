package main

import (
	"fmt"
	"time"
)

func chanDemo() {
	var channels []chan<- int
	for i := 0; i < 10; i++ {
		channels = append(channels, createWorker(i))
	}

	for i := 0; i < len(channels); i++ {
		channels[i] <- 'a' + i
	}

	time.Sleep(time.Millisecond)
}

func worker(id int,ch chan int) {
	go  func() {
		for n := range ch {
			fmt.Printf("channel %d revecived %d\n", id, n)
		}
		//for {
		//	if n,ok := <-ch;ok{
		//		fmt.Printf("channel %d revecived %d\n", id, n)
		//	}else {
		//		break
		//	}
		//}
	}()
}

func createWorker(id int) chan<- int {
	ch := make(chan int)
	go worker(id,ch)
	return ch
}

func bufferChannel() {
	c :=make(chan int,3)
	go worker(1,c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	//close后接收方依然可以收到该类型的0值
	//只能由发送方close
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	//chanDemo()
	bufferChannel()
}
