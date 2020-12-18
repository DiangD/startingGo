package main

import "fmt"

//leetcode hot100 283 移动零

//moveZeroes 双指针
func moveZeroes(nums []int) {
	n, left, right := len(nums), 0, 0
	for right < n {
		if nums[right] != 0 {
			nums[right], nums[left] = nums[left], nums[right]
			left++
		}
		right++
	}
}

//moveZero 直接覆盖
func moveZero(nums []int) {
	index, n := 0, len(nums)
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			nums[index] = nums[i]
			index++
		}
	}
	for i := index; i < n; i++ {
		nums[i] = 0
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZero(nums)
	fmt.Println(nums)
}
