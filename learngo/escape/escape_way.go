package main

import "fmt"

//指针逃逸
func pointerEscape() *user {
	//u 本身为一指针，其指向的内存地址不会是栈而是堆
	u := new(user)
	u.username = "qzh"
	u.password = "123456"
	return u
}

//栈空间不足
// make([]int, 10000, 10000) escapes to heap
func stackSizeExceed() {
	//实际上当栈空间不足以存放当前对象时或无法判断当前切片长度时会将对象分配到堆中
	s := make([]int, 10000, 10000)
	for i := range s {
		s[i] = i
	}
}

//动态类型逃逸（不确定长度大小）
//包括不能确定长度大小的切片...
func dynamicType() {
	var s string
	s = "dynamic type may escape"
	//很多函数参数为interface类型，
	//比如fmt.Println(a …interface{})，编译期间很难确定其参数的具体类型，也能产生逃逸
	fmt.Println(s)
}
