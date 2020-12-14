package main

import (
	"fmt"
)

//leetcode hot100——简单 只出现一次的数字

//最垃圾的hashmap解法
func singleNumber(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		if _, ok := m[num]; !ok {
			m[num] = 1
		} else {
			m[num]++
		}
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return -1
}

//异或
//特性  a ^ 0 = a; a ^ a = 0;满足结合律和交换律
func singleNumberAdv(nums []int) int {
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		ans = ans ^ nums[i]
	}
	return ans
}

func main() {
	//4^1^2^1^2 === (1^1)^(2^2)^4 === 0^4===4
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumberAdv(nums))
}
