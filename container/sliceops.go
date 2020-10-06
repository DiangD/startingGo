package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v,len=%d, cap=%d\n", s, len(s), cap(s))
}
func main() {
	var s []int //zero value for slice is nil 切片的0值为nil

	for i := 0; i < 100; i++ {
		//新底层数组的cap为oldCap * 2
		s = append(s, 2*i+1)
		printSlice(s)
	}

	//初始化slice的几种方式
	s1 := []int{1, 2, 3, 4}
	printSlice(s1)
	//使用内建函数
	s2 := make([]int, 16)
	//len = 10 cap = 32 提前预备好空间，避免多次扩容
	s3 := make([]int, 10, 32)
	printSlice(s2)
	printSlice(s3)

	fmt.Println("Copy slice")
	//1.目标切片 2.原切片
	copy(s2, s1)
	printSlice(s2)

	//为什么没有一些高级的方法？？
	fmt.Println("Delete elements from slice")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)
	fmt.Println("Popping from prev")
	prev := s2[0]
	fmt.Println(s2[1:], prev)
	fmt.Println("Popping from tail")
	tail := s2[len(s2)-1]
	fmt.Println(s2[:len(s2)-1], tail)
}
