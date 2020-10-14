package main

import "fmt"

func main() {
	var i interface{} = "hello"
	//如果使用ok接受两个返回值，成功则返回该类型的值和true
	//失败则返回该类型的0值与false
	s, ok := i.(float64)
	fmt.Println(s, ok)
	//直接panic
	s1 := i.(float64)
	fmt.Println(s1)
}
