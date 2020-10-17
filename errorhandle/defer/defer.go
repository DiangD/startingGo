package main

import (
	"bufio"
	"fmt"
	"os"
	"shmiloveu.fun/startingGo/functional/fib"
)

/**
1. 确保在函数结束时调用（panic、return）
2. 参数在defer语句时计算
3. defer列表先进后出（stack）
*/
func tryDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
		if i > 3 {
			panic("time to panic")
		}
	}
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fib.Fib()
	for i := 0; i < 20; i++ {
		_, _ = fmt.Fprintln(writer, f())
	}
}

func main() {
	tryDefer()
	writeFile("errorhandle/defer/defer_test.txt")
}
