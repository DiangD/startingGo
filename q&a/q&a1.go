package main

import "fmt"

/*
for range 会创建每个元素的副本（val），而这道题&val拿到的都是副本的地址，最后val被赋值为3，于是结果都是3
由于是副本使用对于val的修改都不会作用到原本的结果，一般for range是只读的
*/

func main() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)
	for key, val := range slice {
		//fix
		//tmp := val m[key] = &tmp
		m[key] = &val
	}
	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}
