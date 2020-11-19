package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	_ "shmiloveu.fun/startingGo/demo/redis/db"
	"shmiloveu.fun/startingGo/demo/redis/lock_demo/lock"
	"sync"
	"time"
)

func main() {
	n := 10
	count := 0
	key := "demo:redis:lock"
	val := uuid.NewV1().String()
	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			isLock := lock.Lock(key, val, time.Second)
			defer wg.Done()
			if !isLock {
				fmt.Println("lock fail")
				return
			}
			fmt.Println("lock success")
			count++
			//time.Sleep(time.Millisecond*500)
			lock.Unlock(key, val)
		}()
	}
	wg.Wait()
	fmt.Printf("count=%d", count)
}
