package main

import "fmt"

//初始化顺序：变量初始化->init()->main()

var t = a()

func init() {
	fmt.Println("init in test_init.go")
}

func main() {
	fmt.Println("calling main")
}

func a() int64 {
	fmt.Println("calling a()")
	return 1024
}
