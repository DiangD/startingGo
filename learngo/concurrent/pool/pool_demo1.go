package main

import (
	"fmt"
	"sync"
)

func main() {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance.")
			return struct{}{}
		},
	}

	pool.Get()             //new
	instance := pool.Get() //new
	pool.Get()             //new
	pool.Put(instance)     //exist one instance
	pool.Get()             //get one
}
