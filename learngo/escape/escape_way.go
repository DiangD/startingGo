package main

import "fmt"

//在方法内把局部变量指针返回 局部变量原本应该在栈中分配，在栈中回收。但是由于返回时被外部引用，因此其生命周期大于栈，则溢出。
//发送指针或带有指针的值到 channel 中。 在编译时，是没有办法知道哪个 goroutine 会在 channel 上接收数据。所以编译器没法知道变量什么时候才会被释放。
//在一个切片上存储指针或带指针的值。 一个典型的例子就是 []*string 。这会导致切片的内容逃逸。尽管其后面的数组可能是在栈上分配的，但其引用的值一定是在堆上。
//slice 的背后数组被重新分配了，因为 append 时可能会超出其容量( cap )。 slice 初始化的地方在编译时是可以知道的，它最开始会在栈上分配。如果切片背后的存储要基于运行时的数据进行扩充，就会在堆上分配。
//在 interface 类型上调用方法。 在 interface 类型上调用方法都是动态调度的 —— 方法的真正实现只能在运行时知道。想像一个 io.Reader 类型的变量 r , 调用 r.Read(b) 会使得 r 的值和切片b 的背后存储都逃逸掉，所以会在堆上分配。

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
