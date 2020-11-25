package main

import (
	"fmt"
	"math"
)

//经典最大子序列问题

//思想 max(dp[i-1]+i ,i)
func maxSubArray(nums []int) int {
	var pre int
	var max = nums[0]
	for _, num := range nums {
		pre = int(math.Max(float64(num), float64(pre+num)))
		max = int(math.Max(float64(max), float64(pre)))
	}
	return max
}

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(arr))
}
