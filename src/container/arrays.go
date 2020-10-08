package main

import "fmt"

//数组是值类型，作为参数会对数组进行拷贝
//go 一般不直接使用数组，使用切片
//arr [5]int -> array  arr []int -> slice
func printArray(arr []int) {
	for _, v := range arr {
		fmt.Println(v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3, grid)

	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}
	//转换为切片
	printArray(arr3[:])

}
