package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//记录一道有趣的面试题

//初始版本：输出1-11的数，大多数为11，大量的data-race
//func main() {
//	total, sum := 0, 0
//	for i := 1; i <= 10; i++ {
//		sum += i
//		go func() {
//			total += i
//		}()
//	}
//}

//改进后版本
//解决total在各个协程的data-race可以使用互斥量或者atomic
//使用WaitGroup让主线程先阻塞等待
func main() {
	var total int64 = 0
	var wg sync.WaitGroup
	sum := 0
	wg.Add(10)
	for i := 1; i <= 10; i++ {
		sum += i
		go func(i int) {
			defer wg.Done()
			atomic.AddInt64(&total, int64(i))
		}(i)
	}
	wg.Wait()
	fmt.Printf("total:%d sum:%d", total, sum)
}
