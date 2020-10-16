package main

import (
	"fmt"
)

func updateSlice(arr []int) {
	arr[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	//左开右闭区间
	s := arr[2:6]
	fmt.Println("arr[2:6] =", s)
	fmt.Println("arr[:6] =", arr[:6])
	fmt.Println("arr[2:] =", arr[2:])
	fmt.Println("arr[:] =", arr[:])
	fmt.Println("---------------------")
	fmt.Println("After updateSlice(s)")
	updateSlice(s)
	fmt.Println(s)
	fmt.Println(arr)
	//可以重复取切片
	fmt.Println("Re Slice")
	s = s[2:]
	fmt.Println(s)

	//切片可以向后扩展，不可以向前扩展
	//s[i]不可以超越len 向后扩展不可以超越底层数组的cap
	fmt.Println("Extending slice")
	fmt.Println("arr = ", arr)
	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Printf("s1 = %v,len(s1) = %d,cap(s1) = %d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2 = %v,len(s2) = %d,cap(s2) = %d\n", s2, len(s2), cap(s2))

	//添加元素时如果超越cap，系统重新分配更大的底层数组
	//由于值传递，必须接收返回的slice
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s2 s3 ,s4 ,s5 = ", s2, s3, s4, s5)
	//s4 and s5 no longer view arr
	fmt.Println("arr = ", arr)

}
