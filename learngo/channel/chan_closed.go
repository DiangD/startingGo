package main

import "fmt"

/**
读已经关闭的chan能一直读到东西，但是读到的内容根据通道内关闭前是否有元素而不同。

如果chan关闭前，buffer内有元素还未读, 会正确读到chan内的值，且返回的第二个 bool 值（是否读成功）为true。

如果chan关闭前，buffer内有元素已经被读完，chan内无值，接下来所有接收的值都会非阻塞直接成功，返回channel元素的零值，但是第二个bool值一直为false。

写已经关闭的chan会panic
*/
func main() {
	//初始带buffer的channel
	c := make(chan int, 3)
	c <- 1
	c <- 2
	//关闭channel
	close(c)

	//取buffer
	num, ok := <-c
	num1, ok := <-c
	fmt.Println(num, ok)
	fmt.Println(num1, ok)

	//buffer为空，取0值
	num2, ok := <-c
	num3, ok := <-c
	fmt.Println(num2, ok)
	fmt.Println(num3, ok)

	//panic: send on closed channel
	//无法写入已关闭的channel
	c <- 1
}
