package main

import (
	"fmt"
	"runtime"
	"time"
)
//Goroutine 轻量级线程
//非抢占式的多任务处理，由协程主动交出控制权
//编译器/虚拟机/解释器层面的多任务
//多个协程可以在一个或多个线程中运行
func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i]++
				fmt.Printf("Print from goroutine %d\n",i)
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}

/**
goroutine可能的切换点
io/select channel 等待锁 函数调用 runtime.Gosched()
 */