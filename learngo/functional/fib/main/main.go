package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func fib() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

//定义函数为类型
type intGen func() int

//实现io.reader接口
func (i intGen) Read(p []byte) (n int, err error) {
	next := i()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	//交给strings代理
	return strings.NewReader(s).Read(p)
}

func printContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fib()
	printContents(f)
}
