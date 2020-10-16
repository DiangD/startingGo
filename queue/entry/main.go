package main

import (
	"fmt"
	"queue"
)

func main() {
	q := queue.Queue{}
	for i := 0; i < 10; i++ {
		q.Push(i)
	}
	q.Push("test")
	fmt.Println(q, q.IsEmpty())
	for i := 0; i < 10; i++ {
		q.Pop()
	}
	fmt.Println(q, q.IsEmpty())
}
