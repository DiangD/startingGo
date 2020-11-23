package main

import (
	"bufio"
	"fmt"
	"os"
	"shmiloveu.fun/startingGo/learngo/functional/fib"
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
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	//一般的错误处理逻辑
	if err != nil {
		//自定义error            z
		//err = errors.New("this is a custom error")
		err = fmt.Errorf("msg:%s", err.Error())
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s ,%s ,%s\n", pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		return
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
	//tryDefer()
	writeFile("errorhandle/defer/defer_test.txt")
}
